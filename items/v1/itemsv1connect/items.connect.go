// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: items/v1/items.proto

package itemsv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "maxischmaxi/jstreams/items/v1"
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
	// ItemsServiceName is the fully-qualified name of the ItemsService service.
	ItemsServiceName = "items.ItemsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ItemsServiceGetItemImageByIdProcedure is the fully-qualified name of the ItemsService's
	// GetItemImageById RPC.
	ItemsServiceGetItemImageByIdProcedure = "/items.ItemsService/GetItemImageById"
	// ItemsServiceGetItemInformationByIdProcedure is the fully-qualified name of the ItemsService's
	// GetItemInformationById RPC.
	ItemsServiceGetItemInformationByIdProcedure = "/items.ItemsService/GetItemInformationById"
	// ItemsServiceGetBaseItemsProcedure is the fully-qualified name of the ItemsService's GetBaseItems
	// RPC.
	ItemsServiceGetBaseItemsProcedure = "/items.ItemsService/GetBaseItems"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	itemsServiceServiceDescriptor                      = v1.File_items_v1_items_proto.Services().ByName("ItemsService")
	itemsServiceGetItemImageByIdMethodDescriptor       = itemsServiceServiceDescriptor.Methods().ByName("GetItemImageById")
	itemsServiceGetItemInformationByIdMethodDescriptor = itemsServiceServiceDescriptor.Methods().ByName("GetItemInformationById")
	itemsServiceGetBaseItemsMethodDescriptor           = itemsServiceServiceDescriptor.Methods().ByName("GetBaseItems")
)

// ItemsServiceClient is a client for the items.ItemsService service.
type ItemsServiceClient interface {
	GetItemImageById(context.Context, *connect.Request[v1.GetItemImageByIdRequest]) (*connect.Response[v1.GetItemImageByIdResponse], error)
	GetItemInformationById(context.Context, *connect.Request[v1.GetItemInformationByIdRequest]) (*connect.Response[v1.GetItemInformationByIdResponse], error)
	GetBaseItems(context.Context, *connect.Request[v1.GetBaseItemsRequest]) (*connect.Response[v1.GetBaseItemsResponse], error)
}

// NewItemsServiceClient constructs a client for the items.ItemsService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewItemsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ItemsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &itemsServiceClient{
		getItemImageById: connect.NewClient[v1.GetItemImageByIdRequest, v1.GetItemImageByIdResponse](
			httpClient,
			baseURL+ItemsServiceGetItemImageByIdProcedure,
			connect.WithSchema(itemsServiceGetItemImageByIdMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getItemInformationById: connect.NewClient[v1.GetItemInformationByIdRequest, v1.GetItemInformationByIdResponse](
			httpClient,
			baseURL+ItemsServiceGetItemInformationByIdProcedure,
			connect.WithSchema(itemsServiceGetItemInformationByIdMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getBaseItems: connect.NewClient[v1.GetBaseItemsRequest, v1.GetBaseItemsResponse](
			httpClient,
			baseURL+ItemsServiceGetBaseItemsProcedure,
			connect.WithSchema(itemsServiceGetBaseItemsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// itemsServiceClient implements ItemsServiceClient.
type itemsServiceClient struct {
	getItemImageById       *connect.Client[v1.GetItemImageByIdRequest, v1.GetItemImageByIdResponse]
	getItemInformationById *connect.Client[v1.GetItemInformationByIdRequest, v1.GetItemInformationByIdResponse]
	getBaseItems           *connect.Client[v1.GetBaseItemsRequest, v1.GetBaseItemsResponse]
}

// GetItemImageById calls items.ItemsService.GetItemImageById.
func (c *itemsServiceClient) GetItemImageById(ctx context.Context, req *connect.Request[v1.GetItemImageByIdRequest]) (*connect.Response[v1.GetItemImageByIdResponse], error) {
	return c.getItemImageById.CallUnary(ctx, req)
}

// GetItemInformationById calls items.ItemsService.GetItemInformationById.
func (c *itemsServiceClient) GetItemInformationById(ctx context.Context, req *connect.Request[v1.GetItemInformationByIdRequest]) (*connect.Response[v1.GetItemInformationByIdResponse], error) {
	return c.getItemInformationById.CallUnary(ctx, req)
}

// GetBaseItems calls items.ItemsService.GetBaseItems.
func (c *itemsServiceClient) GetBaseItems(ctx context.Context, req *connect.Request[v1.GetBaseItemsRequest]) (*connect.Response[v1.GetBaseItemsResponse], error) {
	return c.getBaseItems.CallUnary(ctx, req)
}

// ItemsServiceHandler is an implementation of the items.ItemsService service.
type ItemsServiceHandler interface {
	GetItemImageById(context.Context, *connect.Request[v1.GetItemImageByIdRequest]) (*connect.Response[v1.GetItemImageByIdResponse], error)
	GetItemInformationById(context.Context, *connect.Request[v1.GetItemInformationByIdRequest]) (*connect.Response[v1.GetItemInformationByIdResponse], error)
	GetBaseItems(context.Context, *connect.Request[v1.GetBaseItemsRequest]) (*connect.Response[v1.GetBaseItemsResponse], error)
}

// NewItemsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewItemsServiceHandler(svc ItemsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	itemsServiceGetItemImageByIdHandler := connect.NewUnaryHandler(
		ItemsServiceGetItemImageByIdProcedure,
		svc.GetItemImageById,
		connect.WithSchema(itemsServiceGetItemImageByIdMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	itemsServiceGetItemInformationByIdHandler := connect.NewUnaryHandler(
		ItemsServiceGetItemInformationByIdProcedure,
		svc.GetItemInformationById,
		connect.WithSchema(itemsServiceGetItemInformationByIdMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	itemsServiceGetBaseItemsHandler := connect.NewUnaryHandler(
		ItemsServiceGetBaseItemsProcedure,
		svc.GetBaseItems,
		connect.WithSchema(itemsServiceGetBaseItemsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/items.ItemsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ItemsServiceGetItemImageByIdProcedure:
			itemsServiceGetItemImageByIdHandler.ServeHTTP(w, r)
		case ItemsServiceGetItemInformationByIdProcedure:
			itemsServiceGetItemInformationByIdHandler.ServeHTTP(w, r)
		case ItemsServiceGetBaseItemsProcedure:
			itemsServiceGetBaseItemsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedItemsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedItemsServiceHandler struct{}

func (UnimplementedItemsServiceHandler) GetItemImageById(context.Context, *connect.Request[v1.GetItemImageByIdRequest]) (*connect.Response[v1.GetItemImageByIdResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("items.ItemsService.GetItemImageById is not implemented"))
}

func (UnimplementedItemsServiceHandler) GetItemInformationById(context.Context, *connect.Request[v1.GetItemInformationByIdRequest]) (*connect.Response[v1.GetItemInformationByIdResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("items.ItemsService.GetItemInformationById is not implemented"))
}

func (UnimplementedItemsServiceHandler) GetBaseItems(context.Context, *connect.Request[v1.GetBaseItemsRequest]) (*connect.Response[v1.GetBaseItemsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("items.ItemsService.GetBaseItems is not implemented"))
}
