// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: assets/v1/assets.proto

package assetsv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "maxischmaxi/jstreams/assets/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// AssetsServiceName is the fully-qualified name of the AssetsService service.
	AssetsServiceName = "assets.AssetsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AssetsServiceGetRuneIconProcedure is the fully-qualified name of the AssetsService's GetRuneIcon
	// RPC.
	AssetsServiceGetRuneIconProcedure = "/assets.AssetsService/GetRuneIcon"
	// AssetsServiceGetSummonerSpellIconProcedure is the fully-qualified name of the AssetsService's
	// GetSummonerSpellIcon RPC.
	AssetsServiceGetSummonerSpellIconProcedure = "/assets.AssetsService/GetSummonerSpellIcon"
	// AssetsServiceGetItemAssetUrlProcedure is the fully-qualified name of the AssetsService's
	// GetItemAssetUrl RPC.
	AssetsServiceGetItemAssetUrlProcedure = "/assets.AssetsService/GetItemAssetUrl"
	// AssetsServiceGetSpellAssetUrlProcedure is the fully-qualified name of the AssetsService's
	// GetSpellAssetUrl RPC.
	AssetsServiceGetSpellAssetUrlProcedure = "/assets.AssetsService/GetSpellAssetUrl"
	// AssetsServiceGetChampionAbilityAssetUrlProcedure is the fully-qualified name of the
	// AssetsService's GetChampionAbilityAssetUrl RPC.
	AssetsServiceGetChampionAbilityAssetUrlProcedure = "/assets.AssetsService/GetChampionAbilityAssetUrl"
	// AssetsServiceGetChampionPassiveAssetUrlProcedure is the fully-qualified name of the
	// AssetsService's GetChampionPassiveAssetUrl RPC.
	AssetsServiceGetChampionPassiveAssetUrlProcedure = "/assets.AssetsService/GetChampionPassiveAssetUrl"
	// AssetsServiceGetChampionSquareAssetUrlProcedure is the fully-qualified name of the
	// AssetsService's GetChampionSquareAssetUrl RPC.
	AssetsServiceGetChampionSquareAssetUrlProcedure = "/assets.AssetsService/GetChampionSquareAssetUrl"
	// AssetsServiceGetChampionLoadingScreenAssetUrlProcedure is the fully-qualified name of the
	// AssetsService's GetChampionLoadingScreenAssetUrl RPC.
	AssetsServiceGetChampionLoadingScreenAssetUrlProcedure = "/assets.AssetsService/GetChampionLoadingScreenAssetUrl"
	// AssetsServiceGetChampionSplashAssetUrlProcedure is the fully-qualified name of the
	// AssetsService's GetChampionSplashAssetUrl RPC.
	AssetsServiceGetChampionSplashAssetUrlProcedure = "/assets.AssetsService/GetChampionSplashAssetUrl"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	assetsServiceServiceDescriptor                                = v1.File_assets_v1_assets_proto.Services().ByName("AssetsService")
	assetsServiceGetRuneIconMethodDescriptor                      = assetsServiceServiceDescriptor.Methods().ByName("GetRuneIcon")
	assetsServiceGetSummonerSpellIconMethodDescriptor             = assetsServiceServiceDescriptor.Methods().ByName("GetSummonerSpellIcon")
	assetsServiceGetItemAssetUrlMethodDescriptor                  = assetsServiceServiceDescriptor.Methods().ByName("GetItemAssetUrl")
	assetsServiceGetSpellAssetUrlMethodDescriptor                 = assetsServiceServiceDescriptor.Methods().ByName("GetSpellAssetUrl")
	assetsServiceGetChampionAbilityAssetUrlMethodDescriptor       = assetsServiceServiceDescriptor.Methods().ByName("GetChampionAbilityAssetUrl")
	assetsServiceGetChampionPassiveAssetUrlMethodDescriptor       = assetsServiceServiceDescriptor.Methods().ByName("GetChampionPassiveAssetUrl")
	assetsServiceGetChampionSquareAssetUrlMethodDescriptor        = assetsServiceServiceDescriptor.Methods().ByName("GetChampionSquareAssetUrl")
	assetsServiceGetChampionLoadingScreenAssetUrlMethodDescriptor = assetsServiceServiceDescriptor.Methods().ByName("GetChampionLoadingScreenAssetUrl")
	assetsServiceGetChampionSplashAssetUrlMethodDescriptor        = assetsServiceServiceDescriptor.Methods().ByName("GetChampionSplashAssetUrl")
)

