package main

import (
	"fmt"
	functions "formative-11/web-server/functions"
	"net/http"
	"strconv"
)

func main() {
	jariJari := 7
	tinggi := 10

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str := "jariJari: " + strconv.Itoa(jariJari) + ", tinggi: " + strconv.Itoa(tinggi)

		fmt.Fprint(w, str)
		fmt.Fprint(w, ", volume: ", functions.Volume(jariJari, tinggi))
		fmt.Fprint(w, ", luas alas: ", functions.LuasAlas(jariJari))
		fmt.Fprint(w, ", keliling alas: ", functions.KelilingAlas(jariJari))
	})

	fmt.Println("starting web server at http://localhost:8080/")

	http.ListenAndServe(":8080", nil)
}
