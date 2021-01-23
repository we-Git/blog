package main

import (
	"blog/github.com/gohub/blog/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
	//t,_:=time.LoadLocation("Asia/shanghai")
	//fmt.Println(time.Now().In(t))
}
