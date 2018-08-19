package main 


type User struct {
    Id          int     `json:"id"`
    Username    string  `json:"username"`
    Password    string  `json: "password"`
}

type Users []User //list of users