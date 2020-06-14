package main

import (
	"io/ioutil"
	"log"
	"net.matbm/munix/muinstaller/installer"
	"net.matbm/munix/muinstaller/parser"
	"net.matbm/munix/muinstaller/validator"
	"os"
)

const version = "1.0"

func main() {
	log.Printf("starting Munix installer version %s", version)

	if len(os.Args) < 2 {
		log.Fatal("usage: muinstaller [config-file.json]")
	}

	bytes, err := ioutil.ReadFile(os.Args[1])
	check(err)

	conf := parser.InstallConfig{}
	err = parser.ReadConfig(bytes, &conf)
	check(err)

	err = validator.ValidateConfig(conf)
	check(err)

	log.Print("config valid, starting Munix installation")
	err = installer.Install(conf)
	check(err)

	log.Print("install successful [ º_º]")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
