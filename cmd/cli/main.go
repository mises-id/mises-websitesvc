package main

import (
	"flag"
	"fmt"

	"context"
	"time"

	//_ "github.com/mises-id/mises-websitesvc/config"
	"github.com/mises-id/mises-websitesvc/lib/db"

	// This Service
	"github.com/mises-id/mises-websitesvc/handlers"
	"github.com/mises-id/mises-websitesvc/svc/server"
)

func main() {
	// Update addresses if they have been overwritten by flags
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	fmt.Println("setup mongo...")
	db.SetupMongo(ctx)
	//models.EnsureIndex()

	cfg := server.DefaultConfig
	cfg = handlers.SetConfig(cfg)

	server.Run(cfg)
}
