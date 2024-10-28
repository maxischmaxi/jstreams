package utils

import (
	"fmt"
	accountv1 "maxischmaxi/jstreams/account/v1"
	summonerv1 "maxischmaxi/jstreams/summoner/v1"
	"net/url"
	"os"

	"github.com/patrickmn/go-cache"
)

func GetPlatformRoute(region summonerv1.PlatformRoutingValues, route string) (url.URL, error) {
	var uri = fmt.Sprintf("https://%s.api.riotgames.com%s", region, route)
	u, err := url.Parse(uri)

	if err != nil {
		return url.URL{}, err
	}

	q := u.Query()
	q.Set("api_key", os.Getenv("RIOT_API_KEY"))
	u.RawQuery = q.Encode()

	return *u, nil
}

func GetDDragonRoute(route string) (url.URL, error) {
	var uri = fmt.Sprintf("https://ddragon.leagueoflegends.com%s", route)
	u, err := url.Parse(uri)

	if err != nil {
		return url.URL{}, err
	}

	return *u, nil
}

func GetRegionalRoute(region accountv1.RegionalRoutingValues, route string) (url.URL, error) {
	var uri = fmt.Sprintf("https://%s.api.riotgames.com%s", region, route)
	u, err := url.Parse(uri)

	if err != nil {
		return url.URL{}, err
	}

	q := u.Query()
	q.Set("api_key", os.Getenv("RIOT_API_KEY"))
	u.RawQuery = q.Encode()

	return *u, nil
}

func GetCachedValue(c *cache.Cache, uri url.URL) interface{} {
	foo, found := c.Get(uri.String())

	if found {
		return foo
	}

	return nil
}

func SetCachedValue(c *cache.Cache, uri url.URL, value interface{}) {
	c.Set(uri.String(), value, cache.DefaultExpiration)
}

func String(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