// AssetsServiceClient is a client for the assets.AssetsService service.
type AssetsServiceClient interface {
	GetRuneIcon(context.Context, *connect.Request[v1.GetRuneIconRequest]) (*connect.Response[v1.GetRuneIconResponse], error)
	GetSummonerSpellIcon(context.Context, *connect.Request[v1.GetSummonerSpellIconRequest]) (*connect.Response[v1.GetSummonerSpellIconResponse], error)
	GetItemAssetUrl(context.Context, *connect.Request[v1.GetItemAssetUrlRequest]) (*connect.Response[v1.GetItemAssetUrlResponse], error)
	GetSpellAssetUrl(context.Context, *connect.Request[v1.GetSpellAssetUrlRequest]) (*connect.Response[v1.GetSpellAssetUrlResponse], error)
	GetChampionAbilityAssetUrl(context.Context, *connect.Request[v1.GetChampionAbilityAssetUrlRequest]) (*connect.Response[v1.GetChampionAbilityAssetUrlResponse], error)
	GetChampionPassiveAssetUrl(context.Context, *connect.Request[v1.GetChampionPassiveAssetUrlRequest]) (*connect.Response[v1.GetChampionPassiveAssetUrlResponse], error)
	GetChampionSquareAssetUrl(context.Context, *connect.Request[v1.GetChampionSquareAssetUrlRequest]) (*connect.Response[v1.GetChampionSquareAssetUrlResponse], error)
	GetChampionLoadingScreenAssetUrl(context.Context, *connect.Request[v1.GetChampionLoadingScreenAssetUrlRequest]) (*connect.Response[v1.GetChampionLoadingScreenAssetUrlResponse], error)
	GetChampionSplashAssetUrl(context.Context, *connect.Request[v1.GetChampionSplashAssetUrlRequest]) (*connect.Response[v1.GetChampionSplashAssetUrlResponse], error)
}

