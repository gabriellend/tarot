package config

import "flag"

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
	defaultPort = "8000"
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
	Port         string
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

	return &Config{
		Port:         *port,
		StaticDir:    *staticDir,
		TemplatesDir: *templatesDir,
	}, nil
}
