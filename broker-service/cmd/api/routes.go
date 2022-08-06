package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// 用于在服务器端执行预检 CORS 检查
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	//Heartbeat监控端点以检查服务器
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/", app.Broker)

	//broker微服务监听前端请求，之后像认证微服务发出请求，认证微服务接收请求并返回数据
	mux.Post("/handle", app.HandleSubmission)

	return mux
}
