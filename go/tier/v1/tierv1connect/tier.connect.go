// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tier/v1/tier.proto

package tierv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "maxischmaxi/jstreams/tier/v1"
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
	// TierServiceName is the fully-qualified name of the TierService service.
	TierServiceName = "tier.TierService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TierServiceGetTierIconProcedure is the fully-qualified name of the TierService's GetTierIcon RPC.
	TierServiceGetTierIconProcedure = "/tier.TierService/GetTierIcon"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	tierServiceServiceDescriptor           = v1.File_tier_v1_tier_proto.Services().ByName("TierService")
	tierServiceGetTierIconMethodDescriptor = tierServiceServiceDescriptor.Methods().ByName("GetTierIcon")
)

// TierServiceClient is a client for the tier.TierService service.
type TierServiceClient interface {
	GetTierIcon(context.Context, *connect.Request[v1.GetTierIconRequest]) (*connect.Response[v1.GetTierIconResponse], error)
}

// NewTierServiceClient constructs a client for the tier.TierService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTierServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TierServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &tierServiceClient{
		getTierIcon: connect.NewClient[v1.GetTierIconRequest, v1.GetTierIconResponse](
			httpClient,
			baseURL+TierServiceGetTierIconProcedure,
			connect.WithSchema(tierServiceGetTierIconMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// tierServiceClient implements TierServiceClient.
type tierServiceClient struct {
	getTierIcon *connect.Client[v1.GetTierIconRequest, v1.GetTierIconResponse]
}

// GetTierIcon calls tier.TierService.GetTierIcon.
func (c *tierServiceClient) GetTierIcon(ctx context.Context, req *connect.Request[v1.GetTierIconRequest]) (*connect.Response[v1.GetTierIconResponse], error) {
	return c.getTierIcon.CallUnary(ctx, req)
}

// TierServiceHandler is an implementation of the tier.TierService service.
type TierServiceHandler interface {
	GetTierIcon(context.Context, *connect.Request[v1.GetTierIconRequest]) (*connect.Response[v1.GetTierIconResponse], error)
}

// NewTierServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTierServiceHandler(svc TierServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	tierServiceGetTierIconHandler := connect.NewUnaryHandler(
		TierServiceGetTierIconProcedure,
		svc.GetTierIcon,
		connect.WithSchema(tierServiceGetTierIconMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/tier.TierService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TierServiceGetTierIconProcedure:
			tierServiceGetTierIconHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTierServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTierServiceHandler struct{}

func (UnimplementedTierServiceHandler) GetTierIcon(context.Context, *connect.Request[v1.GetTierIconRequest]) (*connect.Response[v1.GetTierIconResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("tier.TierService.GetTierIcon is not implemented"))
}
