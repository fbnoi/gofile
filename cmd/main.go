package main

import (
	"log"
	"os"

	"fbnoi.com/gofile"
)

func main() {
	gofile.Disk("local").WriteString("test.txt", "hello world\t", os.O_CREATE|os.O_RDWR|os.O_APPEND)
	txt, _ := gofile.Disk("local").ReadString("test.txt")
	log.Println(txt)
}
