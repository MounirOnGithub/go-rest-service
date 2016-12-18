package utils

import (
	"github.com/MounirOnGithub/go-rest-service/handler"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name: "Hello",
		Method: "GET",
		Pattern: "/",
		HandlerFunc: handler.Hello,
	},
	Route{
		Name: "GetGlucoses",
		Method: "GET",
		Pattern: "/glucoses",
		HandlerFunc: handler.GetGlucoses,
	},
	Route{
		Name: "GetGlucoseByID",
		Method: "GET",
		Pattern: "/glucose/{id}",
		HandlerFunc: handler.GetGlucoseByID,
	},
	Route{
		Name: "AddGlucose",
		Method: "POST",
		Pattern: "/glucoses",
		HandlerFunc: handler.AddGlucose,
	},
	Route{
		Name: "DeleteGlucose",
		Method: "DELETE",
		Pattern: "/glucose/{id}",
		HandlerFunc: handler.DeleteGlucose,
	},
}
