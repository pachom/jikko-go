package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"text/template"
	"github.com/gorilla/mux"
	// Driver para SQLite3
	//_ "github.com/mattn/go-sqlite3"
)

func sortNumbers(w http.ResponseWriter, r *http.Request) {
	unsorted := []int{3, 5, 5, 6, 8, 3, 4, 4, 7, 7, 1, 1, 2}
	presorted := []int{3, 5, 5, 6, 8, 3, 4, 4, 7, 7, 1, 1, 2}

	sorted := numbersSort(presorted)
	numberLists := make(map[string][]int)
	numberLists["unsorted"] = unsorted

	numberLists["sorted"] = sorted
	json.NewEncoder(w).Encode(numberLists)
}

func numbersSort(numList []int) []int {
	j := 0
	for k := 1; k < 10; k++ {
		for i := 0; i < len(numList); i++ {
			if numList[i] == k && j < k {
				if j != i {
					swap(j, i, &numList)
				}
				j++
				continue
			}
		}
	}
	return numList
}

func swap(origin, new int, pointList *[]int) {
	listTemp := *pointList
	temp := listTemp[origin]
	listTemp[origin] = listTemp[new]
	listTemp[new] = temp
}

var db *sql.DB

func getConn() *sql.DB {
	if db != nil {
		return db
	}

	var err error

	db, err := sql.Open("sqlite3", "data.sqlite")
	if err != nil {
		panic(err.Error())
	}
	log.Println("Data Base connected")
	return db

}

func (n *User) getAll() ([]User, error) {
	db := getConn()
	q := `SELECT
            id, account, password, created_at, updated_at
            FROM users`
	// Execute the query
	rows, err := db.Query(q)
	if err != nil {
		return []User{}, err
	}
	// Will close the resource at the end
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		rows.Scan(
			&n.ID,
			&n.Account,
			&n.Password,
			&n.CreatedAt,
			&n.UpdatedAt,
		)

		users = append(users, *n)
	}
	return users, nil
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// User struct pointer.
	u := new(User)

	users, err := u.getAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// convert the users slice to JSON format.
	j, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Response.
	w.WriteHeader(http.StatusOK)

	// content type
	w.Header().Set("Content-Type", "application/json")

	// Response in JSON.
	w.Write(j)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func main() {

	//dbConn()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/sort", sortNumbers)
	router.HandleFunc("/users", getUsersHandler)

	log.Fatal(http.ListenAndServe(":3000", router))
}
