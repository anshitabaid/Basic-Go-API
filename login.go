package main

import (
    "net/http"
    "encoding/json"
    "io/ioutil"
    "io"
    "fmt"
    //"time"
)


func internalPageHandler (w http.ResponseWriter, r *http.Request){
    username:=getUserName(r)
    if username==""{
        http.Redirect (w,r, "/", 302)
    }
    var resp Resp
    resp.Success="true"
    resp.Message="You are logged in"
    resp.Data.Username=username
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader (http.StatusOK)
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        panic(err)
    }
}

func indexPageHandler (w http.ResponseWriter, r *http.Request){
    var resp Resp
    username:=getUserName (r)
    if username==""{
        fmt.Printf("hihihihih")
    resp.Success="false"
    resp.Message="You are not logged in"
    resp.Data.Username=""
    } else {
        resp.Success="true"
        resp.Message="You are logged in"
        resp.Data.Username=""
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader (http.StatusOK)
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        panic(err)
    }
}


func loginHandler (w http.ResponseWriter, r *http.Request){
    var u User
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err!=nil{
        panic (err)
    }
    if  err:=r.Body.Close (); err!=nil{
        panic (err)
    }
    if err := json.Unmarshal(body, &u); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        
        }
    name:=u.Username 
    pass:=u.Password
    redirectTarget:="/"
    if name!="" && pass!="" { 
        if pass=="password" && name=="username"{
            setSession (name, w)
            redirectTarget ="/internal"
        }
    }

    http.Redirect (w, r, redirectTarget, 302)
    /**
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(u); err != nil {
        panic(err)
        */
}

func logoutHandler (response http.ResponseWriter, request *http.Request){
    clearSession (response)
    //fmt.Printf (getUserName (request))
    //time.Sleep (1*time.Second)
    http.Redirect (response, request, "/", 302)
}
