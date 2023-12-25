package main

import (
	"net/http"
	"log"
	"context"
	"github.com/quii/go-graceful-shutdown"
)


type AcceptanceTests struct {
}

func (acceptancetests AcceptanceTests) SlowHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}


func main() {
	acceptancetests := new(AcceptanceTests)

	ctx := context.Background()

	httpServer := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(acceptancetests.SlowHandler)}

	server := gracefulshutdown.NewServer(httpServer)

	if err := server.ListenAndServe(ctx); err != nil {
		// this will typically happen if our responses aren't written before the ctx deadline, not much can be done
		log.Fatalf("uh oh, didnt shutdown gracefully, some responses may have been lost %v", err)
	}

	// hopefully, you'll always see this instead
	log.Println("shutdown gracefully! all responses were sent")
}