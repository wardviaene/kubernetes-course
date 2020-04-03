package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/save", save)
	http.HandleFunc("/read/", read)
	http.ListenAndServe(":8080", nil)
}

const filenamePrefix string = "/data/"

func save(rw http.ResponseWriter, req *http.Request) {
	data, _ := ioutil.ReadAll(req.Body) //data -> foo=bar
	fmt.Println("KV data", string(data))
	keyVal := strings.Split(string(data), "=")
	key := keyVal[0]   //foo
	value := keyVal[1] //bar

	filename := filenamePrefix + key //e.g. /data/foo
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	fmt.Println("Failed to open file " + filename + " due to " + err.Error())
	if err != nil {
		newFile, createErr := os.Create(filename)
		if createErr != nil {
			errMsg := "Failed to create file " + filename + " due to - " + createErr.Error()
			fmt.Println(errMsg)
			rw.Write([]byte(errMsg))
			return
		}
		file = newFile
		fmt.Println("Created new file", filename)
	}

	defer file.Close()

	if _, err = file.WriteString(value); err != nil {
		errMsg := "Failed to save data " + string(data) + " due to error - " + err.Error()
		fmt.Println(errMsg)
		rw.Write([]byte(errMsg))
		return
	}
	success := "Saved value " + value + " to " + filename
	fmt.Println(success)
	rw.Write([]byte(success))
}
func read(rw http.ResponseWriter, req *http.Request) {
	key := strings.TrimPrefix(req.URL.Path, "/read/") //http://kvstore:8080/read/foo (foo is the key)
	filename := filenamePrefix + key
	value, _ := ioutil.ReadFile(filename)
	if string(value) == "" {
		rw.Write([]byte("Key '" + key + "' does not exist in the store"))
		return
	}
	fmt.Println("Value for key " + key + " is " + string(value))
	rw.Write(value)
}
