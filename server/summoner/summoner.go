package summoner

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	summonerv1 "maxischmaxi/jstreams/summoner/v1"
	"maxischmaxi/jstreams/utils"
	"net/http"

	"connectrpc.com/connect"
)

type SummonerServer struct{}

func (s *SummonerServer) GetSummonerSpells(
	_ context.Context,
	in *connect.Request[summonerv1.GetSummonerSpellsRequest],
) (*connect.Response[summonerv1.GetSummonerSpellsResponse], error) {
	uri, err := utils.GetDDragonRoute(fmt.Sprintf("/cdn/%s/data/en_US/summoner.json", in.Msg.PatchVersion))

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
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

	return connect.NewResponse(&summonerv1.GetSummonerByPuuidResponse{
		Summoner: &summoner,
	}), nil
}
