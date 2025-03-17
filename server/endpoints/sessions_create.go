package endpoints

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/fmotalleb/scrapper-go/config"
	"github.com/fmotalleb/scrapper-go/session"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

func init() {
	registerEndpoint(
		endpoint{
			method:  "POST",
			path:    "/sessions",
			handler: sessionsCreate,
		},
	)
}

func sessionsCreate(c echo.Context) error {
	cfgMap := make(map[string]any)
	timeoutStr := c.QueryParam("timeout")
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		timeout = time.Minute * 5
	}

	err = json.NewDecoder(c.Request().Body).Decode(&cfgMap)
	if err != nil {
		slog.Error("failed to body", slog.Any("err", err))
		return c.String(http.StatusBadRequest, "cannot unmarshal the given json body")
	}
	var cfg config.ExecutionConfig
	err = mapstructure.Decode(cfgMap, &cfg)
	if err != nil {
		slog.Error("failed to read config from body", slog.Any("err", err))
		return c.String(http.StatusBadRequest, "cannot unmarshal the given json body")
	}
	res, err := session.NewSession(cfg, timeout)
	if err != nil {
		slog.Error("failed to execute config", slog.Any("err", err))
		return c.String(http.StatusBadRequest, "failed to execute config. make sure the config is compatible with service")
	}
	return c.JSON(
		http.StatusOK,
		map[string]any{
			"id":      res.ID,
			"timeout": timeout,
		},
	)
}
