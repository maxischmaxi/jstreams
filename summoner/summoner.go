package summoner

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	summonerv1 "maxischmaxi/jstreams/summoner/v1"
	"maxischmaxi/jstreams/utils"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/patrickmn/go-cache"
)

type SummonerServer struct{}

var (
	summonerCache = cache.New(5*time.Minute, 10*time.Minute)
)

func (s *SummonerServer) GetSummonerSpells(
	_ context.Context,
	in *connect.Request[summonerv1.GetSummonerSpellsRequest],
) (*connect.Response[summonerv1.GetSummonerSpellsResponse], error) {
	uri, err := utils.GetDDragonRoute(fmt.Sprintf("/cdn/%s/data/en_US/summoner.json", in.Msg.PatchVersion))

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	if v := utils.GetCachedValue(summonerCache, uri); v != nil {
		return connect.NewResponse(v.(*summonerv1.GetSummonerSpellsResponse)), nil
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var spells summonerv1.GetSummonerSpellsResponse
	err = json.NewDecoder(resp.Body).Decode(&spells)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	utils.SetCachedValue(summonerCache, uri, &spells)

	return connect.NewResponse(&spells), nil
}

func (s *SummonerServer) GetSummonerByPuuid(
	_ context.Context,
	in *connect.Request[summonerv1.GetSummonerByPuuidRequest],
) (*connect.Response[summonerv1.GetSummonerByPuuidResponse], error) {
	path := fmt.Sprintf("/lol/summoner/v4/summoners/by-puuid/%s", in.Msg.Puuid)
	uri, err := utils.GetPlatformRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	if v := utils.GetCachedValue(summonerCache, uri); v != nil {
		return connect.NewResponse(v.(*summonerv1.GetSummonerByPuuidResponse)), nil
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var summoner summonerv1.Summoner
	err = json.NewDecoder(resp.Body).Decode(&summoner)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	val := &summonerv1.GetSummonerByPuuidResponse{
		Summoner: &summoner,
	}

	utils.SetCachedValue(summonerCache, uri, val)

	return connect.NewResponse(val), nil
}
