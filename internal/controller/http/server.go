package http

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/v-351/storage-api/internal/config"
	"github.com/v-351/storage-api/internal/usecase"
)

type Server struct {
	server  *echo.Echo
	usecase usecase.Usecase
	ll      *slog.Logger
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func New(config config.Config, usecase usecase.Usecase) Server {
	s := Server{
		server:  echo.New(),
		usecase: usecase,
		ll: slog.New(slog.NewJSONHandler(os.Stderr,
			&slog.HandlerOptions{Level: slog.LevelDebug})),
	}

	s.server.Validator = &customValidator{validator: validator.New()}

	s.registerRoutes()
	if config.EnableDocs != 0 {
		s.registerSwaggerRoute()
	}

	return s
}

func (s *Server) Start(address string) error {
	return s.server.Start(address)
}

func (s *Server) Shoutdown() {
	s.server.Shutdown(context.Background())
}
