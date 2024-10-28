package version

import (
	"context"
	"encoding/json"
	"log"
	"maxischmaxi/jstreams/utils"
	versionv1 "maxischmaxi/jstreams/version/v1"
	"net/http"

	"github.com/patrickmn/go-cache"

	"connectrpc.com/connect"
)

var (
	versionCache = cache.New(cache.NoExpiration, cache.NoExpiration)
)

type VersionServer struct{}

func (s *VersionServer) GetCurrentVersion(
	_ context.Context,
	in *connect.Request[versionv1.GetCurrentVersionRequest],
) (*connect.Response[versionv1.GetCurrentVersionResponse], error) {
	uri, err := utils.GetDDragonRoute("/api/versions.json")

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	if v := utils.GetCachedValue(versionCache, uri); v != nil {
		return connect.NewResponse(v.(*versionv1.GetCurrentVersionResponse)), nil
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var versions []string
	err = json.NewDecoder(resp.Body).Decode(&versions)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	val := &versionv1.GetCurrentVersionResponse{
		Version: versions[0],
	}

	utils.SetCachedValue(versionCache, uri, val)

	return connect.NewResponse(val), nil
}

func (s *VersionServer) GetVersions(
	_ context.Context,
	in *connect.Request[versionv1.GetVersionsRequest],
) (*connect.Response[versionv1.GetVersionsResponse], error) {
	uri, err := utils.GetDDragonRoute("/api/versions.json")

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	if v := utils.GetCachedValue(versionCache, uri); v != nil {
		return connect.NewResponse(&versionv1.GetVersionsResponse{
			Versions: v.([]string),
		}), nil
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var versions []string
	err = json.NewDecoder(resp.Body).Decode(&versions)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	val := &versionv1.GetVersionsResponse{
		Versions: versions,
	}

	utils.SetCachedValue(versionCache, uri, val)

	return connect.NewResponse(val), nil
}
