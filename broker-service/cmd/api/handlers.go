package main

import (
	"net/http"
)

// Broker代理是一个测试处理程序，确保我们可以从web客户端访问代理
func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker!",
	}
	_ = app.writeJSON(w, http.StatusOK, payload)

}
