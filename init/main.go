package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"junction/db"
	"junction/network"
	"log"
	"net"
	"net/http"
)

var configPathFlag = flag.String("config", "../config/config.toml", "config file not found")

func Server(lc fx.Lifecycle) *gin.Engine {
	router := network.NewNetwork(lc)

	srv := &http.Server{Addr: ":8080", Handler: router} // define a web server

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr) // the web server starts listening on 8080
			if err != nil {
				fmt.Println("[Server] Failed to start HTTP Server at", srv.Addr)
				return err
			}
			go srv.Serve(ln) // process an incoming request in a go routine
			fmt.Println("[Server]Succeeded to start HTTP Server at", srv.Addr)
			return nil

		},
		OnStop: func(ctx context.Context) error {
			srv.Shutdown(ctx) // stop the web server
			fmt.Println("[Server] HTTP Server is stopped")
			return nil
		},
	})

	return router
}

func main() {
	client, err := db.NewDb(*configPathFlag)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app := fx.New(
		fx.Provide(Server),
		fx.Invoke(func(r *gin.Engine) {
			fmt.Print("Server is running on port 8080")
		}),
	)
	app.Run()
}
