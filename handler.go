package main

import (
    "encoding/json"
    "net/http"
    "io"
    "io/ioutil"
    "fmt"
)



func getData (w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader (http.StatusOK)
    if err := json.NewEncoder(w).Encode(users); err != nil {
        panic(err)
    }   
}


func postData (w http.ResponseWriter, r *http.Request){
    var u User
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    //fmt.Println (json.Marshal(body))
    if err!=nil{
        panic (err)
    }
    if  err:=r.Body.Close (); err!=nil{
        panic (err)
    }
    if err := json.Unmarshal(body, &u); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    u=createUser (u)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(u); err != nil {
        panic(err)


    }
}
