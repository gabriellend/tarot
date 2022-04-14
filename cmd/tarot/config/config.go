package config

import (
	"flag"
)

// type Config struct {
// 	Debug bool
// }

// func Load() *Config {
// 	var cfg = &Config{}

// 	flag.BoolVar(&cfg.Debug, "debug", false, "whether to run in debug mode")
// 	flag.Parse()

// 	return cfg
// }

const (
	defaultPort = "8080"
	flagPort    = "port"
	usagePort   = "port to listen and serve on"

	defaultStaticDir = "./web/static/"
	flagStaticDir    = "static"
	usageStaticDir   = "directory for static files"

	defaultTemplatesDir = "./web/templates/"
	flagTemplatesDir    = "templates"
	usageTemplatesDir   = "directory for templates"
)

type Config struct {
	Port         string //Why isn't this type pointer?
	StaticDir    string
	TemplatesDir string
}

func New() (*Config, error) {
	var (
		port      = flag.String(flagPort, defaultPort, usagePort)
		staticDir = flag.String(
			flagStaticDir, defaultStaticDir, usageStaticDir,
		)
		templatesDir = flag.String(
			flagTemplatesDir, defaultTemplatesDir, usageTemplatesDir,
		)
	)

	//Why no flag.Parse() here? How are flags useful if no parsing?

	return &Config{
		Port:         *port, //Why need to specify a pointer when flag.String returns a pointer?
		StaticDir:    *staticDir,
		TemplatesDir: *templatesDir,
	}, nil
}
