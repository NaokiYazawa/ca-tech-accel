package main

import (
	"log"
	"net/http"

	"github.com/karamaru-alpha/ca-tech-accel/handler"
)

func main() {
	// TODO: echo を導入してみよう
	h := http.NewServeMux()
	h.Handle("/ping", handler.Ping())

	h.Handle("/user/add", handler.Add())
	h.Handle("/user/list", handler.List())
	h.Handle("/user/find/", handler.Find()) // ex. /user/find/1
	// TODO: ユーザー名のアップデート処理を追加しよう
	h.Handle("/user/update/", handler.Update())

	srv := http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	// TODO: graceful shutdownに対応しよう
	log.Println("starting server on port 8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to listen and serve: %s", err)
	}
}
