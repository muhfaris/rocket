package router

import (
	"context"
	"fmt"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/samplepg/config"
	handlerv1 "github.com/muhfaris/rocket/examples/samplepg/internal/adapter/inbound/rest/router/v1/handler"
	"github.com/muhfaris/rocket/examples/samplepg/internal/adapter/inbound/rest/router/v1/middleware"
	portadapter "github.com/muhfaris/rocket/examples/samplepg/internal/core/port/inbound/adapter"

	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/muhfaris/rocket/examples/samplepg/internal/adapter/inbound/rest/router/group"
	"github.com/muhfaris/rocket/examples/samplepg/internal/adapter/inbound/rest/router/v1/response"

	portregistry "github.com/muhfaris/rocket/examples/samplepg/internal/core/port/inbound/registry"
)

type Router struct {
	Client *fiber.App
	Port   int
}

func Init(port int, svcs portregistry.Service) portadapter.Rest {
	r := fiber.New(fiber.Config{
		AppName:                  "samplepg",
		EnableSplittingOnParsers: config.App.Fiber.EnableSplittingOnParsers,
		EnablePrintRoutes:        config.App.Fiber.EnablePrintRoutes,
		ErrorHandler:             response.Fail,
	})

	r.Use(requestid.New(requestid.Config{
		ContextKey: fiber.HeaderXRequestID,
	}))
	r.Use(middleware.Latency())

	// public
	// publicGroup := r.Group("/")
	// publicGroup.Get("/health", handlerv1.Health())

	h := handlerv1.New(svcs)
	group.V1(r, h)

	return &Router{
		Client: r,
		Port:   port,
	}
}

func (r *Router) Run() error {
	// gracefully shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var err error

	go func() {
		port := fmt.Sprintf(":%d", r.Port)
		slog.Info("Listening on port", "port", port)
		if err := r.Client.Listen(port); err != nil {
			return
		}
	}()

	<-ctx.Done()
	err = r.Client.Shutdown()
	if err != nil {
		return fmt.Errorf("failed to shutdown: %v", err)
	}

	return err
}
