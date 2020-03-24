package setting

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/hooone/evening/pkg/infra/log"
	ini "gopkg.in/ini.v1"
)

//Scheme supported
type Scheme string

//Scheme supported
const (
	HTTP   Scheme = "http"
	HTTPS  Scheme = "https"
	HTTP2  Scheme = "h2"
	SOCKET Scheme = "socket"

	DefaultHTTPAddr  string = "0.0.0.0"
	RedattedPassword string = "*********"
	AuthProxySyncTTL int    = 60
)

//Build mode
const (
	DEV     = "development"
	PROD    = "production"
	TEST    = "test"
	APPName = "evening"
)

// App settings.
var (
	Env              = DEV
	AppURL           string
	AppSubURL        string
	ServeFromSubPath bool
)

// build
var (
	BuildVersion    string
	BuildCommit     string
	BuildBranch     string
	BuildStamp      int64
	ApplicationName string
)

// Paths
var (
	HomePath       string
	CustomInitPath = "conf/custom.ini"
)

// Global setting objects.
var (
	Raw       *ini.File
	IsWindows bool
)

// for logging purposes
var (
	configFiles                  []string
	appliedCommandLineProperties []string
	appliedEnvOverrides          []string
)

// Http server options
var (
	Protocol         Scheme
	Domain           string
	HTTPAddr         string
	HTTPPort         string
	SSHPort          int
	CertFile         string
	KeyFile          string
	SocketPath       string
	RouterLogging    bool
	DataProxyLogging bool
	DataProxyTimeout int
	StaticRootPath   string
	EnableGzip       bool
	EnforceDomain    bool

	// Http auth
	AdminUser            string
	AdminPassword        string
	LoginCookieName      string
	LoginMaxLifetimeDays int

	// Security settings.
	SecretKey                        string
	SignoutRedirectUrl               string
	DisableGravatar                  bool
	EmailCodeValidMinutes            int
	DataProxyWhiteList               map[string]bool
	DisableBruteForceLoginProtection bool
	CookieSecure                     bool
	CookieSameSiteDisabled           bool
	CookieSameSiteMode               http.SameSite
)

//Cfg : setting for service instance
type Cfg struct {
	Raw    *ini.File
	Logger log.Logger

	// HTTP Server Settings
	AppURL           string
	AppSubURL        string
	ServeFromSubPath bool

	// Paths
	ProvisioningPath string
	DataPath         string
	LogsPath         string

	// Auth
	LoginCookieName              string
	LoginMaxInactiveLifetimeDays int
	LoginMaxLifetimeDays         int
	TokenRotationIntervalMinutes int

	// Security
	DisableInitAdminCreation         bool
	DisableBruteForceLoginProtection bool
	CookieSecure                     bool
	CookieSameSiteDisabled           bool
	CookieSameSiteMode               http.SameSite
}

//CommandLineArgs : parameters from command line
type CommandLineArgs struct {
	Config   string
	HomePath string
	Args     []string
}

func init() {
	IsWindows = runtime.GOOS == "windows"
}

//NewCfg : Cfg constructor
func NewCfg() *Cfg {
	return &Cfg{
		Logger: log.New("settings"),
	}
}

func parseAppURLAndSubURL(section *ini.Section) (string, string, error) {
	appURL, err := valueAsString(section, "root_url", "http://localhost:3000/")
	if err != nil {
		return "", "", err
	}
	if appURL[len(appURL)-1] != '/' {
		appURL += "/"
	}

	// Check if has app suburl.
	url, err := url.Parse(appURL)
	if err != nil {
		log.Fatal(4, "Invalid root_url(%s): %s", appURL, err)
	}
	appSubURL := strings.TrimSuffix(url.Path, "/")

	return appURL, appSubURL, nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func setHomePath(args *CommandLineArgs) {
	if args.HomePath != "" {
		HomePath = args.HomePath
		return
	}

	HomePath, _ = filepath.Abs(".")
	// check if homepath is correct
	if pathExists(filepath.Join(HomePath, "conf/defaults.ini")) {
		return
	}

	// try down one path
	if pathExists(filepath.Join(HomePath, "../conf/defaults.ini")) {
		HomePath = filepath.Join(HomePath, "../")
	}
}

func (cfg *Cfg) loadConfiguration(args *CommandLineArgs) (*ini.File, error) {
	var err error

	// load config defaults
	defaultConfigFile := path.Join(HomePath, "conf/defaults.ini")
	configFiles = append(configFiles, defaultConfigFile)

	// check if config file exists
	if _, err := os.Stat(defaultConfigFile); os.IsNotExist(err) {
		fmt.Println("evening-server Init Failed: Could not find config defaults, make sure homepath command line parameter is set or working directory is homepath")
		os.Exit(1)
	}

	// load defaults
	parsedFile, err := ini.Load(defaultConfigFile)
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to parse defaults.ini, %v", err))
		os.Exit(1)
		return nil, err
	}

	parsedFile.BlockMode = false

	err = cfg.initLogging(parsedFile)
	if err != nil {
		return nil, err
	}

	return parsedFile, err
}

