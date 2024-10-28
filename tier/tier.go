package tier

import (
	"context"
	"fmt"
	"log"
	tierv1 "maxischmaxi/jstreams/tier/v1"
	"maxischmaxi/jstreams/utils"

	"connectrpc.com/connect"
)

type TierServer struct{}

func (s *TierServer) GetTierIcon(
	_ context.Context,
	in *connect.Request[tierv1.GetTierIconRequest],
) (*connect.Response[tierv1.GetTierIconResponse], error) {
	path := fmt.Sprintf("/cdn/%s/img/tier/%s.png", in.Msg.PatchVersion, in.Msg.Tier)
	uri, err := utils.GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&tierv1.GetTierIconResponse{
		Url: uri.String(),
	}), nil
}
