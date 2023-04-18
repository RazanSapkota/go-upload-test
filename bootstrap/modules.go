package bootstrap

import (
	"context"
	"example/go-api/controllers"
	"example/go-api/infrastructure"
	"example/go-api/lib"
	"example/go-api/repository"
	"example/go-api/routes"
	"example/go-api/service"
	"fmt"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	infrastructure.Module,
	lib.Module,
	repository.Module,
	controllers.Module,
	routes.Module,
	service.Module,
	fx.Invoke(registerHooks),
)

func registerHooks(lifecycle fx.Lifecycle, h lib.RequestHandler, userRoute routes.UserRoutes) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Println("Starting application in :5000")
				//userRoute.Setup()
				if err := sentry.Init(sentry.ClientOptions{
					Dsn: "https://eee42c8e696245eba167bf49b2b4ff36@o4504236202000384.ingest.sentry.io/4504236203900928",
				}); err != nil {
					fmt.Printf("Sentry initialization failed: %v\n", err)
				}
				h.Gin.Use(sentrygin.New(sentrygin.Options{Repanic: true}))
				h.Gin.GET("/foo", func(ctx *gin.Context) {
					panic("Error Occured")
				})
				go h.Gin.Run(":5000")
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
