package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/yurist-85/gosignit/internal/helpers"
)

const (
	EnvHost = "HOST"
	EnvPort = "PORT"
	EnvSeed = "SEED"
)

const (
	DefaultHost = "localhost"
	DefaultPort = 3000
	DefaultSeed = ""
)

type ConfigInterface interface {
	HostAndPort() string
	Host() string
	Port() int64
	Seed() string
}

// Configuration stores app settings. Use separated struct for JSON unmarshalling purposes.
type Configuration struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
	Seed string `json:"seed"`
}

// Config contains app Configuration.
type Config struct {
	file   string
	config *Configuration
}

// NewConfig creates new instance of Config and load settings.
func NewConfig() ConfigInterface {
	c := &Config{}
	c.init()

	return c
}

func (c *Config) init() {
	c.config = &Configuration{}
	c.loadEnv()
	c.readParams()

	if c.config.Seed == "" {
		logrus.Fatalln("Seed phrase is not set")
	}
}

// loadEnv reads environment variables or sets defaults.
func (c *Config) loadEnv() {
	logrus.Debugln("Load environment configuration...")
	c.config.Host = helpers.GetEnv(EnvHost, DefaultHost)
	c.config.Port = helpers.GetEnvInt64(EnvPort, DefaultPort)
	c.config.Seed = helpers.GetEnv(EnvSeed, DefaultSeed)

	if c.config.Port < 1024 || c.config.Port > 49151 {
		// TODO Warn about invalid port settings
		// The Well Known Ports (0-1023) – which are reserved for the operating system and core services.
		// The Registered Ports (1024-49151) – which can be used by applications, specific services, and users.
		// The Dynamic and/or Private Ports (49152-65535)
		c.config.Port = 3000
	}

	logrus.WithField("env", c.config).Debugln("Config")
}

// readParams reads configuration from program flags
func (c *Config) readParams() {
	logrus.Debugln("Read configuration from params...")

	host := flag.String("host", "", "Host")
	port := flag.Int64("port", 0, "Port")
	seed := flag.String("seed", "", "Seed phrase")
	filename := flag.String("config", "", "Configuration file")
	flag.Parse()

	if filename != nil && *filename != "" {
		c.readFile(*filename)

		return
	}

	if host != nil && *host != "" {
		c.config.Host = *host
	}

	if port != nil && *port != 0 {
		c.config.Port = *port
	}

	if seed != nil && *seed != "" {
		c.config.Seed = *seed
	}

	logrus.WithField("params", c.config).Debugln("Config after params")
}

// readFile reads configuration from given json file
func (c *Config) readFile(filename string) {
	logrus.Debugln("Load configuration from file...")
	logrus.Debugf("Config file: %s", filename)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		logrus.Fatalln("Config file doesn't exist")

		return
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.WithError(err).Fatalln("Can't read config file")
	}

	if err := json.Unmarshal(data, c.config); err != nil {
		// TODO Warn about invalid file format
		logrus.WithError(err).Fatalln("Corrupted config file. It must be valid JSON.")
	}

	logrus.WithField("json", c.config).Debugln("Config after json")
}

// HostAndPort returns host with port concatenated.
func (c Config) HostAndPort() string {
	return fmt.Sprintf("%s:%d", c.config.Host, c.config.Port)
}

// Host returns host value.
func (c Config) Host() string {
	return c.config.Host
}

// Port returns port value.
func (c Config) Port() int64 {
	return c.config.Port
}

// Seed returns seed value.
func (c Config) Seed() string {
	return c.config.Seed
}
