package main

import (
  "fmt"
  "net/http"
  "time"
  "encoding/json"
)

func main() {
  http.HandleFunc("/hello", viewHandler)
  http.ListenAndServe(":8080", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {  
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-type", "application/json")  

  jsonMsg, err := getResponse() 
  if err != nil {
    http.Error(w, "Oops", http.StatusInternalServerError)
  }
  fmt.Fprintf(w, jsonMsg)
}

func getResponse() (string, error){
  unixtime := int32(time.Now().Unix())
  msg := Message{"Hi", "Hello Web!", unixtime}
  jbMsg, err := json.Marshal(msg)

  if err != nil {    
    return "", err
  }

  jsonMsg := string(jbMsg[:]) // converting byte array to string
  return jsonMsg, nil
}

type Message struct {
  Title string
  Body string
  Time int32
}