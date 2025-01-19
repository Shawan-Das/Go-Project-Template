package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/tools/system/service"
	"github.com/tools/system/util"
)

var _version = "default"

func main() {
	fmt.Println("Starting  iService ", _version)
	defer fmt.Println("Done....")
	// port := flag.Int("p", 7037, "Service listen port")
	bindAddress := flag.String("b", "0.0.0.0", "Bind address")
	verbose := flag.Bool("v", false, "Verbose output")
	flag.Parse()
	if *verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Config file missing")
		fmt.Println("account [flags] <path to config file> ")
		flag.Usage()
		os.Exit(1)
	}
	// -- added dynamic port option
	portint, err := strconv.Atoi(args[0])
	if err == nil {
		fmt.Println("port should be int", portint)
	}
	port := flag.Int("p", portint, "Service listen port")
	// Read the config file
	configBytes, err := ioutil.ReadFile(args[1])
	if err != nil {
		fmt.Println("Unable to read config file ", err)
		os.Exit(1)
	}
	util.ConfigDetails = args[1]
	if systemService := service.NewSystemRestService(configBytes, *verbose); systemService != nil {
		stopSignal := make(chan bool)
		termination := make(chan os.Signal)
		signal.Notify(termination, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-termination
			fmt.Println("SIGTERM/SIGINT received from os")
			stopSignal <- true
		}()
		systemService.Serve(*bindAddress, *port, stopSignal)
	} else {
		fmt.Println("unable to start the service ...")
		os.Exit(2)
	}
}
