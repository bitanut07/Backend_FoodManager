// package main

// import (
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	"goravel/bootstrap"

// 	"github.com/goravel/framework/facades"
// )

// func main() {
// 	bootstrap.Boot()

// 	// Get PORT from Render
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080" // fallback local
// 	}

// 	host := "0.0.0.0:" + port

// 	// Create channel to listen for OS signals
// 	quit := make(chan os.Signal)
// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

// 	// Start HTTP server
// 	go func() {
// 		if err := facades.Route().Run(host); err != nil {
// 			facades.Log().Errorf("Route Run error: %v", err)
// 		}
// 	}()

// 	// Graceful shutdown
// 	go func() {
// 		<-quit
// 		facades.Log().Info("Shutting down server...")

// 		if err := facades.Route().Shutdown(); err != nil {
// 			facades.Log().Errorf("Route Shutdown error: %v", err)
// 		}

// 		os.Exit(0)
// 	}()

// 	select {}
// }

package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/goravel/framework/facades"

	"goravel/bootstrap"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start http server by facades.Route().
	go func() {
		if err := facades.Route().Run(); err != nil {
			facades.Log().Errorf("Route Run error: %v", err)
		}
	}()

	// Listen for the OS signal
	go func() {
		<-quit
		if err := facades.Route().Shutdown(); err != nil {
			facades.Log().Errorf("Route Shutdown error: %v", err)
		}

		os.Exit(0)
	}()

	select {}
}
