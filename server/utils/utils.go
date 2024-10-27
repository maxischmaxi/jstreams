package utils

import (
	"fmt"
	accountv1 "maxischmaxi/jstreams/account/v1"
	summonerv1 "maxischmaxi/jstreams/summoner/v1"
	"net/url"
	"os"
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
