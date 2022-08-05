package main

import (
	"fmt"
	"log"
	"net/http"
)

//端口
const webPort = "80"

type Config struct {
}

func main() {

	log.Printf("Starting broker service on port %s\n", webPort)

	app := Config{}

	// 定义http服务器
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// 启动服务器
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
