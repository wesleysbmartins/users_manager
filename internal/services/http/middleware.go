package http

import "net/http"

type headerObj struct {
	ContentType string
	Application string
}

var header = headerObj{
	ContentType: "Content-type",
	Application: "application/json",
}

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(header.ContentType, header.Application)
		next.ServeHTTP(w, r)
	})
}
