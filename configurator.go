package configurator

import (
	"configurator/info"
	"encoding/json"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	C *configurator
)

const (
	infoKey       = "info"
	propertiesKey = "properties"
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

func (c *configurator) GetString(key string) string {
	return c.v.GetString(key)
}

func (c *configurator) GetBool(key string) bool {
	return c.v.GetBool(key)
}

func (c *configurator) GetInt(key string) int {
	return c.v.GetInt(key)
}

func (c *configurator) GetInt32(key string) int32 {
	return c.v.GetInt32(key)
}

func (c *configurator) GetInt64(key string) int64 {
	return c.v.GetInt64(key)
}

func (c *configurator) GetUint(key string) uint {
	return c.v.GetUint(key)
}

func (c *configurator) GetUint32(key string) uint32 {
	return c.v.GetUint32(key)
}

func (c *configurator) GetUint64(key string) uint64 {
	return c.v.GetUint64(key)
}

func (c *configurator) GetFloat64(key string) float64 {
	return c.v.GetFloat64(key)
}

func (c *configurator) GetTime(key string) time.Time {
	return c.v.GetTime(key)
}

func (c *configurator) GetDuration(key string) time.Duration {
	return c.v.GetDuration(key)
}

func (c *configurator) GetIntSlice(key string) []int {
	return c.v.GetIntSlice(key)
}

func (c *configurator) GetStringSlice(key string) []string {
	return c.v.GetStringSlice(key)
}

func (c *configurator) GetStringMap(key string) map[string]interface{} {
	return c.v.GetStringMap(key)
}

func (c *configurator) GetStringMapString(key string) map[string]string {
	return c.v.GetStringMapString(key)
}

func (c *configurator) GetStringMapStringSlice(key string) map[string][]string {
	return c.v.GetStringMapStringSlice(key)
}

func (c *configurator) GetSizeInBytes(key string) uint {
	return c.v.GetSizeInBytes(key)
}

func (c *configurator) AllSettings() map[string]interface{} {
	return c.v.AllSettings()
}

func AllSettings(w http.ResponseWriter, _ *http.Request) {
	m := make(map[string]interface{})
	m[infoKey] = info.I.Map()
	m[propertiesKey] = C.AllSettings()
	b, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
