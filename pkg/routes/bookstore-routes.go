// package routes

// import (
// 	// "github.com/go-bookStore/pkg/controllers"
// 	"github.com/go-bookStore/pkg/controllers"
// 	"github.com/gorilla/mux"
// )

// var RegisterBookControllerRoutes = func(router *mux.Router) {

// 	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
// 	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
// 	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
// 	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
// 	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

// }
package routes

import (
	"go-bookStore/pkg/controllers"

	"github.com/gorilla/mux"
)

func RegisterBookControllerRoutes(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
