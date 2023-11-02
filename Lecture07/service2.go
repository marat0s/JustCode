// получить данные из сервиса 2: curl http://localhost:8081/data
// получить данные  из сервиса 2 через сервис 1: curl http://localhost:8080/requestB

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ZDRAVSTUY OT SERVISA 2")
	})

	http.HandleFunc("/request1", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:8080/data")
		if err != nil {
			http.Error(w, "OSHIBKA", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		fmt.Fprint(w, "Data from Service 1: ", resp.Status)
	})

	http.ListenAndServe(":8081", nil)
}
