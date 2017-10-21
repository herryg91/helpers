package main

import "github.com/herryg91/helpers/hUrls"

func main() {
	benchmarkFunction()
}

func benchmarkFunction() {
	hUrls.IsURL("http://golang.org")
}
