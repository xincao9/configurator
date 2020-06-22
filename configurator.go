package configurator

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"github.com/xincao9/dkv/client/ms"
	"log"
	"os"
	"strings"
	"time"
)

var (
	C *configurator
)

const (
	fn      = "application"
	env     = "env"
	group   = "group"
	project = "project"
	version = "version"
	master  = "master"
	slaves  = "slaves"
)

func init() {
	inf := &info{}
	inf.env = os.Getenv(env)
	inf.group = os.Getenv(group)
	inf.project = os.Getenv(project)
	inf.version = os.Getenv(version)
	inf.master = os.Getenv(master)
	if os.Getenv(slaves) != "" {
		inf.slaves = strings.Split(os.Getenv(slaves), ",")
	}
	var err error
	C, err = New(inf)
	if err != nil {
		log.Fatalf("Fatal error configurator : %v\n", err)
	}
}

func New(inf *info) (*configurator, error) {
	v := viper.New()
	v.SetConfigName(fn)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	k := inf.key()
	var data string
	if len(inf.slaves) == 0 {
		c, err := ms.New(inf.master, time.Second)
		if err != nil {
			return nil, err
		}
		r, err := c.Get(k)
		if err != nil {
			return nil, err
		}
		if r.Code != 200 {
			return nil, fmt.Errorf("dkv response code: %d, message: %s\n", r.Code, r.Message)
		}
		log.Printf("dkv response: %v\n", r)
		data = r.KV.V
	} else {
		c, err := ms.NewMS(inf.master, inf.slaves, time.Second)
		if err != nil {
			return nil, err
		}
		r, err := c.Get(k)
		if err != nil {
			return nil, err
		}
		if r.Code != 200 {
			return nil, fmt.Errorf("dkv response code: %d, message: %s\n", r.Code, r.Message)
		}
		log.Printf("dkv response: %v\n", r)
		data = r.KV.V
	}
	err := v.ReadConfig(bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}
	return &configurator{v: v}, nill
}

type (
	configurator struct {
		v *viper.Viper
	}
	info struct {
		env     string
		group   string
		project string
		version string
		master  string
		slaves  []string
	}
)

func (inf *info) key() string {
	return fmt.Sprintf("configurator|%s|%s|%s|%s", inf.env, inf.group, inf.project, inf.version)
}

func (c *configurator) Get(key string) interface{} {
	return c.v.Get(key)
}

func (c *configurator) GetString(key string) string {
	return c.v.GetString(key)
}

func (c *configurator) GetInt(key string) int {
	return c.v.GetInt(key)
}

func (c *configurator) GetBool(key string) bool {
	return c.v.GetBool(key)
}

func (c *configurator) GetIntSlice(key string) []int {
	return c.v.GetIntSlice(key)
}