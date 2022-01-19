package handler

import (
	"net/http"
)

func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: pongを返してみる
		if _, err := w.Write([]byte("")); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte("pong"))
		w.WriteHeader(http.StatusOK)
	}
}
