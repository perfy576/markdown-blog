package service

import "net/http"

type Request struct {
	Body string
}

type Response struct {
	Body string
}

type ServiceHandler func(request Request) Response

type HttpHandler func(w http.ResponseWriter, r *http.Request)
