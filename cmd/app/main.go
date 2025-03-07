package main

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/mrzhov/course-app/internal/config"
	"github.com/mrzhov/course-app/internal/modules"
)

func main() {
	c := config.Init()
	modules.Init(c.Echo, c.DB)

	if err := c.Echo.Start(c.Env.PORT); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
