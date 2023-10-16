// получить данные сервиса 1: curl http://localhost:8080/data
// получить данные из сервиса 1 через сервис : curl http://localhost:8081/requestA
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "PRIVET OT SERVISA 1") //привет в другом сервисе =)
	})

	http.HandleFunc("/request2", func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get("http://localhost:8081/data")
		if err != nil {
			http.Error(w, "OSHIBKA", http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		fmt.Fprint(w, "Data from Service 2: ", res.Status)
	})

	http.ListenAndServe(":8080", nil)
}
