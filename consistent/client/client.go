package main

import (
	"assign4/client/consistent"
	"assign4/client/router"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type InputData struct {
	Key   int
	Value string
}
type Data []InputData

func main() {
	var option int32
	//check the options
	commandargs := os.Args[1:]
	if len(commandargs) == 0 {
		fmt.Println("Select an option")
		fmt.Println("1. Execute code as a REST Application")
		fmt.Println("2. Execute code without supplying the data")
		fmt.Scan(&option)
	}
	if option == 2 {
		client := &http.Client{}
		c := consistent.CreateConsistent()
		servera := "one"
		serverb := "two"
		serverc := "three"
		c.AddNode(servera)
		c.AddNode(serverb)
		c.AddNode(serverc)
		userdata := Data{
			{Key: 1, Value: "a"},
			{Key: 2, Value: "b"},
			{Key: 3, Value: "c"},
			{Key: 4, Value: "d"},
			{Key: 5, Value: "e"},
			{Key: 6, Value: "f"},
			{Key: 7, Value: "g"},
			{Key: 8, Value: "h"},
			{Key: 9, Value: "i"},
			{Key: 10, Value: "j"},
		}

		i := 0
		for _, u := range userdata {
			i = i + 1
			server, err := c.AddData(u.Value)
			if err != nil {
				log.Fatal(err)
			}
			url := "http://localhost"
			url += ":"
			var serv string
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
			request, err := http.NewRequest("PUT", param, nil)
			request.Header.Add("Content-Type", "application/json")
			_, health := client.Do(request)
			if health != nil {
				c.RemoveNode(serv)
				userdata = append(userdata, u)
			}
			// implement get for specific key id
			fmt.Println("The below data is inserted in the server", serv)
			PrinteachContents(serv, u)

		}

		//implement get operation
		fmt.Println("*******************************************************")
		fmt.Println("Printing the contents from server with the port no: 3000")
		PrintContents("3000")
		fmt.Println("***********************************************************")
		fmt.Println("Printing the contents from server with the port no : 3001")
		PrintContents("3001")
		fmt.Println("************************************************************")
		fmt.Println("Printing the contents from server with the port no : 3002")
		PrintContents("3002")
	} else if option == 1 {
		router.Initialize()
		router := router.NewRouter()
		log.Fatal(http.ListenAndServe(":8080", router))
	}
}
func PrintContents(port string) {
	client := &http.Client{}
	url := "http://localhost:"
	url += fmt.Sprintf("%s%s", port, "/keys")
	request, err := http.NewRequest("GET", url, nil)
	result, _ := client.Do(request)
	resultdata, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(resultdata))
}
func PrinteachContents(port string, u InputData) {
	client := &http.Client{}
	url := "http://localhost:"
	url += fmt.Sprintf("%s%s%s", port, "/keys/", strconv.Itoa(u.Key))
	request, err := http.NewRequest("GET", url, nil)
	result, _ := client.Do(request)
	resultdata, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(resultdata))
}
