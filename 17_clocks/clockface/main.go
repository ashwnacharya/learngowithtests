package main

import (
	"os"
	"time"
	"github.com/ashwnacharya/learngowithtests/17_clocks"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}