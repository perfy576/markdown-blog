package main

import (
	"fmt"
	"io"
	"net/http"
	"server/config"
	"server/service"
	"server/utils"
	"strconv"

	"github.com/BurntSushi/toml"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

type ServiceHandler func(service.Request, config.Config) service.Response

func ToHttpHandler(handler ServiceHandler, config config.Config, w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	request := service.Request{
		Body: string(body),
	}
	response := handler(request, config)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, response.Body)
}

func main() {

	var serverConfig config.ServerConfig

	// 读取Toml配置文件
	if _, err := toml.DecodeFile(utils.ConverPath("./config/server.toml"), &serverConfig); err != nil {
		fmt.Println("Error decoding Toml file:", err)
		return
	}

	config := config.Config{
		ServerConf: serverConfig,
	}

	http.HandleFunc("/blog/v1/get-markdown-content", func(w http.ResponseWriter, r *http.Request) {
		ToHttpHandler(service.GetMarkDownContent, config, w, r)
	})
	http.HandleFunc("/blog/v1/get-directory-structure", func(w http.ResponseWriter, r *http.Request) {
		ToHttpHandler(service.GetDirectoryStructure, config, w, r)
	})

	addr := config.ServerConf.Server.IP + ":" + strconv.Itoa(config.ServerConf.Server.Port)

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		fmt.Println("Error decoding Toml file:", err)
		return
	}
}
