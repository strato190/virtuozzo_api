package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	//r.Handle("/", http.FileServer(http.Dir("./views/")))
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	//r.Handle("/", jwtMiddleware.Handler(VMHandler)).Methods("GET")
	r.Handle("/vms", jwtMiddleware.Handler(VMHandler)).Methods("GET")
	r.Handle("/get-token", GetTokenHandler).Methods("GET")
	r.Handle("/vms", jwtMiddleware.Handler(VMAddHandler)).Methods("POST")
	//r.Handle("/cts", jwtMiddleware.Handler(CTHandler)).Methods("GET")
	//r.Handle("/cts", jwtMiddleware.Handler(CTAddHandler)).Methods("POST")

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

	//jsn, _ := ioutil.ReadFile("config.json")
	//json.Unmarshal(jsn, &intf)

}
