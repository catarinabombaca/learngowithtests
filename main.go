package main

import (
	"github.com/catarinabombaca/learngowithtests/countdown"
	"os"
	"time"
)

func main() {
	//greet.Greet(os.Stdout, "Elodie")
	//log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(greet.MyGreeterHandler)))

	sleeper := &countdown.ConfigurableSleeper{1 * time.Second, time.Sleep}
	countdown.Countdown(os.Stdout, sleeper)
}
