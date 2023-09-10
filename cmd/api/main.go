package main

import (
	"context"
	"github.com/POMBNK/shtrafovNetTestTask/internal/app"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a := app.NewApp(ctx)
	if err := a.Start(ctx); err != nil {
		log.Fatalln(err)
	}
}
