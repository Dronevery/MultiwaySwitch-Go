package MultiwaySwitch

import (
	"code.google.com/p/log4go"
	"github.com/Unknwon/goconfig"
	"log"
)

const (
	config_file = "config.ini"
)

type level int

const (
	FATAL level = iota
	ERROR
	WARMING
	INFO
)

var cfg *goconfig.ConfigFile

func initConfig() {
	var err error

	cfg, err = goconfig.LoadConfigFile(config_file)
	if err != nil {
		log.Fatalf("CANNOT load config file(%s) : %s\n", config_file, err)
	}
	/*
		role, err := cfg.GetValue("common", "role")
		if err != nil {
			log.Fatalf("CONNOT read role field from config file: %s\n", err)
		}
		config, err := cfg.GetSection("common")
		if err != nil {
			log.Fatalf("CONNOT read common session from config file: %s\n", err)
		}
		config_role, err := cfg.GetSection(role)
		if err != nil {
			log.Fatalf("CONNOT read %s session from config file: %s\n", role, err)
		}
		config = make(map[string]string)
		if err := mergo.Merge(&config, config_role); err != nil {
			log.Fatalf("merge config in difference session failed", role, err)
		}
	*/
}

func config(section, key string, lvl level) string {
	value, err := cfg.GetValue(section, key)
	if err != nil {
		switch lvl {
		case FATAL:
			logger.Critical("Error Read config:", err.Error())
			panic(err.Error())
		case ERROR:
			logger.Error("Error Read config:", err.Error())
		case WARMING:
			logger.Warn("Error Read config:", err.Error())
		case INFO:
			logger.Info("Error Read config:", err.Error())
		}
		return ""
	}
	return value
}
func configCommon(key string, lvl level) string {
	return config("common", key, lvl)
}
func configServer(key string, lvl level) string {
	return config("server", key, lvl)
}
func configClient(key string, lvl level) string {
	return config("client", key, lvl)
}
