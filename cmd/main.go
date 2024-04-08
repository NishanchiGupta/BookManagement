// package main

// import (
// 	"log"
// 	"net/http"

// 	"go-bookStore/pkg/routes"

// 	"github.com/gorilla/mux"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// )

//	func main() {
//		r := mux.NewRouter()
//		routes.RegisterBookControllerRoutes(r)
//		http.Handle("/", r)
//		log.Fatal(http.ListenAndServe("localhost:9010", r))
//	}
package main

import (
	"log"
	"net/http"

	"go-bookStore/pkg/routes"

	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookControllerRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", nil))
}
