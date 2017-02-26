package routes

import (
	"net/http"

	"stocks/handlers"
)

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
		"Stock",
		"GET",
		"/",
		handlers.GetStocks,
	},
	/*Route{
		"Stock",
		"POST",
		"/",
		handlers.AddStock,
	},*/
	Route{
		"Stock",
		"POST",
		"/stock",
		handlers.UpdateStock,
	},
}
