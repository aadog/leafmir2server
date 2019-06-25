package conf

import (
	"encoding/json"
	"github.com/a97077088/leaf/log"
	"io/ioutil"
)

var Server struct {
	LogLevel     string
	LogPath      string
	WSAddr       string
	CertFile     string
	KeyFile      string
	LoginTCPPort int
	SelTCPPort   int
	GameTcpPort  int
	TcpAddr      string
	MaxConnNum   int
	ConsolePort  int
	ProfilePath  string
}

func init() {
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
