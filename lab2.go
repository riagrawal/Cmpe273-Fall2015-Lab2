package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
    "io/ioutil"
  )
  
type req_struct struct{
       Name string
}

func post(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
   var t req_struct  
   err = json.Unmarshal(body, &t)
    if (err != nil || t.Name == "") {
        http.Error(rw, "Bad Request, check request payload", http.StatusBadRequest)
        return
    }
  payload, err := json.Marshal("{\"greeting\":\"Hello, " + t.Name + "!\"}")
  if err != nil {
     http.Error(rw,"Bad Request, check request payload" , http.StatusInternalServerError)
     return
  }
  rw.Header().Set("Content-Type", "application/json")
  rw.Write(payload)
}

func main() {
    mux := httprouter.New()
    mux.POST("/hello",post)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}