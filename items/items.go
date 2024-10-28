package items

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	itemsv1 "maxischmaxi/jstreams/items/v1"
	"maxischmaxi/jstreams/utils"
	"net/http"
	"strings"

	"connectrpc.com/connect"
	"github.com/patrickmn/go-cache"
)

type ItemsServer struct{}

var (
	itemsCache = cache.New(cache.NoExpiration, cache.NoExpiration)
)

func (s *ItemsServer) GetItemImageById(
	_ context.Context,
	in *connect.Request[itemsv1.GetItemImageByIdRequest],
) (*connect.Response[itemsv1.GetItemImageByIdResponse], error) {
	if in.Msg.ItemId == 0 {
		return connect.NewResponse(&itemsv1.GetItemImageByIdResponse{
			Url: "",
		}), nil
	}

	path := fmt.Sprintf("/cdn/%s/img/item/%v.png", in.Msg.PatchVersion, in.Msg.ItemId)
	uri, err := utils.GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&itemsv1.GetItemImageByIdResponse{
		Url: uri.String(),
	}), nil
}

func (s *ItemsServer) GetItemInformationById(
	_ context.Context,
	in *connect.Request[itemsv1.GetItemInformationByIdRequest],
) (*connect.Response[itemsv1.GetItemInformationByIdResponse], error) {
	path := fmt.Sprintf("/cdn/%s/data/en_US/item.json", in.Msg.PatchVersion)
	uri, err := utils.GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	if v := utils.GetCachedValue(itemsCache, uri); v != nil {
		infos := v.(*itemsv1.GetBaseItemsResponse)

		return connect.NewResponse(&itemsv1.GetItemInformationByIdResponse{
			Item: infos.Data[strings.ToLower(utils.String(in.Msg.ItemId))],
		}), nil
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var infos *itemsv1.GetBaseItemsResponse
	err = json.NewDecoder(resp.Body).Decode(&infos)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	utils.SetCachedValue(itemsCache, uri, infos)

	return connect.NewResponse(&itemsv1.GetItemInformationByIdResponse{
		Item: infos.Data[strings.ToLower(utils.String(in.Msg.ItemId))],
	}), nil
}

func (s *ItemsServer) GetBaseItems(
	_ context.Context,
	in *connect.Request[itemsv1.GetBaseItemsRequest],
) (*connect.Response[itemsv1.GetBaseItemsResponse], error) {
	path := fmt.Sprintf("/cdn/%s/data/en_US/item.json", in.Msg.PatchVersion)
	uri, err := utils.GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	if v := utils.GetCachedValue(itemsCache, uri); v != nil {
		return connect.NewResponse(v.(*itemsv1.GetBaseItemsResponse)), nil
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var response itemsv1.GetBaseItemsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	utils.SetCachedValue(itemsCache, uri, &response)

	return connect.NewResponse(&response), nil
}
