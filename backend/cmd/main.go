package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/54m/api_gen-example/backend/interfaces"
	"github.com/54m/api_gen-example/backend/pkg/validator"
	"github.com/54m/echo-routing/output"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initMiddleware() interfaces.MiddlewareList {
	middlewareList := make([]*interfaces.MiddlewareSet, 0)

	middlewareSet := &interfaces.MiddlewareSet{
		Path: "/api/user/",
		MiddlewareFunc: []echo.MiddlewareFunc{
			func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
				return func(context echo.Context) error {
					// /api/user/以下の全てのハンドラーにこの関数を適用
					return nil
				}
			},
		},
	}

	return append(middlewareList, middlewareSet)
}

func initEchoSetting(e *echo.Echo) {
	e.Debug = true
	e.HideBanner = true
	e.Validator = validator.NewValidator()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func main() {
	e := echo.New()

	initEchoSetting(e)

	var (
		middlewareList = initMiddleware()
		repo           = initRepositories()
		props          = initControllerProps(repo)
	)

	interfaces.Bootstrap(props, e, middlewareList, os.Stdout)

	go func() {
		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 10 seconds.
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	output.Do(e.Routes())

	if err := e.Start(":8888"); err != nil {
		panic(err)
	}
}
