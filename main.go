package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rdkvx/desafioTecnico/v2/routes"
	"github.com/rdkvx/desafioTecnico/v2/services"
	"github.com/rdkvx/desafioTecnico/v2/utils"
)

var (
	err error
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	routes.AddRoutes(router)

	server := &http.Server{
		Addr:    utils.HttpPort,
		Handler: router,
	}

	// starting the validator interface
	services.PasswordValidator = services.NewPasswordValidator()
	
	fmt.Printf(utils.ApiStartMsg, utils.HttpPort)

	//skip a line
	fmt.Println()

	// starting API server
	if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Errorf(utils.ApiStartErr, err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Errorf(utils.ApiForcedShutdown, err)
		os.Exit(1)
	}

	fmt.Println(utils.APIStopped)
}
