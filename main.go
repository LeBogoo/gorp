package main

import (
	"fmt"
	"gorp/utils"
	"io"
	"log"
	"net/http"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		configs := utils.ParseConfig()

		for _, config := range configs {
			if config.Host == r.Host {
				// Create a new HTTP request to the config.Server URL
				req, err := http.NewRequest(r.Method, config.Server+r.URL.Path, r.Body)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Set the headers from the original request
				for name, value := range r.Header {
					req.Header.Set(name, value[0])
				}

				// Send the request and get the response
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadGateway)
					return
				}
				defer resp.Body.Close()

				// Copy the headers and body from the response to the client's response writer
				for name, values := range resp.Header {
					for _, value := range values {
						w.Header().Add(name, value)
					}
				}
				w.WriteHeader(resp.StatusCode)
				io.Copy(w, resp.Body)

				return
			}
		}

		w.Write([]byte(fmt.Sprintf("No config for '%s' found.", r.Host)))
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Print("Executing finalHandler")
	w.Write([]byte("OK"))
}

func main() {
	const PORT = 18257
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", middlewareOne(finalHandler))

	log.Printf("Listening on :%d...", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux)
	log.Fatal(err)
}
