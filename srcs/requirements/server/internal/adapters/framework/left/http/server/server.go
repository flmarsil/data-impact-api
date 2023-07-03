package server

import (
	"context"
	"data_impact/srcs/requirements/server/internal/adapters/framework/left/http/server/api"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func HttpLauncher(ms *api.MicroserviceServer) {
	router := NewRouter(ms)

	server := &http.Server{
		Addr:         ":9000",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT)

	done := make(chan struct{}, 1)

	// stop the server properly
	go func() {
		<-s
		fmt.Println("Stopping the server ...")
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		err := server.Shutdown(ctx)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		close(done)
	}()

	fmt.Println("Starting the server. Listening on http://localhost:9000")

	server.ListenAndServe()
	<-done

	fmt.Println("Server has been stopped.")
}
