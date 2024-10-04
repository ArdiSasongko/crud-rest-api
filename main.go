package main

import (
	"golang-rest-api/helper"
	"golang-rest-api/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

// untuk menjalankan main.go, harus menyertakan wire_gen.go
// agar IntializedServer() terbaca
// go run main.go wire_gen.go
// atau go run .

func main() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
