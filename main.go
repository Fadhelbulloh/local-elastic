package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Fadhelbulloh/local-elastic/route"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

		errs <- fmt.Errorf("%v", <-c)
	}()

	go func() {

		gin.SetMode(gin.DebugMode)
		router := gin.Default()
		router.Use(cors.Default())
		router.NoRoute(func(c *gin.Context) {
			c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found or wrong path"})
		})

		route.BasicService(router)
		if e := router.Run(":6969"); e != nil {
			errs <- e
		}
	}()

	log.Fatal(<-errs)
}
