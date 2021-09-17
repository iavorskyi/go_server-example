package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprint(firstVar + secondVar)

	err := ren.JSON(w, http.StatusOK, message)
	if err != nil {
		log.Println("Houston, we got a problem")
	}
}
