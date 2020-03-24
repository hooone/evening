// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hooone/evening/pkg/util"
	"github.com/hooone/evening/pkg/util/errutil"
	"github.com/inconshreveable/log15"
	"github.com/mattn/go-isatty"
	"gopkg.in/ini.v1"
)

var root log15.Logger
var loggersToClose []DisposableHandler
var loggersToReload []ReloadableHandler
var filters map[string]log15.Lvl

func init() {
	loggersToClose = make([]DisposableHandler, 0)
	loggersToReload = make([]ReloadableHandler, 0)
	root = log15.Root()
	filters = map[string]log15.Lvl{}
	root.SetHandler(log15.DiscardHandler())
}

//New a logger context
func New(logger string, ctx ...interface{}) Logger {
	params := append([]interface{}{"logger", logger}, ctx...)
	return root.New(params...)
}

//Trace is level 1 log, same as Debug.Use root context to print log.
func Trace(format string, v ...interface{}) {
	var message string
	if len(v) > 0 {
		message = fmt.Sprintf(format, v...)
	} else {
		message = format
	}

	root.Debug(message)
}

//Debug is level 1 log. Use root context to print log.
func Debug(format string, v ...interface{}) {
	var message string
	if len(v) > 0 {
		message = fmt.Sprintf(format, v...)
	} else {
		message = format
	}

	root.Debug(message)
}

//Info is level 2 log. Use root context to print log.
func Info(format string, v ...interface{}) {
	var message string
	if len(v) > 0 {
		message = fmt.Sprintf(format, v...)
	} else {
		message = format
	}

	root.Info(message)
}

//Warn is level 3 log. Use root context to print log.
func Warn(format string, v ...interface{}) {
	var message string
	if len(v) > 0 {
		message = fmt.Sprintf(format, v...)
	} else {
		message = format
	}

	root.Warn(message)
}

//Error is level 4 log. Use root context to print log.
func Error(skip int, format string, v ...interface{}) {
	root.Error(fmt.Sprintf(format, v...))
}

//Crit is level 5 log. Use root context to print log.
func Crit(skip int, format string, v ...interface{}) {
	root.Crit(fmt.Sprintf(format, v...))
}

//Fatal is level 5 log. Use root context to print log.And application will Exit
func Fatal(skip int, format string, v ...interface{}) {
	root.Crit(fmt.Sprintf(format, v...))
	Close()
	os.Exit(1)
}

//ReadLoggingConfig : set file path and log filter according config file and logger context
func ReadLoggingConfig(modes []string, logsPath string, cfg *ini.File) error {
	Close()

	defaultLevelName, _ := getLogLevelFromConfig("log", "info", cfg)
	defaultFilters := getFilters(util.SplitString(cfg.Section("log").Key("filters").String()))

	handlers := make([]log15.Handler, 0)

	for _, mode := range modes {
		mode = strings.TrimSpace(mode)
		sec, err := cfg.GetSection("log." + mode)
		if err != nil {
			root.Error("Unknown log mode", "mode", mode)
			return errutil.Wrapf(err, "failed to get config section log.%s", mode)
		}

		// Log level.
		_, level := getLogLevelFromConfig("log."+mode, defaultLevelName, cfg)
		modeFilters := getFilters(util.SplitString(sec.Key("filters").String()))
		format := getLogFormat(sec.Key("format").MustString(""))

		var handler log15.Handler

		// Generate log configuration.
		switch mode {
		case "console":
			handler = log15.StreamHandler(os.Stdout, format)
		case "file":
			fileName := sec.Key("file_name").MustString(filepath.Join(logsPath, "evening.log"))
			dpath := filepath.Dir(fileName)
			if err := os.MkdirAll(dpath, os.ModePerm); err != nil {
				root.Error("Failed to create directory", "dpath", dpath, "err", err)
				return errutil.Wrapf(err, "failed to create log directory %q", dpath)
			}
			fileHandler := NewFileWriter()
			fileHandler.Filename = fileName
			fileHandler.Format = format
			fileHandler.Rotate = sec.Key("log_rotate").MustBool(true)
			fileHandler.Maxlines = sec.Key("max_lines").MustInt(1000000)
			fileHandler.Maxsize = 1 << uint(sec.Key("max_size_shift").MustInt(28))
			fileHandler.Daily = sec.Key("daily_rotate").MustBool(true)
			fileHandler.Maxdays = sec.Key("max_days").MustInt64(7)
			if err := fileHandler.Init(); err != nil {
				root.Error("Failed to initialize file handler", "dpath", dpath, "err", err)
				return errutil.Wrapf(err, "failed to initialize file handler")
			}

			loggersToClose = append(loggersToClose, fileHandler)
			loggersToReload = append(loggersToReload, fileHandler)
			handler = fileHandler
		}
		if handler == nil {
			panic(fmt.Sprintf("Handler is uninitialized for mode %q", mode))
		}

		for key, value := range defaultFilters {
			if _, exist := modeFilters[key]; !exist {
				modeFilters[key] = value
			}
		}

		for key, value := range modeFilters {
			if _, exist := filters[key]; !exist {
				filters[key] = value
			}
		}

		handler = logFilterHandler(level, modeFilters, handler)
		handlers = append(handlers, handler)
	}

	root.SetHandler(log15.MultiHandler(handlers...))
	return nil
}

