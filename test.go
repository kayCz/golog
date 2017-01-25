package main

import (
	log "./kaylog"
)

func main() {
	log.Info("hello")
	log.SetLevel(4)
	log.Fatal("fatal")
	log.SetFilePath("./log/go.log")
	log.Info("hello")
}
