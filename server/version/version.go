package version

import (
	"context"
	"encoding/json"
	"log"
	"maxischmaxi/jstreams/utils"
	versionv1 "maxischmaxi/jstreams/version/v1"
	"net/http"

	"connectrpc.com/connect"
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

	return connect.NewResponse(&versionv1.GetCurrentVersionResponse{
		Version: versions[0],
	}), nil
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

	return connect.NewResponse(&versionv1.GetVersionsResponse{
		Versions: versions,
	}), nil
}
