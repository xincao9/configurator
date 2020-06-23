package main

import (
	"configurator/ms"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	master     string
	env        string
	group      string
	project    string
	version    string
	properties string
	key        string
)

func init() {
	master = *flag.String("master", "localhost:9090", "dkv master address")
	env = *flag.String("env", "", "environment")
	group = *flag.String("group", "", "group")
	project = *flag.String("project", "", "project")
	version = *flag.String("version", "", "version")
	properties = *flag.String("properties", "", "properties")
	if flag.Parsed() {
		flag.Parse()
	}
	key = fmt.Sprintf("configurator|%s|%s|%s|%s", env, group, project, version)
	if all([]string{master, env, group, project, version, properties}) == false {
		flag.Usage()
		os.Exit(0)
	}
}

func all(vs []string) bool {
	if len(vs) == 0 {
		return true
	}
	for _, v := range vs {
		if v == "" {
			return false
		}
	}
	return true
}

func main() {
	client, err := ms.New(master, time.Second)
	if err != nil {
		panic(err)
	}
	r, err := client.Put(key, properties)
	if err != nil {
		panic(err)
	}
	log.Printf("%v\n", r)
}
