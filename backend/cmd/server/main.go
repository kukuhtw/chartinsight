// backend/cmd/server/main.go
/*
=============================================================================
Project : ChartInsight — Upload CSV/XLSX → Interactive Charts + AI Insights.
Author  : Kukuh Tripamungkas Wicaksono (Kukuh TW)
Email   : kukuhtw@gmail.com
WhatsApp: https://wa.me/628129893706
LinkedIn: https://id.linkedin.com/in/kukuhtw
=============================================================================
*/

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
