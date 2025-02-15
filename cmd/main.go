package main

import (
	"github.com/codeboris/avito-shop/internal/app"
	"log"
)

func main() {
	application, err := app.New()
	if err != nil {
		log.Fatalf("Не удалось инициализировать приложение: %v", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("Не удалось запустить приложение: %v", err)
	}
}
