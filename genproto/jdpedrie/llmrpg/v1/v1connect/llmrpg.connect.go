// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: jdpedrie/llmrpg/v1/llmrpg.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/jdpedrie/llmrpg/genproto/jdpedrie/llmrpg/v1"
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
	// LLMRPGServiceName is the fully-qualified name of the LLMRPGService service.
	LLMRPGServiceName = "v1.LLMRPGService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// LLMRPGServicePlayProcedure is the fully-qualified name of the LLMRPGService's Play RPC.
	LLMRPGServicePlayProcedure = "/v1.LLMRPGService/Play"
)

// LLMRPGServiceClient is a client for the v1.LLMRPGService service.
type LLMRPGServiceClient interface {
	Play(context.Context) *connect.BidiStreamForClient[v1.PlayRequest, v1.PlayResponse]
}

// NewLLMRPGServiceClient constructs a client for the v1.LLMRPGService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewLLMRPGServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) LLMRPGServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	lLMRPGServiceMethods := v1.File_jdpedrie_llmrpg_v1_llmrpg_proto.Services().ByName("LLMRPGService").Methods()
	return &lLMRPGServiceClient{
		play: connect.NewClient[v1.PlayRequest, v1.PlayResponse](
			httpClient,
			baseURL+LLMRPGServicePlayProcedure,
			connect.WithSchema(lLMRPGServiceMethods.ByName("Play")),
			connect.WithClientOptions(opts...),
		),
	}
}

// lLMRPGServiceClient implements LLMRPGServiceClient.
type lLMRPGServiceClient struct {
	play *connect.Client[v1.PlayRequest, v1.PlayResponse]
}

// Play calls v1.LLMRPGService.Play.
func (c *lLMRPGServiceClient) Play(ctx context.Context) *connect.BidiStreamForClient[v1.PlayRequest, v1.PlayResponse] {
	return c.play.CallBidiStream(ctx)
}

// LLMRPGServiceHandler is an implementation of the v1.LLMRPGService service.
type LLMRPGServiceHandler interface {
	Play(context.Context, *connect.BidiStream[v1.PlayRequest, v1.PlayResponse]) error
}

// NewLLMRPGServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewLLMRPGServiceHandler(svc LLMRPGServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	lLMRPGServiceMethods := v1.File_jdpedrie_llmrpg_v1_llmrpg_proto.Services().ByName("LLMRPGService").Methods()
	lLMRPGServicePlayHandler := connect.NewBidiStreamHandler(
		LLMRPGServicePlayProcedure,
		svc.Play,
		connect.WithSchema(lLMRPGServiceMethods.ByName("Play")),
		connect.WithHandlerOptions(opts...),
	)
	return "/v1.LLMRPGService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case LLMRPGServicePlayProcedure:
			lLMRPGServicePlayHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedLLMRPGServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedLLMRPGServiceHandler struct{}

func (UnimplementedLLMRPGServiceHandler) Play(context.Context, *connect.BidiStream[v1.PlayRequest, v1.PlayResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("v1.LLMRPGService.Play is not implemented"))
}
