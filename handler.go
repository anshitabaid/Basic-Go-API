package main

import (
    "encoding/json"
    "net/http"
    "io"
    "io/ioutil"
    //"log"
    //"fmt"
    //"reflect"
    _ "github.com/go-sql-driver/mysql"
    //"github.com/jinzhu/gorm"
        //_  //"github.com/jinzhu/gorm/dialects/mysql"
)



func getData (w http.ResponseWriter, r *http.Request){
    db:=dbConn ()
    users:=db.Find(&users).Value
    //fmt.Println(reflect.TypeOf (users))
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader (http.StatusOK)
    if err := json.NewEncoder(w).Encode(users); err != nil {
        panic(err)
    }   
}


func postData (w http.ResponseWriter, r *http.Request){
    db:=dbConn()
    //user:=User{Username: "user1", Password: "pass1"}
    //db.Create (&user)
    var u User
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    //fmt.Println (json.Marshal(body))
    if err!=nil{
        panic (err)
    }
    if  err:=r.Body.Close (); err!=nil{
        panic (err)
    }
    //fmt.Println (json.Encode (body))
    
    if err := json.Unmarshal(body, &u); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }
    
    //if err:=json.NewDecoder (r.Body).Decode (&u); err!=nil{
      //  panic (err)
    //}
    db.Create (&u)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(u); err != nil {
        panic(err)
    /*
    user:=createUser (u)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(user); err != nil {
        panic(err)
    //u1,_:=json.Marshal(u)
    db.Create (&body)
    */


    }
}
