package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	post "github.com/dcabral/gokit"
	posthttp "github.com/dcabral/gokit/api"
	"go.uber.org/zap"
)

const svcName = "blog"

func main() {

	logger, _ := zap.NewProduction()

	svc := post.NewService(logger)
	svc = posthttp.NewLoggingMiddleware(svc, logger)

	errs := make(chan error, 2)

	go startHTTPServer(svc, logger, errs)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err := <-errs
	logger.Error(fmt.Sprintf("Post service terminated: %s", err))
}

func startHTTPServer(svc post.PostService, logger *zap.Logger, errs chan error) {
	port := fmt.Sprintf(":%d", 8080)
	logger.Info(fmt.Sprintf("Post service started using http on port %d", 8080))
	errs <- http.ListenAndServe(port, posthttp.MakeHandler(svcName, svc))
}
