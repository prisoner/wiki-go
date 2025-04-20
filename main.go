package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"wiki-go/internal/config"
	"wiki-go/internal/handlers"
	"wiki-go/internal/routes"
	"wiki-go/internal/static"

	// Import goldext package for its initialization side effects
	_ "wiki-go/internal/goldext"
)

func main() {
	cfgFile := flag.String("c", config.ConfigFilePath, "Path to the configuration file")
	flag.Parse()

	// Load configuration
	cfg, err := config.LoadConfig(*cfgFile)
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Ensure the homepage exists
	if err := handlers.EnsureHomepageExists(cfg); err != nil {
		log.Fatal("Error creating homepage:", err)
	}

	// Ensure static assets exist in data directory
	if err := static.EnsureStaticAssetsExist(cfg.Wiki.RootDir); err != nil {
		log.Fatal("Error copying static assets:", err)
	}

	// Update handlers with config
	handlers.InitHandlers(cfg)

	// Setup all routes
	routes.SetupRoutes(cfg)

	// Start the server
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Server starting on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
