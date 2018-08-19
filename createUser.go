package main 

var currentId int
var users Users


func createUser (u User) User{
        currentId+=1
        u.Id=currentId
        users=append (users, u)
        return u
}
