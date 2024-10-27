package entries

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	entriesv1 "maxischmaxi/jstreams/entries/v1"
	"maxischmaxi/jstreams/utils"
	"net/http"

	"connectrpc.com/connect"
)

type EntriesServer struct{}

func (s *EntriesServer) GetEntriesBySummoner(
	_ context.Context,
	in *connect.Request[entriesv1.GetEntriesBySummonerRequest],
) (*connect.Response[entriesv1.GetEntriesBySummonerResponse], error) {
	path := fmt.Sprintf("/lol/league/v4/entries/by-summoner/%s", in.Msg.SummonerId)
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

	var entries []*entriesv1.SummonerEntry
	err = json.NewDecoder(resp.Body).Decode(&entries)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&entriesv1.GetEntriesBySummonerResponse{
		Entries: entries,
	}), nil
}
