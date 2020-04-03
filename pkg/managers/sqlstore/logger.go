package sqlstore

import (
	"fmt"

	glog "github.com/hooone/evening/pkg/log"

	"github.com/go-xorm/core"
)

//XormLogger logger for xorm
type XormLogger struct {
	gLog    glog.Logger
	level   glog.Lvl
	showSQL bool
}

//NewXormLogger wrap logger to XormLogger
func NewXormLogger(level glog.Lvl, gLog glog.Logger) *XormLogger {
	return &XormLogger{
		gLog:    gLog,
		level:   level,
		showSQL: true,
	}
}

// Error implement core.ILogger
func (s *XormLogger) Error(v ...interface{}) {
	if s.level <= glog.LvlError {
		s.gLog.Error(fmt.Sprint(v...))
	}
}

// Errorf implement core.ILogger
func (s *XormLogger) Errorf(format string, v ...interface{}) {
	if s.level <= glog.LvlError {
		s.gLog.Error(fmt.Sprintf(format, v...))
	}
}

// Debug implement core.ILogger
func (s *XormLogger) Debug(v ...interface{}) {
	if s.level <= glog.LvlDebug {
		s.gLog.Debug(fmt.Sprint(v...))
	}
}

// Debugf implement core.ILogger
func (s *XormLogger) Debugf(format string, v ...interface{}) {
	if s.level <= glog.LvlDebug {
		s.gLog.Debug(fmt.Sprintf(format, v...))
	}
}

// Info implement core.ILogger
func (s *XormLogger) Info(v ...interface{}) {
	if s.level <= glog.LvlInfo {
		s.gLog.Info(fmt.Sprint(v...))
	}
}

// Infof implement core.ILogger
func (s *XormLogger) Infof(format string, v ...interface{}) {
	if s.level <= glog.LvlInfo {
		s.gLog.Info(fmt.Sprintf(format, v...))
	}
}

// Warn implement core.ILogger
func (s *XormLogger) Warn(v ...interface{}) {
	if s.level <= glog.LvlWarn {
		s.gLog.Warn(fmt.Sprint(v...))
	}
}

// Warnf implement core.ILogger
func (s *XormLogger) Warnf(format string, v ...interface{}) {
	if s.level <= glog.LvlWarn {
		s.gLog.Warn(fmt.Sprintf(format, v...))
	}
}

// Level implement core.ILogger
func (s *XormLogger) Level() core.LogLevel {
	switch s.level {
	case glog.LvlError:
		return core.LOG_ERR
	case glog.LvlWarn:
		return core.LOG_WARNING
	case glog.LvlInfo:
		return core.LOG_INFO
	case glog.LvlDebug:
		return core.LOG_DEBUG
	default:
		return core.LOG_ERR
	}
}

// SetLevel implement core.ILogger
func (s *XormLogger) SetLevel(l core.LogLevel) {
}

// ShowSQL implement core.ILogger
func (s *XormLogger) ShowSQL(show ...bool) {
	s.gLog.Error("ShowSQL", "show", "show")
	if len(show) == 0 {
		s.showSQL = true
		return
	}
	s.showSQL = show[0]
}

// IsShowSQL implement core.ILogger
func (s *XormLogger) IsShowSQL() bool {
	return s.showSQL
}
