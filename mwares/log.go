package mwares

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"gobpframe/config"
	"gobpframe/utils/helper"
	"gobpframe/utils/logger"
)

const (
	red    = "\033[%d;31m%v\033[0m"
	blue   = "\033[%d;34m%v\033[0m"
	green  = "\033[%d;32m%v\033[0m"
	yellow = "\033[%d;33m%v\033[0m"
	purple = "\033[%d;35m%v\033[0m"
)

func coloredStatus(status int) string {
	switch {
	case 200 <= status && status < 299:
		return fmt.Sprintf(green, 0, status)
	case 300 <= status && status < 399:
		return fmt.Sprintf(blue, 0, status)
	case 400 <= status && status < 499:
		return fmt.Sprintf(yellow, 0, status)
	case 500 <= status && status < 599:
		return fmt.Sprintf(red, 0, status)
	}
	return fmt.Sprint(purple, 0, status)
}

func coloredMethod(method string) string {
	switch {
	case method == "GET":
		return fmt.Sprintf(blue, 1, method)
	case method == "POST":
		return fmt.Sprintf(green, 1, method)
	case method == "PUT":
		return fmt.Sprintf(yellow, 1, method)
	case method == "DELETE":
		return fmt.Sprintf(red, 1, method)
	}
	return fmt.Sprint(purple, 1, method)
}

func Logger(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	endTime := time.Now()

	status := fmt.Sprintf("%d", c.Writer.Status())
	method := strings.ToUpper(c.Request.Method)

	// logging to stdout
	if !helper.IsStdoutRedirectToFile() && !config.GetBool("Logger.NoColor") {
		status = coloredStatus(c.Writer.Status())
		method = coloredMethod(method)
	}

	logger.Infof("| %s | %s | %s | %s | %s",
		status,
		method,
		endTime.Sub(startTime).String(),
		c.ClientIP(),
		c.Request.RequestURI,
		// c.Request.UserAgent(),
	)
}
