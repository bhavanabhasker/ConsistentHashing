package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var Routess = Routes{
	Route{
		"PutData",
		"PUT",
		"/keys/{keys_id}/{value}",
		PutData,
	},
	Route{
		"Get",
		"GET",
		"/keys/{keys_id}",
		Get,
	},
	Route{
		"GetAll",
		"Get",
		"/keys",
		GetAll,
	},
}

type DataStore struct {
	KeyId int    `json:"key"`
	Value string `json:"value"`
}

type Data []DataStore

var Datastore = make(Data, 0)

func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyId := vars["keys_id"]
	index, _ := strconv.Atoi(keyId)
	t := Findkey(index)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
func GetAll(w http.ResponseWriter, r *http.Request) {

	t := FindAll()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
func PutData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyId := vars["keys_id"]
	key, _ := strconv.Atoi(keyId)
	fmt.Println(key)
	value := vars["value"]
	Update(key, value)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
func Findkey(key int) DataStore {
	var t DataStore
	for i := 0; i < len(Datastore); i++ {
		if key == Datastore[i].KeyId {
			t.KeyId = key
			t.Value = Datastore[i].Value
		}
	}
	return t
}
func FindAll() Data {
	return Datastore
}
func Update(key int, value string) {
	var cache DataStore

	cache.KeyId = key
	cache.Value = value
	Datastore = append(Datastore, cache)
}
