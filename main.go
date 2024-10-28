package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/tools/iservice/service"
	"github.com/tools/iservice/util"
)

var _version = "default"

// @title SICL IService API
// @version 2.0
// @description An API to perform `SICL IService API` operations. You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).
// @termsOfService http://swagger.io/terms/

// @contact.name Sumon Sarker
// @contact.email suman@satcombd.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @externalDocs.description OpenAPI
// @externalDocs.url https://swagger.io/resources/open-api/
// TODO: During local development, ensure to set the Swagger host to 'localhost:7026'. However, for online deployment, use the host '118.67.213.45:7026' to seamlessly interact with Swagger.
// @host 118.67.213.45:7037
// @BasePath /
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	fmt.Println("Starting  iService ", _version)
	defer fmt.Println("Done....")
	port := flag.Int("p", 7037, "Service listen port")
	bindAddress := flag.String("b", "0.0.0.0", "Bind address")
	verbose := flag.Bool("v", false, "Verbose output")
	flag.Parse()
	if *verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Config file missing")
		fmt.Println("account [flags] <path to config file> ")
		flag.Usage()
		os.Exit(1)
	}
	// Read the config file
	configBytes, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println("Unable to read config file ", err)
		os.Exit(1)
	}
	util.ConfigDetails = args[0]
	if iService := service.NewRestService(configBytes, *verbose); iService != nil {
		stopSignal := make(chan bool)
		termination := make(chan os.Signal)
		signal.Notify(termination, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-termination
			fmt.Println("SIGTERM/SIGINT received from os")
			stopSignal <- true
		}()
		iService.Serve(*bindAddress, *port, stopSignal)
	} else {
		fmt.Println("unable to start the service ...")
		os.Exit(2)
	}
}
