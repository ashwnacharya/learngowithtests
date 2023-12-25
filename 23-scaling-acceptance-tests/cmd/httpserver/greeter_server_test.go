package main_test

import (
	"testing"
	"fmt"
	"time"
	"net/http"
	"github.com/ashwnacharya/go-specs-greet/specifications"
	"github.com/ashwnacharya/go-specs-greet/adapters"
	"github.com/ashwnacharya/go-specs-greet/adapters/httpserver"
)
func TestGreeterServer(t *testing.T) {
	var (
		port  = "8080"
		baseURL = fmt.Sprintf("http://localhost:%s", port)
		driver = httpserver.Driver{BaseURL: baseURL, Client: &http.Client {
			Timeout: 1 * time.Second,
		}}
	)

	adapters.StartDockerServer(t, port, "httpserver")
	specifications.GreetSpecification(t, &driver)
	specifications.CurseSpecification(t, &driver)
}