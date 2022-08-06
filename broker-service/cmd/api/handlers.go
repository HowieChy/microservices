package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// RequestPayload描述此服务作为HTTP Post请求接受的JSON
type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

// AuthPayLoad是描述身份验证请求的嵌入式类型(在RequestPayLoad中）
type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Broker代理是一个测试处理程序，确保我们可以从web客户端访问代理
func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker!",
	}
	_ = app.writeJSON(w, http.StatusOK, payload)

}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "auth":
		app.authenticate(w, requestPayload.Auth)
	default:
		app.errorJSON(w, errors.New("unknown action"))
	}
}

//在broker中调用authentication微服务并发送回相应的响应
func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	//创建json，我们将发送给auth微服务
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	//调用Http服务
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// 确保我们拿回正确的状态代码
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	//创建我们将读取响应的变量.Body into
	var jsonFromService jsonResponse

	// 从auth服务解码json
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Authenticated!"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)
}
