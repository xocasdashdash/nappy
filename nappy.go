package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	numberOfNaps := 0
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	if len(os.Args) < 2 {
		log.Fatal("At least pass one parameter (a valid time.Duration string), ie 1s")
	}
	nappyTime, err := time.ParseDuration(os.Args[1])
	if err != nil {
		log.Fatalf("error parsing argument %#v", err)
	}
	ticker := time.NewTicker(nappyTime)
	go func() {
		for range ticker.C {
			numberOfNaps = numberOfNaps + 1
			fmt.Printf("I'm on my %d nap..maybe one more\n", numberOfNaps)
		}
	}()
	<-sigs
	ticker.Stop()
	fmt.Printf("Took %d naps. Feel so much better\n", numberOfNaps)
	os.Exit(0)
}
