package main

import "goc/logface"

var log = logface.New(logface.TraceLevel)

func main() {
	convert2bin("close.png")
	select {}
}
