package main
import (
	"database/sql"
	"log"
	"time"
	
	"github.com/lib/pq"
	"fmt"
)

func main() {
	connectiondetails:="user=santhosha dbname=mydb_1 sslmode=disable"
	db,err:=sql.Open("postgres",connectiondetails)
	if(err!=nil) {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("succesfully connected to database !")
}
