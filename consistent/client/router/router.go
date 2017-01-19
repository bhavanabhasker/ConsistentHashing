package router

import (
	"assign4/client/consistent"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type InputData struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var servera = "one"
var serverb = "two"
var serverc = "three"
var serv string
var c = consistent.CreateConsistent()
var routes = Routes{
	Route{
		"GetAll",
		"GET",
		"/keys",
		GetAll,
	},
	Route{
		"Get",
		"GET",
		"/keys/{key_id}",
		Get,
	},
	Route{
		"Put",
		"PUT",
		"/keys/{key_id}/{value}",
		Put,
	},
}

func Initialize() {
	c = consistent.CreateConsistent()
	c.AddNode(servera)
	c.AddNode(serverb)
	c.AddNode(serverc)
}
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}

func Get(w http.ResponseWriter, r *http.Request) {
	var t InputData
	client := &http.Client{}
	vars := mux.Vars(r)
	key := vars["key_id"]
	url := "http://localhost:"
	url += fmt.Sprintf("%s%s%s", serv, "/keys/", key)
	request, err := http.NewRequest("GET", url, nil)
	result, _ := client.Do(request)
	resultdata, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}
	if e := json.Unmarshal(resultdata, &t); e != nil {
		log.Fatal(e)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
func GetAll(w http.ResponseWriter, r *http.Request) {
	var t []InputData
	client := &http.Client{}
	url := "http://localhost:"
	url += fmt.Sprintf("%s%s", serv, "/keys")
	request, err := http.NewRequest("GET", url, nil)
	result, _ := client.Do(request)
	resultdata, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}
	if e := json.Unmarshal(resultdata, &t); e != nil {
		log.Fatal(e)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
func Put(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	vars := mux.Vars(r)
	key := vars["key_id"]
	value := vars["value"]
	var userdata InputData
	userdata.Key, _ = strconv.Atoi(key)
	userdata.Value = value
	var userdat []InputData
	userdat = append(userdat, userdata)
	for _, u := range userdat {
		url := "http://localhost"
		url += ":"
		server, _ := c.AddData(u.Value)
		if server == servera {
			serv = "3000"
		} else if server == serverb {
			serv = "3001"
		} else {
			serv = "3002"
		}
		url += fmt.Sprintf("%s%s", serv, "/keys")
		param := url
		param += fmt.Sprintf("%s%s%s%s", "/", strconv.Itoa(u.Key), "/", u.Value)
		request, _ := http.NewRequest("PUT", param, nil)
		request.Header.Add("Content-Type", "application/json")
		_, health := client.Do(request)
		t := userdata
		if health != nil {
			c.RemoveNode(server)
			message := "Server not working.Try Put again"
			if err := json.NewEncoder(w).Encode(message); err != nil {
				panic(err)
			}
		} else {

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(t); err != nil {
				panic(err)
			}
		}
	}
}
