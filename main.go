package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/MounirOnGithub/go-rest-service/handler"
	"github.com/MounirOnGithub/go-rest-service/utils"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/meatballhat/negroni-logrus"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
)

var (
	// command line parameters
	port      = 8080
	logLevel  = "debug"
	logFormat = "text_color"

	// Version is the version of the software
	Version string
	// BuildStmp is the build date
	BuildStmp string
	// GitHash is the git build hash
	GitHash string
)

func main() {

	cliApp := cli.NewApp()

	timeStmp, err := strconv.Atoi(BuildStmp)
	if err != nil {
		timeStmp = 0
	}

	cliApp.Version = Version + ", build on " + time.Unix(int64(timeStmp), 0).String() + ", git hash " + GitHash
	cliApp.Name = "go rest service"
	cliApp.Authors = []cli.Author{{Name: "mkh"}}
	cliApp.Copyright = "Mounir Khanouri" + strconv.Itoa(time.Now().Year())
	cliApp.Usage = "Example of go REST service with JWT"

	cliApp.Flags = []cli.Flag{
		cli.IntFlag{
			Value: port,
			Name:  "port",
			Usage: "Set the listening port of the web server",
		},
		cli.StringFlag{
			Value: logLevel,
			Name:  "logl",
			Usage: "Set the output log level (debug, info, warning, error)",
		},
		cli.StringFlag{
			Value: logFormat,
			Name:  "logf",
			Usage: "Set the log formatter (logstash or text)",
		},
	}

	cliApp.Action = func(c *cli.Context) error {
		port = c.Int("port")
		logLevel = c.String("logl")
		logFormat = c.String("logf")

		fmt.Print("* --------------------------------------------------- *\n")
		fmt.Printf("|   port                    : %d\n", port)
		fmt.Printf("|   logger level            : %s\n", logLevel)
		fmt.Printf("|   logger format           : %s\n", logFormat)
		fmt.Print("* --------------------------------------------------- *\n")

		n := negroni.New()
		n.Use(negronilogrus.NewMiddlewareFromLogger(logrus.StandardLogger(), "go-rest-service"))

		// Recovery middleware for responding 500 while having a panic
		recovery := negroni.NewRecovery()
		recovery.PrintStack = false
		n.Use(recovery)

		// Router
		r := mux.NewRouter()
		r.HandleFunc("/login", handler.LogIn).Methods(http.MethodPost)

		// User sub router
		userSubRouter := mux.NewRouter().PathPrefix("/user").Subrouter().StrictSlash(true)
		userSubRouter.HandleFunc("/", handler.Hello).Methods(http.MethodGet)

		// Using middleware for the user sub router
		r.PathPrefix("/user").Handler(negroni.New(
			negroni.HandlerFunc(utils.JWTValidationMiddleware),
			negroni.Wrap(userSubRouter),
		))

		n.UseHandler(r)
		n.Run(":8080")
		return nil
	}

	cliApp.Run(os.Args)

}
