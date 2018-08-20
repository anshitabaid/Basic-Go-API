package main 


type User struct {
    Id          int     `gorm:"primary_key"; "AUTO_INCREMENT"`
    Username    string  `gorm: "username"; "size:255"`
    Password    string  `gorm: "password"; "size:255"`
}

type Users []User //list of users