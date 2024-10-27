package champions

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	championsv1 "maxischmaxi/jstreams/champions/v1"
	"maxischmaxi/jstreams/utils"
	"net/http"

	"connectrpc.com/connect"
)

type ChampionsServer struct{}

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

	return connect.NewResponse(&response), nil
}
