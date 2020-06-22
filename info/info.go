package info

import (
	"configurator/ms"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	FN         = "application"
	Ext        = "yaml"
	env        = "env"
	group      = "group"
	project    = "project"
	version    = "version"
	master     = "master"
	slaves     = "slaves"
	path       = "HOME"
	syncPeriod = 30
)

var (
	I *info
)

func init() {
	if I != nil {
		return
	}
	var err error
	I, err = newInfo()
	if err != nil {
		log.Fatalf("Fatal error configurator newInfo : %v\n", err)
	}
	err = I.sync()
	if err != nil {
		log.Fatalf("Fatal error configurator sync : %v\n", err)
	}
	go func() {
		for range time.Tick(time.Second * syncPeriod) {
			err = I.sync()
			if err != nil {
				log.Printf("sync err: %v", err)
			}
		}
	}()
}

func newInfo() (*info, error) {
	inf := &info{}
	inf.env = os.Getenv(env)
	inf.group = os.Getenv(group)
	inf.project = os.Getenv(project)
	inf.version = os.Getenv(version)
	inf.master = os.Getenv(master)
	inf.Path = os.Getenv(path)
	var err error
	if os.Getenv(slaves) != "" {
		inf.slaves = strings.Split(os.Getenv(slaves), ",")
		inf.client, err = ms.NewMS(inf.master, inf.slaves, time.Second)
	} else {
		inf.client, err = ms.New(inf.master, time.Second)
	}
	return inf, err
}

type info struct {
	env     string
	group   string
	project string
	version string
	master  string
	slaves  []string
	Path    string
	client  *ms.Client
}

func (inf *info) key() string {
	return fmt.Sprintf("configurator|%s|%s|%s|%s", inf.env, inf.group, inf.project, inf.version)
}

func (inf *info) sync() error {
	k := inf.key()
	r, err := inf.client.Get(k)
	if err != nil {
		return err
	}
	if r.Code != 200 {
		return fmt.Errorf("dkv response code: %d, message: %s\n", r.Code, r.Message)
	}
	log.Printf("dkv response: %v\n", r)
	d := r.KV.V
	if d == "" {
		return nil
	}
	err = ioutil.WriteFile(filepath.Join(inf.Path, fmt.Sprintf("%s.%s", FN, Ext)), []byte(d), 0644)
	if err != nil {
		return err
	}
	return nil
}
