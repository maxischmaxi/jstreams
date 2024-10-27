package matches

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	matchesv1 "maxischmaxi/jstreams/matches/v1"
	"maxischmaxi/jstreams/utils"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/patrickmn/go-cache"
)

type MatchesServer struct{}

var (
	matchesCache = cache.New(5*time.Minute, 10*time.Minute)
)

func (s *MatchesServer) GetMatchIdsByPuuid(
	_ context.Context,
	in *connect.Request[matchesv1.GetMatchIdsByPuuidRequest],
) (*connect.Response[matchesv1.GetMatchIdsByPuuidResponse], error) {
	path := fmt.Sprintf("/lol/match/v5/matches/by-puuid/%s/ids", in.Msg.Puuid)
	uri, err := utils.GetRegionalRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	if v := utils.GetCachedValue(matchesCache, uri); v != nil {
		return connect.NewResponse(v.(*matchesv1.GetMatchIdsByPuuidResponse)), nil
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var matchIds []string
	err = json.NewDecoder(resp.Body).Decode(&matchIds)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	val := &matchesv1.GetMatchIdsByPuuidResponse{
		MatchIds: matchIds,
	}

	utils.SetCachedValue(matchesCache, uri, val)

	return connect.NewResponse(val), nil
}

func (s *MatchesServer) GetMatchByMatchId(
	_ context.Context,
	in *connect.Request[matchesv1.GetMatchByMatchIdRequest],
) (*connect.Response[matchesv1.GetMatchByMatchIdResponse], error) {
	path := fmt.Sprintf("/lol/match/v5/matches/%s", in.Msg.MatchId)
	uri, err := utils.GetRegionalRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	if v := utils.GetCachedValue(matchesCache, uri); v != nil {
		return connect.NewResponse(v.(*matchesv1.GetMatchByMatchIdResponse)), nil
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var match matchesv1.Match
	err = json.NewDecoder(resp.Body).Decode(&match)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	val := &matchesv1.GetMatchByMatchIdResponse{
		Match: &match,
	}

	utils.SetCachedValue(matchesCache, uri, val)

	return connect.NewResponse(val), nil
}

func (s *MatchesServer) GetMatchTimeline(
	_ context.Context,
	in *connect.Request[matchesv1.GetMatchTimelineRequest],
) (*connect.Response[matchesv1.GetMatchTimelineResponse], error) {
	path := fmt.Sprintf("/lol/match/v5/matches/%s/timeline", in.Msg.MatchId)
	uri, err := utils.GetRegionalRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	if v := utils.GetCachedValue(matchesCache, uri); v != nil {
		return connect.NewResponse(v.(*matchesv1.GetMatchTimelineResponse)), nil
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var timeline matchesv1.Timeline
	err = json.NewDecoder(resp.Body).Decode(&timeline)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	val := &matchesv1.GetMatchTimelineResponse{
		Timeline: &timeline,
	}

	utils.SetCachedValue(matchesCache, uri, val)

	return connect.NewResponse(val), nil
}
