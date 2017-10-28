package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var intf Instances
	var err error

	r := mux.NewRouter()

	//r.Handle("/", http.FileServer(http.Dir("./views/")))
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.Handle("/vms", jwtMiddleware.Handler(VMHandler)).Methods("GET")
	r.Handle("/get-token", GetTokenHandler).Methods("GET")
	//	r.Handle("/vms", jwtMiddleware.Handler(VmAddHandler)).Methods("POST")

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

	jsn, _ := ioutil.ReadFile("config.json")
	json.Unmarshal(jsn, &intf)
	if len(intf.Cts) > 0 {
		for _, c := range intf.Cts {
			err = createCT(c)
			err = configCT(c)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	if len(intf.Vms) > 0 {
		for _, v := range intf.Vms {
			//err = create_vm(v)
			err = configVM(v)
			err = configVMNetwork(v)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
