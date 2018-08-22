package main 

import(
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"
	"io"
	"io/ioutil"
	_ "github.com/go-sql-driver/mysql"
)

func display(w http.ResponseWriter, r *http.Request){
	db,err := sql.Open("mysql", "root:lolseeyou123@/gotest")
	if err!=nil {
		panic(err)
	}
	
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

		//fmt.Println(id)
		//fmt.Println(sntc)
		//fmt.Println(auth)
	}
	
}

func add(w http.ResponseWriter, r *http.Request){
	var q Quote
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err!=nil {
		panic(err)
	}

	if  err:=r.Body.Close (); err!=nil{
        panic (err)
	}

	if err := json.Unmarshal(body, &q); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
	}

	db,err := sql.Open("mysql", "root:lolseeyou123@/gotest")
	if err!=nil {
		panic(err)
	}

	stmt,err:=db.Prepare("INSERT INTO quote(sentence,author) VALUES (?,?)")
	if err!=nil {
		panic(err)
	}

	fmt.Print(q)

	res,err := stmt.Exec(q.sentence,q.author)
	if err!=nil {
		panic(err)
	}

	id,err := res.LastInsertId()
	if err!=nil {
		panic(err)
	}

	fmt.Println(id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(q); err != nil {
        panic(err)
    }

}
