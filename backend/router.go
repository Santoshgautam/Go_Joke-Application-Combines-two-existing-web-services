package backend

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"html/template"
)

var controller = &Controller{Repository: Repository{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"GetNewJoke",
		"GET",
		"/GetNewJoke",
		controller.GetNewJoke,
	}}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	//serving curl as requirment
	router.HandleFunc("/", controller.GetNewJoke)
	//adding a ui
	router.HandleFunc("/ui", Home)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Home(w http.ResponseWriter, r *http.Request) {
	templates, err := template.ParseFiles("templates/home_page.html")
	if err != nil {
		fmt.Println("error")
	}
	//If errors show an internal server error message
	if err := templates.Execute(w, "Why Newton why?");
		err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
