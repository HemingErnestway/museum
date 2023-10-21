package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
)

type Server struct {
	Port    string
	Timeout string
}

type Service struct {
	Method string
	Url    string
}

type Config struct {
	Server
	Api map[string]Service
}

var config Config

func Get() *Config {
	return &config
}

var isEnv = regexp.MustCompile(`\$\{(.*?)}`)

func init() {
	file, err := os.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
	//if isEnv.MatchString(config.Server.Port) {
	//	find := isEnv.FindStringSubmatch(config.Server.Port)
	//	fmt.Println(find)
	//	sub := strings.SplitN(find[1], ":", 1)
	//	envOs := os.Getenv(sub[0])
	//	if len(envOs) > 0 {
	//		config.Server.Port = envOs
	//	} else {
	//		config.Server.Port = sub[1]
	//	}
	//}
}
