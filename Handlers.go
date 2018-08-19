package main 

import(
	//"encoding/json"
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)




func display(w http.ResponseWriter, r *http.Request){
	db,err := sql.Open("mysql", "root:<password>@/gotest")
	if err!=nil {
		panic(err)
	}
	fmt.Println(err)
	rows,err := db.Query("SELECT * FROM quote")
	if err!=nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var sntc string
		var auth string
		err = rows.Scan(&id,&sntc,&auth)
		if err!=nil {
		panic(err)
		}
		fmt.Println(id)
		fmt.Println(sntc)
		fmt.Println(auth)
		
	}
}

func add(w http.ResponseWriter, r *http.Request){
	fmt.Println("here")
}
