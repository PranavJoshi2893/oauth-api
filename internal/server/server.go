package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PranavJoshi2893/oauth-api/internal/config"
)

type Server struct {
	httpServer *http.Server
	db         *sql.DB
}

func New(config *config.Config, db *sql.DB) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: config.ServerPort,
		},
		db: db,
	}
}

func (s *Server) RunWithGracefulShutdown() error {

	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)
		log.Printf("Server is running on port %v\n", s.httpServer.Addr)
		err := s.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		return fmt.Errorf("server error: %v", err)
	case <-quit:
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		log.Println("Shutting down")
		if err := s.httpServer.Shutdown(ctx); err != nil {
			return fmt.Errorf("forced shutdown: %v", err)
		}

		log.Println("shutdown")
	}
	return nil

}
