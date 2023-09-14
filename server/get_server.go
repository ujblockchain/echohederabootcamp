package server

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
)

// we pass echo app instance as a parameter
func SetServer(app *echo.Echo) {
	server := &http2.Server{
		// set the number of concurrent streams
		// that can run at the same time
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		//set idle timeout
		IdleTimeout: 10 * time.Second,
	}

	// run server this will return an error is server does not run
	// we can then log the error
	if err := app.StartH2CServer(":3000", server); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
