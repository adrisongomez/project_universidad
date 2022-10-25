package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:generate go run github.com/prisma/prisma-client-go generate

func main() {
	const (
		port = ":3000"
	)
	e := echo.New()
	if err := e.Start(port); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
