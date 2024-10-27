package account

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	accountv1 "maxischmaxi/jstreams/account/v1"
	"maxischmaxi/jstreams/utils"

	"connectrpc.com/connect"
	"github.com/patrickmn/go-cache"
)

type AccountServer struct{}

var (
	accountCache = cache.New(5*time.Minute, 10*time.Minute)
)

func (s *AccountServer) GetAccountByGamenameAndTagline(
	_ context.Context,
	in *connect.Request[accountv1.GetAccountByGamenameAndTaglineRequest],
) (*connect.Response[accountv1.GetAccountByGamenameAndTaglineResponse], error) {
	path := fmt.Sprintf("/riot/account/v1/accounts/by-riot-id/%s/%s", in.Msg.Gamename, in.Msg.Tagline)
	uri, err := utils.GetRegionalRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	if v := utils.GetCachedValue(accountCache, uri); v != nil {
		return connect.NewResponse(v.(*accountv1.GetAccountByGamenameAndTaglineResponse)), nil
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var account accountv1.Account
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	val := &accountv1.GetAccountByGamenameAndTaglineResponse{
		Account: &account,
	}

	utils.SetCachedValue(accountCache, uri, val)

	return connect.NewResponse(val), nil
}

func (s *AccountServer) GetAccountByPuuid(
	_ context.Context,
	in *connect.Request[accountv1.GetAccountByPuuidRequest],
) (*connect.Response[accountv1.GetAccountByPuuidResponse], error) {
	path := fmt.Sprintf("/riot/account/v1/accounts/by-puuid/%s", in.Msg.Puuid)
	uri, err := utils.GetRegionalRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	if v := utils.GetCachedValue(accountCache, uri); v != nil {
		return connect.NewResponse(v.(*accountv1.GetAccountByPuuidResponse)), nil
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var account accountv1.Account
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	val := &accountv1.GetAccountByPuuidResponse{
		Account: &account,
	}

	utils.SetCachedValue(accountCache, uri, val)

	return connect.NewResponse(val), nil
}

func (s *AccountServer) GetAccountProfileIcon(
	_ context.Context,
	in *connect.Request[accountv1.GetAccountProfileIconRequest],
) (*connect.Response[accountv1.GetAccountProfileIconResponse], error) {
	path := fmt.Sprintf("/cdn/%s/img/profileicon/%v.png", in.Msg.PatchVersion, in.Msg.ProfileIconId)
	uri, err := utils.GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&accountv1.GetAccountProfileIconResponse{
		Url: uri.String(),
	}), nil
}
