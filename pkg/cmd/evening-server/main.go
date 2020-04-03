package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/hooone/evening/pkg/log"
	"github.com/hooone/evening/pkg/setting"
)

var version = "0.0.1"
var commit = "SUNHY"
var buildBranch = "master"
var buildstamp string

func main() {
	//执行参数解析
	var (
		configFile = flag.String("config", "", "path to config file")
		homePath   = flag.String("homepath", "", "path to evening install/home path, defaults to working directory")
		pidFile    = flag.String("	", "", "path to pid file")

		v = flag.Bool("v", false, "prints current version and exits")
	)
	flag.Parse()

	//输入-v查询程序版本
	if *v {
		fmt.Printf("Version %s (commit: %s, branch: %s)\n", version, commit, buildBranch)
		os.Exit(0)
	}

	//设置编译信息
	buildstampInt64, _ := strconv.ParseInt(buildstamp, 10, 64)
	if buildstampInt64 == 0 {
		buildstampInt64 = time.Now().Unix()
	}

	setting.BuildVersion = version
	setting.BuildCommit = commit
	setting.BuildStamp = buildstampInt64
	setting.BuildBranch = buildBranch

	server := NewServer(*configFile, *homePath, *pidFile)

	go listenToSystemSignals(server)

	err := server.Run()

	code := server.ExitCode(err)
	log.Close()

	os.Exit(code)
}

func listenToSystemSignals(server *Server) {
	signalChan := make(chan os.Signal, 1)
	sighupChan := make(chan os.Signal, 1)

	signal.Notify(sighupChan, syscall.SIGHUP)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-sighupChan:
			//终端断开时，logger重新初始化
			log.Reload()
		case sig := <-signalChan:
			//程序结束
			server.Shutdown(fmt.Sprintf("System signal: %s", sig))
		}
	}
}
