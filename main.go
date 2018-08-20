package main 

import (
    //"fmt"
    "log"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
 _  "github.com/jinzhu/gorm/dialects/mysql"
)

func dbConn() (db *gorm.DB){
    db, err := gorm.Open("mysql", "gouser:gopassword@/userdb")
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
    
    //db.Debug().DropTableIfExists(&User{}) 
    //Drops table if already exists
    //db.Debug().AutoMigrate(&User{}) 
    //db.CreateTable(&User{})
    //user:=User{Username: "user1", Password: "pass1"}
    //fmt.Println(db.NewRecord (user))
    //db.Create(&user)
    //fmt.Println(db.NewRecord(user) )
    //db.Find (&user)
    router:=NewRouter()
    log.Fatal (http.ListenAndServe(":8080", router))

}