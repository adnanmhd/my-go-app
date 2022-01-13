package common

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	Port            string `default:":1323"`
	RootURL         string `split_words:"true" default:"/backend-menu"`
	RootAPI         string `split_words:"true" default:"/api/engine"`
	LogLevel        string `split_words:"true" default:"debug"`
	ConfigFile      string `split_words:"true" default:"config/config_assembly.json"`
	KeyFile         string `split_words:"true" default:""`
	CertificateFile string `split_words:"true" default:""`
	VersionApp      string `split_words:"true" default:"0.0.1"`
	NameApp         string `split_words:"true" default:"BE Menu"`
}

type ErrorMessage struct {
	HTTPStatus int      `json:"http"`
	Message    []string `json:"message"`
}

type VersionName struct {
	Version string `json:"version"`
	Name    string `json:"name"`
}

var Config Configuration
var FileConfig fileConfig

type fileConfig struct {
	Database struct {
		MySql struct {
			Name     string `json:"name"`
			Host     string `json:"host"`
			Port     string `json:"port"`
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"mysql"`
	} `json:"database"`
}

var VersionApp VersionName
var ErrorMessages map[string]ErrorMessage

func InitConfig() {
	err := envconfig.Process("PC", &Config)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.WithField("config", Config).Info("Config successfully loaded")

	VersionApp.Name = Config.NameApp
	VersionApp.Version = Config.VersionApp

	file, err := ioutil.ReadFile(Config.ConfigFile)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal([]byte(file), &FileConfig)
}