//Close all loggers
func Close() {
	for _, logger := range loggersToClose {
		logger.Close()
	}
	loggersToClose = make([]DisposableHandler, 0)
}

//Reload all loggers
func Reload() {
	for _, logger := range loggersToReload {
		logger.Reload()
	}
}

var logLevels = map[string]log15.Lvl{
	"trace":    log15.LvlDebug,
	"debug":    log15.LvlDebug,
	"info":     log15.LvlInfo,
	"warn":     log15.LvlWarn,
	"error":    log15.LvlError,
	"critical": log15.LvlCrit,
}

func getLogLevelFromConfig(key string, defaultName string, cfg *ini.File) (string, log15.Lvl) {
	levelName := cfg.Section(key).Key("level").MustString(defaultName)
	levelName = strings.ToLower(levelName)
	level := getLogLevelFromString(levelName)
	return levelName, level
}

func getLogLevelFromString(levelName string) log15.Lvl {
	level, ok := logLevels[levelName]

	if !ok {
		root.Error("Unknown log level", "level", levelName)
		return log15.LvlError
	}

	return level
}

func getFilters(filterStrArray []string) map[string]log15.Lvl {
	filterMap := make(map[string]log15.Lvl)

	for _, filterStr := range filterStrArray {
		parts := strings.Split(filterStr, ":")
		if len(parts) > 1 {
			filterMap[parts[0]] = getLogLevelFromString(parts[1])
		}
	}

	return filterMap
}

func getLogFormat(format string) log15.Format {
	switch format {
	case "console":
		if isatty.IsTerminal(os.Stdout.Fd()) {
			return log15.TerminalFormat()
		}
		return log15.LogfmtFormat()
	case "text":
		return log15.LogfmtFormat()
	case "json":
		return log15.JsonFormat()
	default:
		return log15.LogfmtFormat()
	}
}

func logFilterHandler(maxLevel log15.Lvl, filters map[string]log15.Lvl, h log15.Handler) log15.Handler {
	return log15.FilterHandler(func(r *log15.Record) (pass bool) {

		if len(filters) > 0 {
			for i := 0; i < len(r.Ctx); i += 2 {
				key, ok := r.Ctx[i].(string)
				if ok && key == "logger" {
					loggerName, strOk := r.Ctx[i+1].(string)
					if strOk {
						if filterLevel, ok := filters[loggerName]; ok {
							return r.Lvl <= filterLevel
						}
					}
				}
			}
		}

		return r.Lvl <= maxLevel
	}, h)
}
