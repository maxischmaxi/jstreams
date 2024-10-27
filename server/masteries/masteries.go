package masteries

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	masteriesv1 "maxischmaxi/jstreams/masteries/v1"
	"maxischmaxi/jstreams/utils"
	"net/http"

	"connectrpc.com/connect"
)

type MasteriesServer struct{}

func (s *MasteriesServer) GetChampionMasteriesByPuuidByChampion(
	_ context.Context,
	in *connect.Request[masteriesv1.GetChampionMasteriesByChampionRequeset],
) (*connect.Response[masteriesv1.GetChampionMasteriesByChampionResponse], error) {
	return connect.NewResponse(&masteriesv1.GetChampionMasteriesByChampionResponse{}), nil
}

func (s *MasteriesServer) GetChampionMateriesByPuuidByChampion(
	_ context.Context,
	in *connect.Request[masteriesv1.GetChampionMasteriesByChampionRequeset],
) (*connect.Response[masteriesv1.GetChampionMasteriesByChampionResponse], error) {
	path := fmt.Sprintf("/lol/champion-mastery/v4/champion-masteries/by-puuid/%s/by-champion/%v", in.Msg.Puuid, in.Msg.ChampionId)
	uri, err := utils.GetPlatformRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
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

	return connect.NewResponse(&masteriesv1.GetChampionMasteriesByChampionResponse{
		ChampionMastery: &mastery,
	}), nil
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

	return connect.NewResponse(&masteriesv1.GetChampionMasteriesResponse{
		ChampionMasteries: masteries,
	}), nil
}
