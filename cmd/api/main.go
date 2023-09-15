package main

import (
	"context"
	"github.com/POMBNK/shtrafovNetTestTask/internal/app"
	"github.com/POMBNK/shtrafovNetTestTask/pkg/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx = logger.ContextWithLogger(ctx, logger.NewLogger())
	a := app.NewApp(ctx)

	logger.L(ctx).Info("Starting application...")

	if err := a.Start(ctx); err != nil {
		logger.WithError(ctx, err).Fatal("app.Run")
	}
}
