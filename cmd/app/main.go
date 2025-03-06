package main

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/mrzhov/course-app/internal/common/config"
	"github.com/mrzhov/course-app/internal/common/routes"
)

func main() {
	c := config.Init()
	routes.Register(c.Echo, c.DB)

	if err := c.Echo.Start(c.Env.PORT); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
