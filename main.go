package main

import (
	"github.com/catarinabombaca/learngowithtests/greet"
	"log"
	"net/http"
	"os"
)

func main() {
	greet.Greet(os.Stdout, "Elodie")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(greet.MyGreeterHandler)))
}
