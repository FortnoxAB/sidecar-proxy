package main

import (
	"flag"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/jonaz/gograce"
	"github.com/sirupsen/logrus"
)

func main() {
	listenAddress := flag.String("listen-addr", os.Getenv("LISTEN_ADDRESS"), "Port to listen on")
	forwardAddress := flag.String("proxy-addr", os.Getenv("PROXY_ADDRESS"), "IP /w port to forward to")
	logLevel := flag.String("log-level", os.Getenv("LOG_LEVEL"), "Set proxy's log level")
	// flag.Var(&metricNames, "metric-name", "Set which metrics to collect")
	flag.Parse()

	if *logLevel != "" {
		lvl, err := logrus.ParseLevel(*logLevel)
		if err != nil {
			logrus.Error(err)
			return
		}
		logrus.SetLevel(lvl)
	}

	forwardURL, err := url.Parse(*forwardAddress)
	if err != nil {
		logrus.Errorf("invalid proxy-addr: %s", err)
		return
	}

	proxy := NewSingleHostReverseProxy(forwardURL)
	srv, shutdown := gograce.NewServerWithTimeout(3 * time.Second)
	srv.Handler = proxy
	srv.Addr = *listenAddress

	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logrus.Error(err)
	}
	<-shutdown
}

func NewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	// targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
		req.URL.RawPath = target.RawPath
		req.Host = target.Host
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
		logrus.Debug(req.URL.String())
	}
	return &httputil.ReverseProxy{Director: director}
}
