package configurator

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConfigurator_AllSettings(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(AllSettings))
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Error(fmt.Errorf("期望的响应码为200，实际为%d", resp.StatusCode))
	}
	d, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	t.Logf("allsettings: %s\n", string(d))
}

func TestConfigurator_Get(t *testing.T) {
	t.Log(C.Get("redis.host"))
	t.Log(C.Get("redis.port"))
}
