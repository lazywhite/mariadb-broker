package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/kubernetes-sigs/service-catalog/contrib/pkg/broker/server"
	"github.com/kubernetes-sigs/service-catalog/pkg"
	"github.com/lazywhite/mariadb-broker/controller"
)

var options struct {
	Port int
}

func init() {
	flag.IntVar(&options.Port, "port", 8005, "use '--port' option to specify the port for broker to listen on")
	flag.Parse()
}

func main() {
	if flag.Arg(0) == "version" {
		fmt.Printf("%s/%s\n", path.Base(os.Args[0]), pkg.VERSION)
		return
	}

	server.Run(context.TODO(), ":" + strconv.Itoa(options.Port), controller.CreateController())
}
