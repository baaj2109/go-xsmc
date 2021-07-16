package sysinit

import (
	"hellobeego/utils"
	"time"

	cache "github.com/patrickmn/go-cache"
)

func init() {
	utils.Cache = cache.New(60*time.Minute, 120*time.Minute)
	initDB()
}
