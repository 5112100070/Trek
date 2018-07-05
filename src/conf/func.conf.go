package conf

import (
	"log"

	"github.com/ruizu/gcfg"
)

func ReadConfig(filePath string) (Config, bool) {
	var c Config
	if err := gcfg.ReadFileInto(&c, filePath); err != nil {
		log.Printf("%s\n", err)
		return Config{}, false
	}
	return c, true
}
