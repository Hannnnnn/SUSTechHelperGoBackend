package main

import "net/http"
import router "SUSTechHelperGoBackend/router"
import "log"

func main() {

	http.HandleFunc("/gpa", router.GPAQuery)
	http.HandleFunc("/course", router.CourseQuery)
	http.HandleFunc("/login", router.LoginQuery)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
