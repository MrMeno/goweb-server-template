package Utils

import (
	"MetaWebServer/DataReflect/Config"
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

/*
register a global key for action
*/
var (
	GB_CACHE   = cache.New(24*time.Hour, 7*24*time.Hour)
	DB_G_KEY   = "db"
	USR_G_KEY  = "user"
	AU_G_KEY   = "auth"
	PROP_G_KEY = "props"
)

type CacheManager interface {
	Add(string, interface{})
	Remove(string)
	Get(string) (interface{}, bool)
}
type CacheService struct {
}

func (c *CacheService) Add(key string, t interface{}) {
	GB_CACHE.Set(key, t, cache.DefaultExpiration)
}

func (c *CacheService) Remove(k string) {
	GB_CACHE.Delete(k)
}
func (c *CacheService) Get(k string) (interface{}, bool) {
	switch k {
	case DB_G_KEY:
		temp, res := GB_CACHE.Get(k)
		if res {
			return temp, true
		} else {
			myConfigIntf := CFG[Config.DBConfig]{}
			temp, err := myConfigIntf.ReadFromYaml("./Config/datasource.yaml")
			if err != nil {
				return nil, false
			} else {
				c.Add(DB_G_KEY, *temp)
				log.Printf("DB Cache Loaded")
				return temp, true
			}
		}
	case AU_G_KEY:
		temp, res := GB_CACHE.Get(AU_G_KEY)
		if res {
			return temp, true
		} else {
			myConfigIntf := CFG[Config.AuthConfig]{}
			temp, err := myConfigIntf.ReadFromYaml("./Config/auth.yaml")
			if err != nil {
				return nil, false
			} else {
				c.Add(AU_G_KEY, *temp)
				log.Printf("Auth Cache Loaded")
				return temp, true
			}
		}
	case PROP_G_KEY:
		temp, res := GB_CACHE.Get(PROP_G_KEY)
		if res {
			return temp, true
		} else {
			myConfigIntf := CFG[Config.PropsConfig]{}
			temp, err := myConfigIntf.ReadFromYaml("./Config/properties.yaml")
			if err != nil {
				return nil, false
			} else {
				c.Add(PROP_G_KEY, *temp)
				log.Printf("Property Cache Loaded")
				return temp, true
			}
		}
	default:
		return nil, false
	}
}