// NewAssetsServiceClient constructs a client for the assets.AssetsService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAssetsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AssetsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &assetsServiceClient{
		getRuneIcon: connect.NewClient[v1.GetRuneIconRequest, v1.GetRuneIconResponse](
			httpClient,
			baseURL+AssetsServiceGetRuneIconProcedure,
			connect.WithSchema(assetsServiceGetRuneIconMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getSummonerSpellIcon: connect.NewClient[v1.GetSummonerSpellIconRequest, v1.GetSummonerSpellIconResponse](
			httpClient,
			baseURL+AssetsServiceGetSummonerSpellIconProcedure,
			connect.WithSchema(assetsServiceGetSummonerSpellIconMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getItemAssetUrl: connect.NewClient[v1.GetItemAssetUrlRequest, v1.GetItemAssetUrlResponse](
			httpClient,
			baseURL+AssetsServiceGetItemAssetUrlProcedure,
			connect.WithSchema(assetsServiceGetItemAssetUrlMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getSpellAssetUrl: connect.NewClient[v1.GetSpellAssetUrlRequest, v1.GetSpellAssetUrlResponse](
			httpClient,
			baseURL+AssetsServiceGetSpellAssetUrlProcedure,
			connect.WithSchema(assetsServiceGetSpellAssetUrlMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getChampionAbilityAssetUrl: connect.NewClient[v1.GetChampionAbilityAssetUrlRequest, v1.GetChampionAbilityAssetUrlResponse](
			httpClient,
			baseURL+AssetsServiceGetChampionAbilityAssetUrlProcedure,
			connect.WithSchema(assetsServiceGetChampionAbilityAssetUrlMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getChampionPassiveAssetUrl: connect.NewClient[v1.GetChampionPassiveAssetUrlRequest, v1.GetChampionPassiveAssetUrlResponse](
			httpClient,
			baseURL+AssetsServiceGetChampionPassiveAssetUrlProcedure,
			connect.WithSchema(assetsServiceGetChampionPassiveAssetUrlMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getChampionSquareAssetUrl: connect.NewClient[v1.GetChampionSquareAssetUrlRequest, v1.GetChampionSquareAssetUrlResponse](
			httpClient,
			baseURL+AssetsServiceGetChampionSquareAssetUrlProcedure,
			connect.WithSchema(assetsServiceGetChampionSquareAssetUrlMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getChampionLoadingScreenAssetUrl: connect.NewClient[v1.GetChampionLoadingScreenAssetUrlRequest, v1.GetChampionLoadingScreenAssetUrlResponse](
			httpClient,
			baseURL+AssetsServiceGetChampionLoadingScreenAssetUrlProcedure,
			connect.WithSchema(assetsServiceGetChampionLoadingScreenAssetUrlMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getChampionSplashAssetUrl: connect.NewClient[v1.GetChampionSplashAssetUrlRequest, v1.GetChampionSplashAssetUrlResponse](
			httpClient,
			baseURL+AssetsServiceGetChampionSplashAssetUrlProcedure,
			connect.WithSchema(assetsServiceGetChampionSplashAssetUrlMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// assetsServiceClient implements AssetsServiceClient.
type assetsServiceClient struct {
	getRuneIcon                      *connect.Client[v1.GetRuneIconRequest, v1.GetRuneIconResponse]
	getSummonerSpellIcon             *connect.Client[v1.GetSummonerSpellIconRequest, v1.GetSummonerSpellIconResponse]
	getItemAssetUrl                  *connect.Client[v1.GetItemAssetUrlRequest, v1.GetItemAssetUrlResponse]
	getSpellAssetUrl                 *connect.Client[v1.GetSpellAssetUrlRequest, v1.GetSpellAssetUrlResponse]
	getChampionAbilityAssetUrl       *connect.Client[v1.GetChampionAbilityAssetUrlRequest, v1.GetChampionAbilityAssetUrlResponse]
	getChampionPassiveAssetUrl       *connect.Client[v1.GetChampionPassiveAssetUrlRequest, v1.GetChampionPassiveAssetUrlResponse]
	getChampionSquareAssetUrl        *connect.Client[v1.GetChampionSquareAssetUrlRequest, v1.GetChampionSquareAssetUrlResponse]
	getChampionLoadingScreenAssetUrl *connect.Client[v1.GetChampionLoadingScreenAssetUrlRequest, v1.GetChampionLoadingScreenAssetUrlResponse]
	getChampionSplashAssetUrl        *connect.Client[v1.GetChampionSplashAssetUrlRequest, v1.GetChampionSplashAssetUrlResponse]
}

// GetRuneIcon calls assets.AssetsService.GetRuneIcon.
func (c *assetsServiceClient) GetRuneIcon(ctx context.Context, req *connect.Request[v1.GetRuneIconRequest]) (*connect.Response[v1.GetRuneIconResponse], error) {
	return c.getRuneIcon.CallUnary(ctx, req)
}

// GetSummonerSpellIcon calls assets.AssetsService.GetSummonerSpellIcon.
func (c *assetsServiceClient) GetSummonerSpellIcon(ctx context.Context, req *connect.Request[v1.GetSummonerSpellIconRequest]) (*connect.Response[v1.GetSummonerSpellIconResponse], error) {
	return c.getSummonerSpellIcon.CallUnary(ctx, req)
}

// GetItemAssetUrl calls assets.AssetsService.GetItemAssetUrl.
func (c *assetsServiceClient) GetItemAssetUrl(ctx context.Context, req *connect.Request[v1.GetItemAssetUrlRequest]) (*connect.Response[v1.GetItemAssetUrlResponse], error) {
	return c.getItemAssetUrl.CallUnary(ctx, req)
}

// GetSpellAssetUrl calls assets.AssetsService.GetSpellAssetUrl.
func (c *assetsServiceClient) GetSpellAssetUrl(ctx context.Context, req *connect.Request[v1.GetSpellAssetUrlRequest]) (*connect.Response[v1.GetSpellAssetUrlResponse], error) {
	return c.getSpellAssetUrl.CallUnary(ctx, req)
}

// GetChampionAbilityAssetUrl calls assets.AssetsService.GetChampionAbilityAssetUrl.
func (c *assetsServiceClient) GetChampionAbilityAssetUrl(ctx context.Context, req *connect.Request[v1.GetChampionAbilityAssetUrlRequest]) (*connect.Response[v1.GetChampionAbilityAssetUrlResponse], error) {
	return c.getChampionAbilityAssetUrl.CallUnary(ctx, req)
}

// GetChampionPassiveAssetUrl calls assets.AssetsService.GetChampionPassiveAssetUrl.
func (c *assetsServiceClient) GetChampionPassiveAssetUrl(ctx context.Context, req *connect.Request[v1.GetChampionPassiveAssetUrlRequest]) (*connect.Response[v1.GetChampionPassiveAssetUrlResponse], error) {
	return c.getChampionPassiveAssetUrl.CallUnary(ctx, req)
}

// GetChampionSquareAssetUrl calls assets.AssetsService.GetChampionSquareAssetUrl.
func (c *assetsServiceClient) GetChampionSquareAssetUrl(ctx context.Context, req *connect.Request[v1.GetChampionSquareAssetUrlRequest]) (*connect.Response[v1.GetChampionSquareAssetUrlResponse], error) {
	return c.getChampionSquareAssetUrl.CallUnary(ctx, req)
}

// GetChampionLoadingScreenAssetUrl calls assets.AssetsService.GetChampionLoadingScreenAssetUrl.
func (c *assetsServiceClient) GetChampionLoadingScreenAssetUrl(ctx context.Context, req *connect.Request[v1.GetChampionLoadingScreenAssetUrlRequest]) (*connect.Response[v1.GetChampionLoadingScreenAssetUrlResponse], error) {
	return c.getChampionLoadingScreenAssetUrl.CallUnary(ctx, req)
}

// GetChampionSplashAssetUrl calls assets.AssetsService.GetChampionSplashAssetUrl.
func (c *assetsServiceClient) GetChampionSplashAssetUrl(ctx context.Context, req *connect.Request[v1.GetChampionSplashAssetUrlRequest]) (*connect.Response[v1.GetChampionSplashAssetUrlResponse], error) {
	return c.getChampionSplashAssetUrl.CallUnary(ctx, req)
}

// AssetsServiceHandler is an implementation of the assets.AssetsService service.
type AssetsServiceHandler interface {
	GetRuneIcon(context.Context, *connect.Request[v1.GetRuneIconRequest]) (*connect.Response[v1.GetRuneIconResponse], error)
	GetSummonerSpellIcon(context.Context, *connect.Request[v1.GetSummonerSpellIconRequest]) (*connect.Response[v1.GetSummonerSpellIconResponse], error)
	GetItemAssetUrl(context.Context, *connect.Request[v1.GetItemAssetUrlRequest]) (*connect.Response[v1.GetItemAssetUrlResponse], error)
	GetSpellAssetUrl(context.Context, *connect.Request[v1.GetSpellAssetUrlRequest]) (*connect.Response[v1.GetSpellAssetUrlResponse], error)
	GetChampionAbilityAssetUrl(context.Context, *connect.Request[v1.GetChampionAbilityAssetUrlRequest]) (*connect.Response[v1.GetChampionAbilityAssetUrlResponse], error)
	GetChampionPassiveAssetUrl(context.Context, *connect.Request[v1.GetChampionPassiveAssetUrlRequest]) (*connect.Response[v1.GetChampionPassiveAssetUrlResponse], error)
	GetChampionSquareAssetUrl(context.Context, *connect.Request[v1.GetChampionSquareAssetUrlRequest]) (*connect.Response[v1.GetChampionSquareAssetUrlResponse], error)
	GetChampionLoadingScreenAssetUrl(context.Context, *connect.Request[v1.GetChampionLoadingScreenAssetUrlRequest]) (*connect.Response[v1.GetChampionLoadingScreenAssetUrlResponse], error)
	GetChampionSplashAssetUrl(context.Context, *connect.Request[v1.GetChampionSplashAssetUrlRequest]) (*connect.Response[v1.GetChampionSplashAssetUrlResponse], error)
}

// NewAssetsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAssetsServiceHandler(svc AssetsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	assetsServiceGetRuneIconHandler := connect.NewUnaryHandler(
		AssetsServiceGetRuneIconProcedure,
		svc.GetRuneIcon,
		connect.WithSchema(assetsServiceGetRuneIconMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	assetsServiceGetSummonerSpellIconHandler := connect.NewUnaryHandler(
		AssetsServiceGetSummonerSpellIconProcedure,
		svc.GetSummonerSpellIcon,
		connect.WithSchema(assetsServiceGetSummonerSpellIconMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	assetsServiceGetItemAssetUrlHandler := connect.NewUnaryHandler(
		AssetsServiceGetItemAssetUrlProcedure,
		svc.GetItemAssetUrl,
		connect.WithSchema(assetsServiceGetItemAssetUrlMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	assetsServiceGetSpellAssetUrlHandler := connect.NewUnaryHandler(
		AssetsServiceGetSpellAssetUrlProcedure,
		svc.GetSpellAssetUrl,
		connect.WithSchema(assetsServiceGetSpellAssetUrlMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	assetsServiceGetChampionAbilityAssetUrlHandler := connect.NewUnaryHandler(
		AssetsServiceGetChampionAbilityAssetUrlProcedure,
		svc.GetChampionAbilityAssetUrl,
		connect.WithSchema(assetsServiceGetChampionAbilityAssetUrlMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	assetsServiceGetChampionPassiveAssetUrlHandler := connect.NewUnaryHandler(
		AssetsServiceGetChampionPassiveAssetUrlProcedure,
		svc.GetChampionPassiveAssetUrl,
		connect.WithSchema(assetsServiceGetChampionPassiveAssetUrlMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	assetsServiceGetChampionSquareAssetUrlHandler := connect.NewUnaryHandler(
		AssetsServiceGetChampionSquareAssetUrlProcedure,
		svc.GetChampionSquareAssetUrl,
		connect.WithSchema(assetsServiceGetChampionSquareAssetUrlMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	assetsServiceGetChampionLoadingScreenAssetUrlHandler := connect.NewUnaryHandler(
		AssetsServiceGetChampionLoadingScreenAssetUrlProcedure,
		svc.GetChampionLoadingScreenAssetUrl,
		connect.WithSchema(assetsServiceGetChampionLoadingScreenAssetUrlMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	assetsServiceGetChampionSplashAssetUrlHandler := connect.NewUnaryHandler(
		AssetsServiceGetChampionSplashAssetUrlProcedure,
		svc.GetChampionSplashAssetUrl,
		connect.WithSchema(assetsServiceGetChampionSplashAssetUrlMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/assets.AssetsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AssetsServiceGetRuneIconProcedure:
			assetsServiceGetRuneIconHandler.ServeHTTP(w, r)
		case AssetsServiceGetSummonerSpellIconProcedure:
			assetsServiceGetSummonerSpellIconHandler.ServeHTTP(w, r)
		case AssetsServiceGetItemAssetUrlProcedure:
			assetsServiceGetItemAssetUrlHandler.ServeHTTP(w, r)
		case AssetsServiceGetSpellAssetUrlProcedure:
			assetsServiceGetSpellAssetUrlHandler.ServeHTTP(w, r)
		case AssetsServiceGetChampionAbilityAssetUrlProcedure:
			assetsServiceGetChampionAbilityAssetUrlHandler.ServeHTTP(w, r)
		case AssetsServiceGetChampionPassiveAssetUrlProcedure:
			assetsServiceGetChampionPassiveAssetUrlHandler.ServeHTTP(w, r)
		case AssetsServiceGetChampionSquareAssetUrlProcedure:
			assetsServiceGetChampionSquareAssetUrlHandler.ServeHTTP(w, r)
		case AssetsServiceGetChampionLoadingScreenAssetUrlProcedure:
			assetsServiceGetChampionLoadingScreenAssetUrlHandler.ServeHTTP(w, r)
		case AssetsServiceGetChampionSplashAssetUrlProcedure:
			assetsServiceGetChampionSplashAssetUrlHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAssetsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAssetsServiceHandler struct{}

func (UnimplementedAssetsServiceHandler) GetRuneIcon(context.Context, *connect.Request[v1.GetRuneIconRequest]) (*connect.Response[v1.GetRuneIconResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("assets.AssetsService.GetRuneIcon is not implemented"))
}

func (UnimplementedAssetsServiceHandler) GetSummonerSpellIcon(context.Context, *connect.Request[v1.GetSummonerSpellIconRequest]) (*connect.Response[v1.GetSummonerSpellIconResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("assets.AssetsService.GetSummonerSpellIcon is not implemented"))
}

func (UnimplementedAssetsServiceHandler) GetItemAssetUrl(context.Context, *connect.Request[v1.GetItemAssetUrlRequest]) (*connect.Response[v1.GetItemAssetUrlResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("assets.AssetsService.GetItemAssetUrl is not implemented"))
}

func (UnimplementedAssetsServiceHandler) GetSpellAssetUrl(context.Context, *connect.Request[v1.GetSpellAssetUrlRequest]) (*connect.Response[v1.GetSpellAssetUrlResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("assets.AssetsService.GetSpellAssetUrl is not implemented"))
}

func (UnimplementedAssetsServiceHandler) GetChampionAbilityAssetUrl(context.Context, *connect.Request[v1.GetChampionAbilityAssetUrlRequest]) (*connect.Response[v1.GetChampionAbilityAssetUrlResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("assets.AssetsService.GetChampionAbilityAssetUrl is not implemented"))
}

func (UnimplementedAssetsServiceHandler) GetChampionPassiveAssetUrl(context.Context, *connect.Request[v1.GetChampionPassiveAssetUrlRequest]) (*connect.Response[v1.GetChampionPassiveAssetUrlResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("assets.AssetsService.GetChampionPassiveAssetUrl is not implemented"))
}

func (UnimplementedAssetsServiceHandler) GetChampionSquareAssetUrl(context.Context, *connect.Request[v1.GetChampionSquareAssetUrlRequest]) (*connect.Response[v1.GetChampionSquareAssetUrlResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("assets.AssetsService.GetChampionSquareAssetUrl is not implemented"))
}

func (UnimplementedAssetsServiceHandler) GetChampionLoadingScreenAssetUrl(context.Context, *connect.Request[v1.GetChampionLoadingScreenAssetUrlRequest]) (*connect.Response[v1.GetChampionLoadingScreenAssetUrlResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("assets.AssetsService.GetChampionLoadingScreenAssetUrl is not implemented"))
}

func (UnimplementedAssetsServiceHandler) GetChampionSplashAssetUrl(context.Context, *connect.Request[v1.GetChampionSplashAssetUrlRequest]) (*connect.Response[v1.GetChampionSplashAssetUrlResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("assets.AssetsService.GetChampionSplashAssetUrl is not implemented"))
}