//Load : Initialization for Cfg from command line
func (cfg *Cfg) Load(args *CommandLineArgs) error {
	setHomePath(args)

	ApplicationName = APPName

	iniFile, err := cfg.loadConfiguration(args)
	if err != nil {
		return err
	}

	cfg.Raw = iniFile
	Raw = cfg.Raw

	server := iniFile.Section("server")

	AppURL, AppSubURL, err = parseAppURLAndSubURL(server)
	if err != nil {
		return err
	}
	ServeFromSubPath = server.Key("serve_from_sub_path").MustBool(false)

	cfg.AppURL = AppURL
	cfg.AppSubURL = AppSubURL
	cfg.ServeFromSubPath = ServeFromSubPath

	Protocol = HTTP
	protocolStr, err := valueAsString(server, "protocol", "http")
	if err != nil {
		return err
	}
	if protocolStr == "https" {
		Protocol = HTTPS
		CertFile = server.Key("cert_file").String()
		KeyFile = server.Key("cert_key").String()
	}
	if protocolStr == "h2" {
		Protocol = HTTP2
		CertFile = server.Key("cert_file").String()
		KeyFile = server.Key("cert_key").String()
	}
	if protocolStr == "socket" {
		Protocol = SOCKET
		SocketPath = server.Key("socket").String()
	}

	HTTPAddr, err = valueAsString(server, "http_addr", DefaultHTTPAddr)
	if err != nil {
		return err
	}
	HTTPPort, err = valueAsString(server, "http_port", "3000")
	if err != nil {
		return err
	}
	staticRoot, err := valueAsString(server, "static_root_path", "")
	if err != nil {
		return err
	}
	StaticRootPath = makeAbsolute(staticRoot, HomePath)

	// read security settings
	security := iniFile.Section("security")
	SecretKey, err = valueAsString(security, "secret_key", "")

	samesiteString, err := valueAsString(security, "cookie_samesite", "lax")
	if err != nil {
		return err
	}

	if samesiteString == "disabled" {
		CookieSameSiteDisabled = true
		cfg.CookieSameSiteDisabled = CookieSameSiteDisabled
	} else {
		validSameSiteValues := map[string]http.SameSite{
			"lax":    http.SameSiteLaxMode,
			"strict": http.SameSiteStrictMode,
			"none":   http.SameSiteNoneMode,
		}

		if samesite, ok := validSameSiteValues[samesiteString]; ok {
			CookieSameSiteMode = samesite
			cfg.CookieSameSiteMode = CookieSameSiteMode
		} else {
			CookieSameSiteMode = http.SameSiteLaxMode
			cfg.CookieSameSiteMode = CookieSameSiteMode
		}
	}
	// auth
	auth := iniFile.Section("auth")

	LoginCookieName, err = valueAsString(auth, "login_cookie_name", "evening_session")
	cfg.LoginCookieName = LoginCookieName
	if err != nil {
		return err
	}
	cfg.LoginMaxInactiveLifetimeDays = auth.Key("login_maximum_inactive_lifetime_days").MustInt(7)

	LoginMaxLifetimeDays = auth.Key("login_maximum_lifetime_days").MustInt(30)
	cfg.LoginMaxLifetimeDays = LoginMaxLifetimeDays

	cfg.TokenRotationIntervalMinutes = auth.Key("token_rotation_interval_minutes").MustInt(10)
	if cfg.TokenRotationIntervalMinutes < 2 {
		cfg.TokenRotationIntervalMinutes = 2
	}
	SignoutRedirectUrl, err = valueAsString(auth, "signout_redirect_url", "")

	return nil
}

// LogConfigSources : log config source file path
func (cfg *Cfg) LogConfigSources() {
	var text bytes.Buffer

	for _, file := range configFiles {
		cfg.Logger.Info("Config loaded from", "file", file)
	}

	if len(appliedCommandLineProperties) > 0 {
		for _, prop := range appliedCommandLineProperties {
			cfg.Logger.Info("Config overridden from command line", "arg", prop)
		}
	}

	if len(appliedEnvOverrides) > 0 {
		text.WriteString("\tEnvironment variables used:\n")
		for _, prop := range appliedEnvOverrides {
			cfg.Logger.Info("Config overridden from Environment variable", "var", prop)
		}
	}

	cfg.Logger.Info("Path Home", "path", HomePath)
	cfg.Logger.Info("Path Data", "path", cfg.DataPath)
	cfg.Logger.Info("Path Logs", "path", cfg.LogsPath)
	cfg.Logger.Info("Path Provisioning", "path", cfg.ProvisioningPath)
	cfg.Logger.Info("App mode " + Env)
}

func (cfg *Cfg) initLogging(file *ini.File) error {
	logModeStr, err := valueAsString(file.Section("log"), "mode", "console")
	if err != nil {
		return err
	}
	// split on comma
	logModes := strings.Split(logModeStr, ",")
	// also try space
	if len(logModes) == 1 {
		logModes = strings.Split(logModeStr, " ")
	}
	logsPath, err := valueAsString(file.Section("paths"), "logs", "")
	if err != nil {
		return err
	}
	cfg.LogsPath = makeAbsolute(logsPath, HomePath)
	return log.ReadLoggingConfig(logModes, cfg.LogsPath, file)
}

// valueAsString : get key value from config file's section
func valueAsString(section *ini.Section, keyName string, defaultValue string) (value string, err error) {
	defer func() {
		if error := recover(); error != nil {
			err = errors.New("Invalid value for key '" + keyName + "' in configuration file")
		}
	}()

	return section.Key(keyName).MustString(defaultValue), nil
}

// makeAbsolute : make sure the path is absolute path
func makeAbsolute(path string, root string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(root, path)
}
