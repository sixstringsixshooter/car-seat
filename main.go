package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/sixstringsixshooter/car-seat/config"
)

const usage = `version: %s - git: %s
Usage: %s [-c config file]
Options:
	-h            help menu
	-v            show version
	-c path       configuration file path
`

func main() {
	flag.Usage = func() {
		w := os.Stderr
		for _, arg := range os.Args {
			if arg == "-h" {
				w = os.Stdout
				break
			}
		}
		fmt.Fprintln(w, usage)
	}

	b, err := ioutil.ReadFile("./config/testdata/test-config.toml")
	if err != nil {
		log.Fatalln(err)
	}

	var conf config.Config
	if _, err := toml.Decode(string(b), &conf); err != nil {
		log.Fatalln(err)
	}

	c, err := newController(&conf)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("starting API...")
	if err := http.ListenAndServe(":8080", c.router); err != nil {
		log.Fatal(err.Error())
	}
}
