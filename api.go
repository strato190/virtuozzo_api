package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

//VMHandler get vms list
var VMHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	hostedVms, _ := getVMS()
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(hostedVms))
})

//CTHandler get vms list
var CTHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	hostedCts, _ := getCTS()
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(hostedCts))
})

//VMAddHandler add's vm to the host
var VMAddHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var intf Instances

	jsn, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(jsn, &intf)
	for _, v := range intf.Vms {

		if err := createVM(v); err != nil {
			log.Fatal(err)
			payload, _ := json.Marshal(err)
			w.Write([]byte(payload))
		}
		if err := configVM(v); err != nil {
			log.Fatal(err)
			payload, _ := json.Marshal(err)
			w.Write([]byte(payload))
		}
		if err := configVMNetwork(v); err != nil {
			log.Fatal(err)
			payload, _ := json.Marshal(err)
			w.Write([]byte(payload))
		}

	}
	w.Header().Set("Content-Type", "application/json")

})

//get token

var mySigningKey = []byte("SuperSecretKey")

//GetTokenHandler gets token for user
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if checkAuth(w, r) {
		/* Create the token */
		token := jwt.New(jwt.SigningMethodHS256)

		/* Create a map to store our claims */
		claims := token.Claims.(jwt.MapClaims)

		/* Set token claims */
		claims["admin"] = true
		claims["name"] = "Ado Kukic"
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		/* Sign the token with our secret */
		tokenString, _ := token.SignedString(mySigningKey)

		/* Finally, write the token to the browser window */
		w.Write([]byte(tokenString))
	} else {
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
	}
})

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
