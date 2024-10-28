package champions

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	championsv1 "maxischmaxi/jstreams/champions/v1"
	"maxischmaxi/jstreams/utils"
	"net/http"
	"strconv"

	"connectrpc.com/connect"
	"github.com/patrickmn/go-cache"
)

type ChampionsServer struct{}

var (
	championsCache = cache.New(cache.NoExpiration, cache.NoExpiration)
)

func (s *ChampionsServer) GetChampions(
	_ context.Context,
	in *connect.Request[championsv1.GetChampionsRequest],
) (*connect.Response[championsv1.GetChampionsResponse], error) {
	path := fmt.Sprintf("/cdn/%s/data/en_US/champion.json", in.Msg.PatchVersion)
	uri, err := utils.GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	if v := utils.GetCachedValue(championsCache, uri); v != nil {
		return connect.NewResponse(v.(*championsv1.GetChampionsResponse)), nil
	}

	var response championsv1.GetChampionsResponse
	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	utils.SetCachedValue(championsCache, uri, &response)

	return connect.NewResponse(&response), nil
}

func (s *ChampionsServer) GetChampByChampId(
	_ context.Context,
	in *connect.Request[championsv1.GetChampByChampIdRequest],
) (*connect.Response[championsv1.GetChampByChampIdResponse], error) {
	path := fmt.Sprintf("/cdn/%s/data/en_US/champion.json", in.Msg.PatchVersion)
	uri, err := utils.GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	if v := utils.GetCachedValue(championsCache, uri); v != nil {
		champId := strconv.FormatInt(int64(in.Msg.ChampId), 10)

		for _, champ := range v.(*championsv1.GetChampionsResponse).Data {
			if champ.Key == champId {
				return connect.NewResponse(&championsv1.GetChampByChampIdResponse{
					Champion: champ,
				}), nil
			}
		}
	}

	var response championsv1.GetChampionsResponse
	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	utils.SetCachedValue(championsCache, uri, &response)

	champId := strconv.FormatInt(int64(in.Msg.ChampId), 10)

	for _, champ := range response.Data {
		if champ.Key == champId {
			return connect.NewResponse(&championsv1.GetChampByChampIdResponse{
				Champion: champ,
			}), nil
		}
	}

	return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("champion not found"))
}
