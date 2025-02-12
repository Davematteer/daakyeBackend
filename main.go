package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte("Bitch we up"))

}

func main() {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	fmt.Println("This is the ussd app")
	//http.HandleFunc("/", ussd_callback)
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
