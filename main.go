package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndpoint)).Methods("GET")

	log.Println("Listen oon port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Signup invoked")
	w.Write([]byte("succesfully called signup"))
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login invoked")
	w.Write([]byte("succesfully called login"))
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protected invoked")
	w.Write([]byte("succesfully called protectedEndpoint"))
}

// TokenVerifyMiddleWare function
func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("TokenVerifyMiddleWare invoked")
	return nil
}
