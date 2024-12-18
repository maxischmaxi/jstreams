// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: account/v1/account.proto

package accountv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "maxischmaxi/jstreams/account/v1"
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
	// AccountServiceName is the fully-qualified name of the AccountService service.
	AccountServiceName = "account.AccountService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AccountServiceGetAccountByGamenameAndTaglineProcedure is the fully-qualified name of the
	// AccountService's GetAccountByGamenameAndTagline RPC.
	AccountServiceGetAccountByGamenameAndTaglineProcedure = "/account.AccountService/GetAccountByGamenameAndTagline"
	// AccountServiceGetAccountProfileIconProcedure is the fully-qualified name of the AccountService's
	// GetAccountProfileIcon RPC.
	AccountServiceGetAccountProfileIconProcedure = "/account.AccountService/GetAccountProfileIcon"
	// AccountServiceGetAccountByPuuidProcedure is the fully-qualified name of the AccountService's
	// GetAccountByPuuid RPC.
	AccountServiceGetAccountByPuuidProcedure = "/account.AccountService/GetAccountByPuuid"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	accountServiceServiceDescriptor                              = v1.File_account_v1_account_proto.Services().ByName("AccountService")
	accountServiceGetAccountByGamenameAndTaglineMethodDescriptor = accountServiceServiceDescriptor.Methods().ByName("GetAccountByGamenameAndTagline")
	accountServiceGetAccountProfileIconMethodDescriptor          = accountServiceServiceDescriptor.Methods().ByName("GetAccountProfileIcon")
	accountServiceGetAccountByPuuidMethodDescriptor              = accountServiceServiceDescriptor.Methods().ByName("GetAccountByPuuid")
)

// AccountServiceClient is a client for the account.AccountService service.
type AccountServiceClient interface {
	GetAccountByGamenameAndTagline(context.Context, *connect.Request[v1.GetAccountByGamenameAndTaglineRequest]) (*connect.Response[v1.GetAccountByGamenameAndTaglineResponse], error)
	GetAccountProfileIcon(context.Context, *connect.Request[v1.GetAccountProfileIconRequest]) (*connect.Response[v1.GetAccountProfileIconResponse], error)
	GetAccountByPuuid(context.Context, *connect.Request[v1.GetAccountByPuuidRequest]) (*connect.Response[v1.GetAccountByPuuidResponse], error)
}

