package main

import (
	"context"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"zgo/pkg/api"
	"zgo/pkg/auth"
	"zgo/pkg/hasher"
	"zgo/pkg/persistence"
)

func mkRouter() *chi.Mux {
	r := chi.NewRouter()

	authenticator := auth.NewTokenAuthenticator(
		persistence.NewMemoryUserPersistence(),
		hasher.NewBCryptHasher(),
		auth.NewJwtTokenGenerator("secret", time.Hour),
	)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		api.SendJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	r.Post("/tokens", auth.IssueToken(authenticator))

	return r
}

func main() {
	r := mkRouter()

	srv := http.Server{
		Addr:    ":8080",
		Handler: chi.ServerBaseContext(context.TODO(), r),
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		srv.Shutdown(ctx)
	}()

	log.Fatal(srv.ListenAndServe())
}
