package engine

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/fmotalleb/scrapper-go/config"
	"github.com/playwright-community/playwright-go"
)

func ExecuteConfig(config config.ExecutionConfig) (map[string]any, error) {
	vars := initializeVariables(config.Pipeline.Vars)
	pw, err := playwright.Run()
	if err != nil {
		slog.Error("could not start Playwright", slog.Any("err", err))
		return nil, fmt.Errorf("could not start Playwright: %v", err)
	}
	defer func() {
		if err := pw.Stop(); err != nil {
			slog.Warn("Failed to stop Playwright session", slog.Any("err", err))
		}
	}()

	slog.Info("Playwright initialized")

	browser, err := pw.Chromium.Launch(config.Pipeline.BrowserParams)
	if err != nil {
		slog.Error("could not launch browser", slog.Any("err", err))
		return nil, fmt.Errorf("could not launch browser: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		slog.Error("could not create page", slog.Any("err", err))
		return nil, fmt.Errorf("could not create page: %v", err)
	}

	result := make(map[string]any)
	for _, step := range config.Pipeline.Steps {
		slog.Debug("Executing step", slog.Any("step", step))
		if err := executeStep(page, step, vars, result); err != nil {
			slog.Error("Error executing step", slog.Any("err", err), slog.Any("step", step))
			return result, fmt.Errorf("Error executing step: %v, step: %v", err, step)
		}
	}

	slog.Info("Execution finished", slog.Any("vars_snapshot", vars.Snapshot()), slog.Any("result", result))

	if config.Pipeline.KeepRunning != "" {
		sleepTime, err := time.ParseDuration(config.Pipeline.KeepRunning)
		if err != nil {
			slog.Error("Cannot parse duration in KeepRunning", slog.Any("err", err))
			return result, fmt.Errorf("cannot parse given duration in keep running: %s", config.Pipeline.KeepRunning)
		}
		slog.Info("Sleeping for duration", slog.Duration("duration", sleepTime))
		time.Sleep(sleepTime)
	}

	return result, nil
}
