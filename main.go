package main

import (
	"fmt"
	"net/http"
	"os"
)

func weatherHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/weather", weatherHandler)

	port := 7575
	fmt.Printf("Server is running on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
