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
)

var (
	err error
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: cors.DefaultConfig().AllowMethods,
		AllowHeaders: cors.DefaultConfig().AllowHeaders,
		MaxAge:       12 * time.Hour,
	}))

	routes.AddRoutes(router)

	httpPort := ":8080"

	server := &http.Server{
		Addr:    httpPort,
		Handler: router,
	}
/* 
	go func() {
		fmt.Printf("API sendo ouvida na porta %s", httpPort)
		if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Errorf("nao foi possivel iniciar o servidor: %v", err)
			os.Exit(1)
		}
	}() */


	// iniciando o servidor da API	
	fmt.Printf("API sendo ouvida na porta %s", httpPort)
	if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Errorf("nao foi possivel iniciar o servidor: %v", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Errorf("o servidor foi forçado a parar %s", err)
		os.Exit(1)
	}

	fmt.Println("servidor parado com sucesso")
}
