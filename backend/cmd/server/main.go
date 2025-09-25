// backend/cmd/server/main.go
package main

import (
	"log"
	"os"

	"github.com/yourname/csvxlchart/backend/internal/config"
	"github.com/yourname/csvxlchart/backend/internal/router"
)

func main() {
	cfg := config.Load()
	r := router.New(cfg)

	addr := ":" + cfg.Port
	log.Printf("listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Println("server exited:", err)
		os.Exit(1)
	}
}
