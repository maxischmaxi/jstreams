package assets

import (
	"context"
	"fmt"
	"log"
	assetsv1 "maxischmaxi/jstreams/assets/v1"
	"maxischmaxi/jstreams/utils"
	"strings"

	"connectrpc.com/connect"
)

type AssetsServer struct{}

func (s *AssetsServer) GetRuneIcon(
	_ context.Context,
	in *connect.Request[assetsv1.GetRuneIconRequest],
) (*connect.Response[assetsv1.GetRuneIconResponse], error) {
	return connect.NewResponse(&assetsv1.GetRuneIconResponse{
		Url: "",
	}), nil
}

func (s *AssetsServer) GetItemAssetUrl(
	_ context.Context,
	in *connect.Request[assetsv1.GetItemAssetUrlRequest],
) (*connect.Response[assetsv1.GetItemAssetUrlResponse], error) {
	return connect.NewResponse(&assetsv1.GetItemAssetUrlResponse{
		Url: "",
	}), nil
}

func (s *AssetsServer) GetRuneIconRequest(
	_ context.Context,
	in *connect.Request[assetsv1.GetRuneIconRequest],
) (*connect.Response[assetsv1.GetRuneIconResponse], error) {
	path := fmt.Sprintf("/cdn/img/%s/img/%v", in.Msg.PatchVersion, in.Msg.Style)
	uri, err := utils.GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetRuneIconResponse{
		Url: uri.String(),
	}), nil
}

func (s *AssetsServer) GetSummonerSpellIcon(
	_ context.Context,
	in *connect.Request[assetsv1.GetSummonerSpellIconRequest],
) (*connect.Response[assetsv1.GetSummonerSpellIconResponse], error) {
	path := fmt.Sprintf("/cdn/%s/img/spell/%s", in.Msg.PatchVersion, in.Msg.Image.Full)
	uri, err := utils.GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetSummonerSpellIconResponse{
		Url: uri.String(),
	}), nil
}

func (s *AssetsServer) GetItemAssetUrlRequest(
	_ context.Context,
	in *connect.Request[assetsv1.GetItemAssetUrlRequest],
) (*connect.Response[assetsv1.GetItemAssetUrlResponse], error) {
	path := fmt.Sprintf("/cdn/%s/data/en_US/%s.json", in.Msg.PatchVersion, in.Msg.ItemName)
	uri, err := utils.GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetItemAssetUrlResponse{
		Url: uri.String(),
	}), nil
}

func (s *AssetsServer) GetSpellAssetUrl(
	_ context.Context,
	in *connect.Request[assetsv1.GetSpellAssetUrlRequest],
) (*connect.Response[assetsv1.GetSpellAssetUrlResponse], error) {
	uri, err := utils.GetDDragonRoute(fmt.Sprintf("/cdn/%s/img/spell/%s.png", in.Msg.PatchVersion, in.Msg.SpellName))

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetSpellAssetUrlResponse{
		Url: uri.String(),
	}), nil
}

func (s *AssetsServer) GetChampionAbilityAssetUrl(
	_ context.Context,
	in *connect.Request[assetsv1.GetChampionAbilityAssetUrlRequest],
) (*connect.Response[assetsv1.GetChampionAbilityAssetUrlResponse], error) {
	var championName = strings.ReplaceAll(in.Msg.ChampionName, " ", "")
	uri, err := utils.GetDDragonRoute(fmt.Sprintf("/cdn/%s/img/champion/%s_%v.png", in.Msg.PatchVersion, championName, in.Msg.AbilityNumber))

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetChampionAbilityAssetUrlResponse{
		Url: uri.String(),
	}), nil
}

func (s *AssetsServer) GetChampionPassiveAssetUrl(
	_ context.Context,
	in *connect.Request[assetsv1.GetChampionPassiveAssetUrlRequest],
) (*connect.Response[assetsv1.GetChampionPassiveAssetUrlResponse], error) {
	var championName = strings.ReplaceAll(in.Msg.ChampionName, " ", "")
	path := fmt.Sprintf("/cdn/%s/img/passive/%s_P.png", in.Msg.PatchVersion, championName)
	uri, err := utils.GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetChampionPassiveAssetUrlResponse{
		Url: uri.String(),
	}), nil
}

func (s *AssetsServer) GetChampionSquareAssetUrl(
	_ context.Context,
	in *connect.Request[assetsv1.GetChampionSquareAssetUrlRequest],
) (*connect.Response[assetsv1.GetChampionSquareAssetUrlResponse], error) {
	uri, err := utils.GetDDragonRoute(fmt.Sprintf("/cdn/%s/img/champion/%s", in.Msg.PatchVersion, in.Msg.ChampionName))

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetChampionSquareAssetUrlResponse{
		Url: uri.String(),
	}), nil
}

func (s *AssetsServer) GetChampionLoadingScreenAssetUrl(
	_ context.Context,
	in *connect.Request[assetsv1.GetChampionLoadingScreenAssetUrlRequest],
) (*connect.Response[assetsv1.GetChampionLoadingScreenAssetUrlResponse], error) {
	var championName = strings.ReplaceAll(in.Msg.ChampionName, " ", "")
	uri, err := utils.GetDDragonRoute(fmt.Sprintf("/cdn/img/champion/loading/%s_%v.jpg", championName, in.Msg.SkinNumber))

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetChampionLoadingScreenAssetUrlResponse{
		Url: uri.String(),
	}), nil
}

func (s *AssetsServer) GetChampionSplashAssetUrl(
	_ context.Context,
	in *connect.Request[assetsv1.GetChampionSplashAssetUrlRequest],
) (*connect.Response[assetsv1.GetChampionSplashAssetUrlResponse], error) {
	var championName = strings.ReplaceAll(in.Msg.ChampionName, " ", "")
	uri, err := utils.GetDDragonRoute(fmt.Sprintf("/cdn/img/champion/splash/%s_%v.jpg", championName, in.Msg.SkinNumber))

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetChampionSplashAssetUrlResponse{
		Url: uri.String(),
	}), nil
}
