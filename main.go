package main

import (
	"api/app/stringtest"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {

	// 配置
	var (
		httpAddr = flag.String("http.addr", ":8088", "HTTP listen address")
	)
	flag.Parse()

	// 日志
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	var s stringtest.Service
	{
		s = stringtest.NewStringService()
		s = stringtest.LoggingMiddleware(logger)(s)
	}

	var h http.Handler
	{
		h = stringtest.MakeHTTPHandler(s, log.With(logger, "component", "HTTP"))
	}
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	logger.Log("exit", <-errs)
}
