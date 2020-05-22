package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//User 首字母大写，代表对外部可见，首字母小写代表对外部不可见，
//适用于所有对象，包括函数、方法
type User struct {
	ID int
	Name string
	sex int
}

func TestStaticStruct(response http.ResponseWriter, request *http.Request) {

	//user := new(User)
	user := &User{}
	user.ID = 1
	user.Name = "du"
	user.sex = 1
	data, _ := json.Marshal(user)
	fmt.Fprintf(response, string(data))
}

func TestJsonStruct(response http.ResponseWriter, request *http.Request){
	defer request.Body.Close()
	data,_ := ioutil.ReadAll(request.Body)
	user := &User{}
	_ = json.Unmarshal(data,user)
	userJson, _ := json.Marshal(user)

	fmt.Fprintf(response, string(userJson))
}

func TestUrlParam(response http.ResponseWriter,request *http.Request){
	request.ParseForm()

	fmt.Println(request.Form)
	fmt.Println("Path: ", request.URL.Path)
	fmt.Println("scheme:", request.URL.Scheme)
	fmt.Println(request.Form["url_long"])
	for k,v := range request.Form{
		fmt.Println("key: ", k)
        fmt.Println("val: ", strings.Join(v, " "))
	}
	fmt.Fprintf(response, "TestUrlParam success")
}

func main() {
	http.HandleFunc("/testStaticStruct", TestStaticStruct)
	http.HandleFunc("/testWebStruct",TestJsonStruct)
	http.HandleFunc("/testTestUrlParam",TestUrlParam)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

