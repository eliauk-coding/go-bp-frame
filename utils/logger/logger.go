package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"gobpframe/config"
	"gobpframe/utils/helper"

	uuid "github.com/satori/go.uuid"
)

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var (
	logMinLevel   = DebugLevel
	logTimeFormat = "01/02/06 15:04:05.000"
)

var strLevelMap = map[string]Level{
	"all":     DebugLevel,
	"debug":   DebugLevel,
	"info":    InfoLevel,
	"warn":    WarnLevel,
	"warning": WarnLevel,
	"fatal":   FatalLevel,
	"error":   ErrorLevel,
}

var levelColorMap = map[Level]string{
	DebugLevel: "\033[1;36m%v\033[0m",
	InfoLevel:  "\033[1;34m%v\033[0m",
	WarnLevel:  "\033[1;33m%v\033[0m",
	ErrorLevel: "\033[1;31m%v\033[0m",
	FatalLevel: "\033[1;35m%v\033[0m",
}

var levelFileMap = map[Level]string{
	DebugLevel: "debug",
	InfoLevel:  "info",
	WarnLevel:  "warning",
	ErrorLevel: "error",
	FatalLevel: "fatal",
}

var levelStrMap = map[Level]string{
	DebugLevel: "DBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERRO",
	FatalLevel: "FATA",
}

func init() {
	level := "info"
	if helper.IsDebug() {
		level = "all"
	} else {
		lvl := config.GetStr("Logger.Level")
		if lvl != "" {
			level = lvl
		}
	}
	SetLevel(level)
}

func SetLevel(str string) {
	level, ok := strLevelMap[str]
	if !ok {
		level = InfoLevel
	}
	logMinLevel = level
}

func Debug(args ...interface{}) {
	logPrint(DebugLevel, args...)
}

func Debugf(tpl string, args ...interface{}) {
	logPrintf(DebugLevel, tpl, args...)
}

func Info(args ...interface{}) {
	logPrint(InfoLevel, args...)
}

func Infof(tpl string, args ...interface{}) {
	logPrintf(InfoLevel, tpl, args...)
}

func Warn(args ...interface{}) {
	logPrint(WarnLevel, args...)
}

func Warnf(tpl string, args ...interface{}) {
	logPrintf(WarnLevel, tpl, args...)
}

func Error(args ...interface{}) {
	logPrint(WarnLevel, args...)
}

func Errorf(tpl string, args ...interface{}) {
	logPrintf(WarnLevel, tpl, args...)
}

func Fatal(args ...interface{}) {
	logPrint(FatalLevel, args...)
	panic("panic - by fatal error")
}

func Fatalf(tpl string, args ...interface{}) {
	logPrintf(FatalLevel, tpl, args...)
	panic("panic - by fatal error")
}

func logPrintf(lvl Level, tpl string, args ...interface{}) {
	if lvl < logMinLevel {
		return
	}

	// logging to stdout
	if !helper.IsStdoutRedirectToFile() {
		lvlStr := " [" + levelStrMap[lvl] + "] "
		nowStr := time.Now().Format(logTimeFormat)
		if !config.GetBool("Logger.NoColor") {
			white := "\033[0;97m%v\033[0m"
			lvlStr = fmt.Sprintf(levelColorMap[lvl], lvlStr)
			nowStr = fmt.Sprintf(white, nowStr)
		}
		tpl = nowStr + lvlStr + tpl + "\n"
		fmt.Printf(tpl, args...)
		return
	}

	// redirect to file
	logData := map[string]interface{}{
		"time":     time.Now().Format(time.RFC3339),
		"app_name": config.GetStr("Server.AppName"),
		"level":    levelFileMap[lvl],
		"msg":      fmt.Sprintf(tpl, args...),
	}

	if lvl == DebugLevel {
		pc, filePath, line, _ := runtime.Caller(3)
		pcName := runtime.FuncForPC(pc).Name()
		_, file := filepath.Split(filePath)
		logData["caller"] = fmt.Sprintf("%s:%d %s", file, line, pcName)
		logData["traceid"] = uuid.NewV4().String()
	}

	logStr, _ := json.Marshal(logData)
	fmt.Fprintln(os.Stdout, string(logStr))
}

func logPrint(lvl Level, args ...interface{}) {
	logPrintf(lvl, strings.Trim(strings.Repeat("%v ", len(args)), " "), args...)
}
