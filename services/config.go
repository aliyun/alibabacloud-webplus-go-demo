package services

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Debug      bool              `yaml:"debug,omitempty"`
	Logging    *LoggingConfig    `yaml:"logging,omitempty"`
	Site       *SiteConfig       `yaml:"site,omitempty"`
	Quickstart *QuickStartConfig `yaml:"quickstart,omitempty"`
	App        *App              `yaml:"app,omitempty"`
	Env        *Env              `yaml:"env,omitempty"`
	Next       *NextConfig       `yaml:"next,omitempty"`
	WebPlus    *WebPlus          `yaml:"webplus,omitempty"`
}

type LoggingConfig struct {
	Level string `yaml:"level,omitempty"`
}

type SiteConfig struct {
	Id string `yaml:"id,omitempty"`
}

type QuickStartConfig struct {
	Doc  *QuickStartDoc  `yaml:"doc,omitempty"`
	Repo *QuickStartRepo `yaml:"repo,omitempty"`
}

type QuickStartDoc struct {
	Url string `yaml:"url,omitempty"`
}

type QuickStartRepo struct {
	Name string `yaml:"name,omitempty"`
	Url  string `yaml:"url,omitempty"`
}

type App struct {
	Url string `yaml:"url,omitempty"`
}

type Env struct {
	Url string `yaml:"url,omitempty"`
}

type NextConfig struct {
	Step *Step `yaml:"step,omitempty"`
}

type Step struct {
	Show    bool     `yaml:"show,omitempty"`
	Package *Package `yaml:"package,omitempty"`
}

type Package struct {
	Url string `yaml:"url,omitempty"`
}

type WebPlus struct {
	Console *Console `yaml:"console,omitempty"`
}

type Console struct {
	Url string `yaml:"url,omitempty"`
}

func LoadConfig() *Config {
	config := Config{}
	yamlFile, _ := ioutil.ReadFile("conf/config.yml")
	yaml.Unmarshal(yamlFile, &config)
	return &config
}
