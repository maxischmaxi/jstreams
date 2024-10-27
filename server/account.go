package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	accountv1 "maxischmaxi/jstreams/account/v1"

	"connectrpc.com/connect"
)

type accountServer struct{}

func (s *accountServer) GetAccountByGamenameAndTagline(
	_ context.Context,
	in *connect.Request[accountv1.GetAccountByGamenameAndTaglineRequest],
) (*connect.Response[accountv1.GetAccountByGamenameAndTaglineResponse], error) {
	path := fmt.Sprintf("/riot/account/v1/accounts/by-riot-id/%s/%s", in.Msg.Gamename, in.Msg.Tagline)
	uri, err := GetRegionalRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
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

	return connect.NewResponse(&accountv1.GetAccountByGamenameAndTaglineResponse{
		Account: &account,
	}), nil
}

func (s *accountServer) GetAccountByPuuid(
	_ context.Context,
	in *connect.Request[accountv1.GetAccountByPuuidRequest],
) (*connect.Response[accountv1.GetAccountByPuuidResponse], error) {
	path := fmt.Sprintf("/riot/account/v1/accounts/by-puuid/%s", in.Msg.Puuid)
	uri, err := GetRegionalRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
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

	return connect.NewResponse(&accountv1.GetAccountByPuuidResponse{
		Account: &account,
	}), nil
}

func (s *accountServer) GetAccountProfileIcon(
	_ context.Context,
	in *connect.Request[accountv1.GetAccountProfileIconRequest],
) (*connect.Response[accountv1.GetAccountProfileIconResponse], error) {
	var uri = fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img/profileicon/%v.png", in.Msg.PatchVersion, in.Msg.ProfileIconId)
	return connect.NewResponse(&accountv1.GetAccountProfileIconResponse{
		Url: uri,
	}), nil
}
