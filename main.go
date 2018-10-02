package main 

import (
    //"fmt"
    //"os"
    "log"
    //"encoding/json"
    //"reflect"
    //"io/ioutil"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
 _  "github.com/jinzhu/gorm/dialects/mysql"
)

type  DBCred struct

{
    Username string `json:"username"`
    Password string `json:"password"`
    Name string     `json:"name"`
}

func dbConn() (db *gorm.DB){
    /*
    var cred DBCred
    configFile, err:=os.Open ("config.json")
    defer configFile.Close ()
    if err!=nil{
        log.Println (err)
    }
    log.Println ("Opened config.json")
    jsonParser:=json.NewDecoder (configFile)
    jsonParser.Decode (&cred)
    fmt.Println (cred)
    */

    db, err := gorm.Open("mysql", "root:<password>@/gotest")
    //defer db.Close()
    if err!=nil{
        panic(err)
        log.Println("Connection Failed to Open")
    }
    log.Println("Connection Established")

    db.LogMode(true)
    return db
}

func main (){
    router:=NewRouter()
    log.Fatal (http.ListenAndServe(":8080", router))

}
