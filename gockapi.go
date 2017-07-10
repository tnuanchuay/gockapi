package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
)

const API_FOLDER = "api"
const PORT = 8080

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/{file:[a-z/]+}", getFileAndResponse)

	http.Handle("/", router)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}

func getFileAndResponse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	file := params["file"]
	file = strings.Replace(file, "/", "-", -1)
	path := fmt.Sprintf("./%s/%s.%s", API_FOLDER, file, r.Method)
	fmt.Println(path)
	bfile, err := ioutil.ReadFile(path)
	if err != nil {
		w.WriteHeader(404)
	}else{
		w.Write([]byte(bfile))
	}


}