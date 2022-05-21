// Converts AMBWeather station messages into JSON and forwards it on to a configured server using HTTP POST
// Copyright 2022 Kyle Ceschi
// Distributed under the terms of the GNU Public License (GPLv3)

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Tuple struct {
	Name  string
	Value string
}
type AW2JConfig struct {
	PostHost   string
	ListenPort string
}

//the config, don't write to this struct because i didn't bother implementing any concurrency protection for it
//and it's shared by the routine that parses data.
var CONFIG AW2JConfig

func main() {
	//parse config file
	data, readFileError := os.ReadFile("config.json")
	if readFileError != nil {
		log.Fatal("Failed to read config.json" + readFileError.Error())
	}
	unMarshalError := json.Unmarshal(data, &CONFIG)
	if unMarshalError != nil {
		log.Fatal("Failed to unmarshal config.json" + unMarshalError.Error())
	}
	log.Println("Config loaded")
	log.Println(fmt.Sprintf("%#v", CONFIG))
	handler := func(w http.ResponseWriter, req *http.Request) {
		log.Println("Got New Request")
		log.Println(req)
		//the amb weather station sends GET requests with the data encoded in the URI so ignoring everything except GET
		if req.Method == "GET" {
			//fire off a routine to parse the data, the data has a timestamp in it so down stream
			//processors can deal with things if they take to long and arrive out of order
			go processData(req.RequestURI)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
		return
	}
	http.HandleFunc("/", handler)
	log.Println("Starting server on port " + CONFIG.ListenPort)
	log.Fatal(http.ListenAndServe(":80", nil))
}
func processData(uri string) {
	t, parseURIError := parseURI(uri)
	if parseURIError != nil {
		log.Println(parseURIError.Error())
		return
	}
	b, marshalError := json.Marshal(t)
	if marshalError != nil {
		log.Println(marshalError.Error())
		return
	}
	log.Println("Processed Data")
	log.Println(string(b))
	postError := postData(b)
	if postError != nil {
		log.Println(postError.Error())
	}
}

func parseURI(uri string) ([]Tuple, error) {
	//each data section is split by & in the uri
	strs := strings.Split(uri, "&")
	var arr []Tuple
	if len(strs) == 0 {
		//no data
		return arr, nil
	}
	//build an array of name value pairs from the uri data
	for s := range strs {
		//each item has a name and a value split by an equals sign
		d := strings.Split(strs[s], "=")
		if len(d) != 2 {
			//bad data
			log.Println("Got bad data skipping " + fmt.Sprintf("%#v", d))
			continue
		}
		arr = append(arr, Tuple{Name: d[0], Value: d[1]})
	}
	return arr, nil
}

func postData(data []byte) error {
	resp, postError := http.Post(CONFIG.PostHost, "application/json", bytes.NewBuffer(data))
	if postError != nil {
		return postError
	}
	if resp.StatusCode != 200 {
		return errors.New("Server Error: \n " + fmt.Sprintf("%#v", resp))
	}
	return nil
}
