// author: asydevc <asydev@163.com>
// date: 2021-02-13

package db

import (
	"github.com/asydevc/log/v2"
	"sync"
)

var (
	Config *configuration
)

func init() {
	new(sync.Once).Do(func() {
		log.Debug("init db package.")
		Config = new(configuration)
		Config.initialize()
	})
}
