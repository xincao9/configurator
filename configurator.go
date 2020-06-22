package configurator

import (
	"configurator/info"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var (
	C *configurator
)

func init() {
	var err error
	C, err = new()
	if err != nil {
		log.Fatalf("Fatal error configurator init: %v\n", err)
	}
}

func new() (*configurator, error) {
	v := viper.New()
	v.SetConfigName(info.FN)
	v.SetConfigType(info.Ext)
	v.AddConfigPath(info.I.Path)
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config file change: %s op: %d\n", in.Name, in.Op)
	})
	return &configurator{v: v}, nil
}

type configurator struct {
	v *viper.Viper
}

func (c *configurator) Get(key string) interface{} {
	return c.v.Get(key)
}
