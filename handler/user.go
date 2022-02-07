package handler

import (
	"encoding/json"
	"io"
	"path/filepath"
	"strconv"
	"strings"

	"net/http"

	"github.com/karamaru-alpha/ca-tech-accel/model"
)

// Add 新規ユーザ登録
func Add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		var req SignupRequest
		if err := json.Unmarshal(body, &req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		user, err := model.Create(req.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		resp, err := json.Marshal(
			SignupResponse{
				ID:   user.ID,
				Name: user.Name,
			},
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if _, err := w.Write(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// List ユーザー一覧の取得
func List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := model.List()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		res, err := json.Marshal(
			UserListResponse{
				Users: users,
			},
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if _, err := w.Write(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// Find ユーザー取得
func Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sub := strings.TrimPrefix(r.URL.Path, "/user/find")
		_, id := filepath.Split(sub)
		userID, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		user, err := model.Find(userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		res, err := json.Marshal(
			UserFindResponse{
				ID:   user.ID,
				Name: user.Name,
			},
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if _, err := w.Write(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// Update ユーザー更新
func Update() http.HandlerFunc {
	// 関数を返す
	return func(w http.ResponseWriter, r *http.Request) {
		// 一括読み込み
		body, err := io.ReadAll(r.Body)
		// エラー処理
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		var req UserUpdateRequest
		// json.Marshalは構造体をjsonに変換する
		if err := json.Unmarshal(body, &req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		// idを取得
		sub := strings.TrimPrefix(r.URL.Path, "/user/update")
		_, id := filepath.Split(sub)
		userID, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		model.Update(userID, req.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

type (
	SignupRequest struct {
		Name string `json:"name"`
	}
	SignupResponse struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	UserListResponse struct {
		Users []model.User `json:"users"`
	}

	UserFindResponse struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	UserUpdateRequest struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)
