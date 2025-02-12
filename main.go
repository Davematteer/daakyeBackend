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

func ussd_callback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	// the stuff being sent to my backend from africastalking

	session_id := r.FormValue("sessionId")
	service_code := r.FormValue("serviceCode")
	phone_number := r.FormValue("phoneNumber")
	text := r.FormValue("text")

	_ = fmt.Sprint("%s,%s,%s", session_id, service_code, phone_number)

	if len(text) == 0 {

	}
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
