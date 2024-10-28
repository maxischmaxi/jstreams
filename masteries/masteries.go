package masteries

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	masteriesv1 "maxischmaxi/jstreams/masteries/v1"
	"maxischmaxi/jstreams/utils"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/patrickmn/go-cache"
)

type MasteriesServer struct{}

var (
	masteriesCache = cache.New(5*time.Minute, 10*time.Minute)
)

func (s *MasteriesServer) GetChampionMasteriesByPuuidByChampion(
	_ context.Context,
	in *connect.Request[masteriesv1.GetChampionMasteriesByChampionRequeset],
) (*connect.Response[masteriesv1.GetChampionMasteriesByChampionResponse], error) {
	path := fmt.Sprintf("/lol/champion-mastery/v4/champion-masteries/by-puuid/%s/by-champion/%v", in.Msg.Puuid, in.Msg.ChampionId)
	uri, err := utils.GetPlatformRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	if v := utils.GetCachedValue(masteriesCache, uri); v != nil {
		return connect.NewResponse(v.(*masteriesv1.GetChampionMasteriesByChampionResponse)), nil
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var mastery masteriesv1.ChampionMastery
	err = json.NewDecoder(resp.Body).Decode(&mastery)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	val := &masteriesv1.GetChampionMasteriesByChampionResponse{
		ChampionMastery: &mastery,
	}

	utils.SetCachedValue(masteriesCache, uri, val)

	return connect.NewResponse(val), nil
}

func (s *MasteriesServer) GetChampionMasteriesByPuuid(
	_ context.Context,
	in *connect.Request[masteriesv1.GetChampionMasteriesRequeset],
) (*connect.Response[masteriesv1.GetChampionMasteriesResponse], error) {
	path := fmt.Sprintf("/lol/champion-mastery/v4/champion-masteries/by-puuid/%s", in.Msg.Puuid)
	uri, err := utils.GetPlatformRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	if v := utils.GetCachedValue(masteriesCache, uri); v != nil {
		return connect.NewResponse(v.(*masteriesv1.GetChampionMasteriesResponse)), nil
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var masteries []*masteriesv1.ChampionMastery

	err = json.NewDecoder(resp.Body).Decode(&masteries)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	val := &masteriesv1.GetChampionMasteriesResponse{
		ChampionMasteries: masteries,
	}

	utils.SetCachedValue(masteriesCache, uri, val)

	return connect.NewResponse(val), nil
}
