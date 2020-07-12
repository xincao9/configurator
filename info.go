package configurator

import (
    "fmt"
    "github.com/subosito/gotenv"
    "github.com/xincao9/dkv/client"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
    "sync"
    "time"
)

const (
	FN         = "application"
	Ext        = "json"
	envKey     = "env"
	groupKey   = "group"
	projectKey = "project"
	versionKey = "version"
	masterKey  = "master"
	slavesKey  = "slaves"
	pathKey    = "path"
	path       = "HOME"
	syncPeriod = 30
)

var (
	I *info
    once = &sync.Once{}
)

func init () {
    startInfo()
}

func startInfo () {
    once.Do(initInfo)
}

func initInfo() {
    var err error
    err = gotenv.Load(filepath.Join(os.Getenv(path), ".env"))
    if err != nil {
        log.Fatalln(err)
        log.Printf("warning file '%s' not found", filepath.Join(os.Getenv(path), ".env"))
    }
	if I != nil {
		return
	}
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
	inf.env = os.Getenv(envKey)
	inf.group = os.Getenv(groupKey)
	inf.project = os.Getenv(projectKey)
	inf.version = os.Getenv(versionKey)
	inf.master = os.Getenv(masterKey)
	inf.Path = filepath.Join(os.Getenv(path), inf.group, inf.project, inf.version)
	err := os.MkdirAll(inf.Path, 0755)
	if err != nil {
		return inf, err
	}
	if os.Getenv(slavesKey) != "" {
		inf.slaves = strings.Split(os.Getenv(slavesKey), ",")
		inf.client, err = client.NewMS(inf.master, inf.slaves, time.Second)
	} else {
		inf.client, err = client.New(inf.master, time.Second)
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
	client  *client.Client
}

func (inf *info) key() string {
	return fmt.Sprintf("configurator|%s|%s|%s|%s", inf.env, inf.group, inf.project, inf.version)
}

func (inf *info) sync() error {
	k := inf.key()
	r, err := inf.client.GetOrRealtime(k, true)
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

func (inf *info) Map() map[string]interface{} {
	m := make(map[string]interface{})
	m[envKey] = inf.env
	m[groupKey] = inf.group
	m[projectKey] = inf.project
	m[versionKey] = inf.version
	m[masterKey] = inf.master
	m[slavesKey] = inf.slaves
	m[pathKey] = inf.Path
	return m
}
