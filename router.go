package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	"fmt"
)

var router *gin.Engine

func route() {
	// Create our web server
	router = gin.Default()

	// Add our static file routes
	router.Static("/assets", "./public/assets/")

	// Figure out host:port
	addr := fmt.Sprintf("%s:%d", *host, *port)

	// Create our server based off of net/http so we can
	// control it more
	s := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Listen and serve
	s.ListenAndServe()
}
