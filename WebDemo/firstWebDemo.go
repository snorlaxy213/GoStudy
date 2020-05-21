package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Id int
	Name string
	sex int
}



func Hello(response http.ResponseWriter, request *http.Request) {

	user := &User{}
	user.Id = 1
	user.Name = "du"
	user.sex = 1
	data, _ := json.Marshal(user)
	fmt.Fprintf(response, string(data))
}



func GetJson(response http.ResponseWriter, request *http.Request){
	defer request.Body.Close()
	data,_ := ioutil.ReadAll(request.Body)
	user := &User{}
	_ = json.Unmarshal(data,user)

	user.Id = 2
	userJson, _ := json.Marshal(user)
	fmt.Fprintf(response, string(userJson))

}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/getJson",GetJson)
	http.ListenAndServe(":8080", nil)
}

