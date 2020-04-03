package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"path"
	"sync"

	"github.com/hooone/evening/pkg/api/middleware"
	"github.com/hooone/evening/pkg/api/routing"
	httpstatic "github.com/hooone/evening/pkg/api/static"
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/log"
	"github.com/hooone/evening/pkg/registry"
	"github.com/hooone/evening/pkg/services/action"
	"github.com/hooone/evening/pkg/services/auth"
	"github.com/hooone/evening/pkg/services/card"
	"github.com/hooone/evening/pkg/services/field"
	"github.com/hooone/evening/pkg/services/navigation"
	"github.com/hooone/evening/pkg/services/parameter"
	"github.com/hooone/evening/pkg/services/style"
	"github.com/hooone/evening/pkg/setting"
	"github.com/hooone/evening/pkg/util/errutil"
	macaron "gopkg.in/macaron.v1"
)

func init() {
	registry.Register(&registry.Descriptor{
		Name:         "HTTPServer",
		Instance:     &HTTPServer{},
		InitPriority: registry.High, //优先级
	})
}

//HTTPServer framework:macaron
type HTTPServer struct {
	log     log.Logger
	macaron *macaron.Macaron
	context context.Context
	httpSrv *http.Server

	RouteRegister routing.RouteRegister `inject:""`
	Bus           bus.Bus               `inject:""`

	AuthTokenService  *auth.UserTokenService        `inject:""`
	NavigationService *navigation.NavigationService `inject:""`
	CardService       *card.CardService             `inject:""`
	FieldService      *field.FieldService           `inject:""`
	ActionService     *action.ActionService         `inject:""`
	StyleService      *style.StyleService           `inject:""`
	ParameterService  *parameter.ParameterService   `inject:""`

	Cfg *setting.Cfg `inject:""`
}

//Init :services initialization
func (hs *HTTPServer) Init() error {
	hs.log = log.New("http.server")

	hs.macaron = hs.newMacaron()
	hs.registerRoutes()

	return nil
}

//Run register route and start http listener
//ctx : use to control service lifecycle
func (hs *HTTPServer) Run(ctx context.Context) error {
	hs.log.Info("HTTP Server Run")
	hs.context = ctx

	hs.applyRoutes()
	hs.httpSrv = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", setting.HTTPAddr, setting.HTTPPort),
		Handler: hs.macaron,
	}

	// http liste
	var listener net.Listener
	var err error
	listener, err = net.Listen("tcp", hs.httpSrv.Addr)
	if err != nil {
		return errutil.Wrapf(err, "failed to open listener on address %s", hs.httpSrv.Addr)
	}
	hs.log.Info("HTTP Server Listen", "address", listener.Addr().String(), "protocol",
		setting.Protocol, "subUrl", setting.AppSubURL, "socket", setting.SocketPath)

	//协程异步锁
	var wg sync.WaitGroup
	wg.Add(1)

	//接收ctx.Done()传来的程序关闭事件，释放异步锁，使service结束
	// handle http shutdown on server context done
	go func() {
		defer wg.Done()

		<-ctx.Done()
		if err := hs.httpSrv.Shutdown(context.Background()); err != nil {
			hs.log.Error("Failed to shutdown server", "error", err)
		}
	}()

	if err := hs.httpSrv.Serve(listener); err != nil {
		if err == http.ErrServerClosed {
			hs.log.Debug("server was shutdown gracefully")
			return nil
		}
		return err
	}

	wg.Wait()

	return nil
}

//new macaron对象
func (hs *HTTPServer) newMacaron() *macaron.Macaron {
	macaron.Env = setting.Env
	m := macaron.New()

	// automatically set HEAD for every GET
	m.SetAutoHead(true)

	return m
}

func (hs *HTTPServer) applyRoutes() {
	// start with middlewares & static routes
	hs.addMiddlewaresAndStaticRoutes()
	// then add view routes & api routes
	hs.RouteRegister.Register(hs.macaron)
	// lastly not found route
	//hs.macaron.NotFound(hs.NotFoundHandler)
}

//添加中间件及静态路由
func (hs *HTTPServer) addMiddlewaresAndStaticRoutes() {
	m := hs.macaron
	//绑定静态文件
	hs.mapStatic(m, setting.StaticRootPath, "dist", "/static")
	// hs.mapStatic(m, setting.StaticRootPath, "robots.txt", "robots.txt")

	//用户身份验证
	m.Use(middleware.GetContextHandler(hs.AuthTokenService))

	//把public>views中的文件作为页面渲染模板
	m.Use(macaron.Renderer(macaron.RenderOptions{
		Directory:  path.Join(setting.StaticRootPath, "dist"),
		IndentJSON: macaron.Env != macaron.PROD,
		Delims:     macaron.Delims{Left: "[[", Right: "]]"},
	}))
}

//把静态文件注册进中间件，在中间件中判断如果是文件GET请求，则把文件内容写入ctx，并返回true
func (hs *HTTPServer) mapStatic(m *macaron.Macaron, rootDir string, dir string, prefix string) {
	headers := func(c *macaron.Context) {
		c.Resp.Header().Set("Cache-Control", "public, max-age=3600")
	}

	if prefix == "public/build" {
		headers = func(c *macaron.Context) {
			c.Resp.Header().Set("Cache-Control", "public, max-age=31536000")
		}
	}

	if setting.Env == setting.DEV {
		headers = func(c *macaron.Context) {
			c.Resp.Header().Set("Cache-Control", "max-age=0, must-revalidate, no-cache")
		}
	}
	logger := log.New("test")
	logger.Info("setting", "StaticRootPath", rootDir)
	m.Use(httpstatic.Static(
		path.Join(rootDir, dir),
		httpstatic.StaticOptions{
			SkipLogging: true,
			Prefix:      prefix,
			AddHeaders:  headers,
		},
	))
}
