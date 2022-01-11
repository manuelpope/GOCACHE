package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var cache = NewCache()

func ReturnAllKeysValues(w http.ResponseWriter, r *http.Request) {
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

func EchoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	metadata := map[string]string{}
	if err := decoder.Decode(&metadata); err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	cache.Set(metadata["key"], metadata["value"])
	fmt.Fprintf(w, "Payload: %v\n", metadata)
}

func RemoveRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	metadata := map[string][]string{}
	if err := decoder.Decode(&metadata); err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	go cache.Delete(metadata["keys"]...)
	fmt.Fprintf(w, "removedList: %v", metadata["keys"])

}

func main() {

	cache.Set("uno", "1")
	cache.Set("dos", "2")
	http.HandleFunc("/", EchoString)
	http.HandleFunc("/cache", ReturnAllKeysValues)
	http.HandleFunc("/addregister", PostRequest)
	http.HandleFunc("/deleteregister", RemoveRequest)

	log.Println("listing at port 10000")
	log.Println(http.ListenAndServe(":10000", nil))

}
