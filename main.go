package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var cache = NewCache()

func returnAllKeysValues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mapAux := map[string]string{}
	key := r.URL.Query().Get("key")

	if key == "" {
		log.Println("Url Param 'key' is missing")
		mapAux, _ = cache.All()
		jsonString, _ := json.Marshal(mapAux)
		w.Write(jsonString)
		return
	}

	mapAux[key], _ = cache.Get(key)

	jsonString, _ := json.Marshal(mapAux)
	w.Write(jsonString)

}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func main() {

	cache.Set("uno", "1")
	cache.Set("dos", "2")
	http.HandleFunc("/", echoString)
	http.HandleFunc("/cache", returnAllKeysValues)

	log.Println(http.ListenAndServe(":10000", nil))

}
