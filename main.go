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

	_ = fmt.Sprintf("%s,%s,%s", session_id, service_code, phone_number)

	if len(text) == 0 {
		w.Write([]byte("Nigga what do you wanna see\n1. Account details\n2. Make a deposit?\n3. Check insurance plans"))
		return
	} else {
		switch text {
		case "1":
			w.Write([]byte("CON Select an account option\n1. Account number\n2. Account Balance"))
			return
		case "2":
			w.Write([]byte(fmt.Sprintf("END Your Phone Number is %s", phone_number)))
			return
		case "1*1":
			w.Write([]byte("END Your Account Number is ACC1001"))
			return
		case "1*2":
			w.Write([]byte("END Your Balance is NGN 20,000"))
			return
		default:
			w.Write([]byte("END Invalid input"))
			return
		}
	}
}

func main() {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	fmt.Println("This is the ussd app")
	http.HandleFunc("/", ussd_callback)
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