// NewAccountServiceClient constructs a client for the account.AccountService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAccountServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AccountServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &accountServiceClient{
		getAccountByGamenameAndTagline: connect.NewClient[v1.GetAccountByGamenameAndTaglineRequest, v1.GetAccountByGamenameAndTaglineResponse](
			httpClient,
			baseURL+AccountServiceGetAccountByGamenameAndTaglineProcedure,
			connect.WithSchema(accountServiceGetAccountByGamenameAndTaglineMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getAccountProfileIcon: connect.NewClient[v1.GetAccountProfileIconRequest, v1.GetAccountProfileIconResponse](
			httpClient,
			baseURL+AccountServiceGetAccountProfileIconProcedure,
			connect.WithSchema(accountServiceGetAccountProfileIconMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getAccountByPuuid: connect.NewClient[v1.GetAccountByPuuidRequest, v1.GetAccountByPuuidResponse](
			httpClient,
			baseURL+AccountServiceGetAccountByPuuidProcedure,
			connect.WithSchema(accountServiceGetAccountByPuuidMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// accountServiceClient implements AccountServiceClient.
type accountServiceClient struct {
	getAccountByGamenameAndTagline *connect.Client[v1.GetAccountByGamenameAndTaglineRequest, v1.GetAccountByGamenameAndTaglineResponse]
	getAccountProfileIcon          *connect.Client[v1.GetAccountProfileIconRequest, v1.GetAccountProfileIconResponse]
	getAccountByPuuid              *connect.Client[v1.GetAccountByPuuidRequest, v1.GetAccountByPuuidResponse]
}

// GetAccountByGamenameAndTagline calls account.AccountService.GetAccountByGamenameAndTagline.
func (c *accountServiceClient) GetAccountByGamenameAndTagline(ctx context.Context, req *connect.Request[v1.GetAccountByGamenameAndTaglineRequest]) (*connect.Response[v1.GetAccountByGamenameAndTaglineResponse], error) {
	return c.getAccountByGamenameAndTagline.CallUnary(ctx, req)
}

// GetAccountProfileIcon calls account.AccountService.GetAccountProfileIcon.
func (c *accountServiceClient) GetAccountProfileIcon(ctx context.Context, req *connect.Request[v1.GetAccountProfileIconRequest]) (*connect.Response[v1.GetAccountProfileIconResponse], error) {
	return c.getAccountProfileIcon.CallUnary(ctx, req)
}

// GetAccountByPuuid calls account.AccountService.GetAccountByPuuid.
func (c *accountServiceClient) GetAccountByPuuid(ctx context.Context, req *connect.Request[v1.GetAccountByPuuidRequest]) (*connect.Response[v1.GetAccountByPuuidResponse], error) {
	return c.getAccountByPuuid.CallUnary(ctx, req)
}

// AccountServiceHandler is an implementation of the account.AccountService service.
type AccountServiceHandler interface {
	GetAccountByGamenameAndTagline(context.Context, *connect.Request[v1.GetAccountByGamenameAndTaglineRequest]) (*connect.Response[v1.GetAccountByGamenameAndTaglineResponse], error)
	GetAccountProfileIcon(context.Context, *connect.Request[v1.GetAccountProfileIconRequest]) (*connect.Response[v1.GetAccountProfileIconResponse], error)
	GetAccountByPuuid(context.Context, *connect.Request[v1.GetAccountByPuuidRequest]) (*connect.Response[v1.GetAccountByPuuidResponse], error)
}

// NewAccountServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAccountServiceHandler(svc AccountServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	accountServiceGetAccountByGamenameAndTaglineHandler := connect.NewUnaryHandler(
		AccountServiceGetAccountByGamenameAndTaglineProcedure,
		svc.GetAccountByGamenameAndTagline,
		connect.WithSchema(accountServiceGetAccountByGamenameAndTaglineMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServiceGetAccountProfileIconHandler := connect.NewUnaryHandler(
		AccountServiceGetAccountProfileIconProcedure,
		svc.GetAccountProfileIcon,
		connect.WithSchema(accountServiceGetAccountProfileIconMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServiceGetAccountByPuuidHandler := connect.NewUnaryHandler(
		AccountServiceGetAccountByPuuidProcedure,
		svc.GetAccountByPuuid,
		connect.WithSchema(accountServiceGetAccountByPuuidMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/account.AccountService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AccountServiceGetAccountByGamenameAndTaglineProcedure:
			accountServiceGetAccountByGamenameAndTaglineHandler.ServeHTTP(w, r)
		case AccountServiceGetAccountProfileIconProcedure:
			accountServiceGetAccountProfileIconHandler.ServeHTTP(w, r)
		case AccountServiceGetAccountByPuuidProcedure:
			accountServiceGetAccountByPuuidHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAccountServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAccountServiceHandler struct{}

func (UnimplementedAccountServiceHandler) GetAccountByGamenameAndTagline(context.Context, *connect.Request[v1.GetAccountByGamenameAndTaglineRequest]) (*connect.Response[v1.GetAccountByGamenameAndTaglineResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("account.AccountService.GetAccountByGamenameAndTagline is not implemented"))
}

func (UnimplementedAccountServiceHandler) GetAccountProfileIcon(context.Context, *connect.Request[v1.GetAccountProfileIconRequest]) (*connect.Response[v1.GetAccountProfileIconResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("account.AccountService.GetAccountProfileIcon is not implemented"))
}

func (UnimplementedAccountServiceHandler) GetAccountByPuuid(context.Context, *connect.Request[v1.GetAccountByPuuidRequest]) (*connect.Response[v1.GetAccountByPuuidResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("account.AccountService.GetAccountByPuuid is not implemented"))
}
