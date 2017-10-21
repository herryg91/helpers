package main

import (
	"fmt"
	"time"

	"github.com/herryg91/helpers/hUrls"
)

func main() {
	benchmarkFunction()
}

func benchmarkFunction() {
	timeStart := time.Now()
	hUrls.IsURL("http://golang.org")
	fmt.Println("[package:hUrls][func:IsURL] ", time.Since(timeStart).Nanoseconds(), "ns")
}
