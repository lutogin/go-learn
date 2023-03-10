package gorilla_server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func handlerFromQuery(w http.ResponseWriter, r *http.Request) {
	payload := mux.Vars(r)
	name := r.URL.Query().Get("name")
	id := payload["id"]
	response := fmt.Sprintf("Product is %s and name=%s", id, name)
	fmt.Fprint(w, response)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age, _ := strconv.Atoi(r.FormValue("age"))
	response := fmt.Sprintf("Name is %s and age is %d", name, age)
	fmt.Fprint(w, response)
}

func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/products/{id:[0-9]+}?name=Vasia", handlerFromQuery)
	router.HandleFunc("/api", handlePost)
	http.Handle("/", router)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
