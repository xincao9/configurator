package config

import (
    "configurator/api/component/constant"
    "flag"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    "log"
    "net/http"
    "os"
    "os/exec"
    "strings"
)

var (
    C *viper.Viper
)

func init() {
    d := flag.Bool("d", false, "run app as a daemon with -d=true")
    c := flag.String("conf", "configurator-api.yaml", "configure file")
    if flag.Parsed() == false {
        flag.Parse()
    }
    if *d {
        args := os.Args[1:]
        i := 0
        for ; i < len(args); i++ {
            if args[i] == "-d=true" {
                args[i] = "-d=false"
                break
            }
        }
        cmd := exec.Command(os.Args[0], args...)
        cmd.Start()
        fmt.Println("[PID]", cmd.Process.Pid)
        os.Exit(0)
    }
    C = viper.New()
    for _, t := range []string{"yaml", "yml"} {
        if strings.HasSuffix(*c, t) {
            i := strings.LastIndex(*c, t)
            *c = string([]byte(*c)[:i-1])
        }
    }
    C.SetConfigName(*c)
    C.SetConfigType("yaml")
    C.AddConfigPath("./resources/conf")
    C.AddConfigPath("/tmp/configurator/log")
    C.AddConfigPath("/usr/local/configurator/conf")
    C.SetDefault(constant.LoggerDir, "/tmp/configurator/log")
    C.SetDefault(constant.LoggerLevel, "debug")
    C.SetDefault(constant.ServerMode, "debug")
    C.SetDefault(constant.ServerPort, 8080)
    C.SetDefault(constant.ServerCorsAccessControlAllowOrigin, "http://localhost:8081")
    C.SetDefault(constant.ManagerServerPort, 8090)
    C.SetDefault(constant.DBDataSourceName, "root:asdf@tcp(localhost:3306)/configurator?charset=utf8&parseTime=true")
    C.SetDefault(constant.DKVAddress, "localhost:9090")
    err := C.ReadInConfig()
    if err != nil {
        log.Fatalf("Fatal error conf : %v\n", err)
    }
}

func Route(engine *gin.Engine) {
    engine.GET("/config", func(c *gin.Context) {
        c.JSON(http.StatusOK, C.AllSettings())
    })
}
