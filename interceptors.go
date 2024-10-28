package main

import (
	"context"

	"connectrpc.com/connect"
)

func LoggingInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			// log.Printf("[%s] %s %s %s", req.HTTPMethod(), req.Peer().Protocol, req.Peer().Addr, req.Peer().Query)
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
