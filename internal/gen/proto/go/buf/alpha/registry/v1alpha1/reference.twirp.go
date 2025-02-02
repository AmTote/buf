// Copyright 2020-2021 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-twirp v8.0.0, DO NOT EDIT.
// source: buf/alpha/registry/v1alpha1/reference.proto

package registryv1alpha1

import context "context"
import fmt "fmt"
import http "net/http"
import ioutil "io/ioutil"
import json "encoding/json"
import strconv "strconv"
import strings "strings"

import protojson "google.golang.org/protobuf/encoding/protojson"
import proto "google.golang.org/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the twirp package used in your project.
// A compilation error at this line likely means your copy of the
// twirp package needs to be updated.
const _ = twirp.TwirpPackageIsVersion7

// ==========================
// ReferenceService Interface
// ==========================

// ReferenceService is a service that provides RPCs that allow the BSR to query
// for reference information.
type ReferenceService interface {
	// GetReferenceByName takes a reference name and returns the
	// reference either as a tag, branch, or commit.
	GetReferenceByName(context.Context, *GetReferenceByNameRequest) (*GetReferenceByNameResponse, error)
}

// ================================
// ReferenceService Protobuf Client
// ================================

type referenceServiceProtobufClient struct {
	client      HTTPClient
	urls        [1]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewReferenceServiceProtobufClient creates a Protobuf client that implements the ReferenceService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewReferenceServiceProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) ReferenceService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(clientOpts.PathPrefix(), "buf.alpha.registry.v1alpha1", "ReferenceService")
	urls := [1]string{
		serviceURL + "GetReferenceByName",
	}

	return &referenceServiceProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *referenceServiceProtobufClient) GetReferenceByName(ctx context.Context, in *GetReferenceByNameRequest) (*GetReferenceByNameResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "buf.alpha.registry.v1alpha1")
	ctx = ctxsetters.WithServiceName(ctx, "ReferenceService")
	ctx = ctxsetters.WithMethodName(ctx, "GetReferenceByName")
	caller := c.callGetReferenceByName
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *GetReferenceByNameRequest) (*GetReferenceByNameResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*GetReferenceByNameRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*GetReferenceByNameRequest) when calling interceptor")
					}
					return c.callGetReferenceByName(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*GetReferenceByNameResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*GetReferenceByNameResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *referenceServiceProtobufClient) callGetReferenceByName(ctx context.Context, in *GetReferenceByNameRequest) (*GetReferenceByNameResponse, error) {
	out := new(GetReferenceByNameResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ============================
// ReferenceService JSON Client
// ============================

type referenceServiceJSONClient struct {
	client      HTTPClient
	urls        [1]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewReferenceServiceJSONClient creates a JSON client that implements the ReferenceService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewReferenceServiceJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) ReferenceService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(clientOpts.PathPrefix(), "buf.alpha.registry.v1alpha1", "ReferenceService")
	urls := [1]string{
		serviceURL + "GetReferenceByName",
	}

	return &referenceServiceJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *referenceServiceJSONClient) GetReferenceByName(ctx context.Context, in *GetReferenceByNameRequest) (*GetReferenceByNameResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "buf.alpha.registry.v1alpha1")
	ctx = ctxsetters.WithServiceName(ctx, "ReferenceService")
	ctx = ctxsetters.WithMethodName(ctx, "GetReferenceByName")
	caller := c.callGetReferenceByName
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *GetReferenceByNameRequest) (*GetReferenceByNameResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*GetReferenceByNameRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*GetReferenceByNameRequest) when calling interceptor")
					}
					return c.callGetReferenceByName(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*GetReferenceByNameResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*GetReferenceByNameResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *referenceServiceJSONClient) callGetReferenceByName(ctx context.Context, in *GetReferenceByNameRequest) (*GetReferenceByNameResponse, error) {
	out := new(GetReferenceByNameResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ===============================
// ReferenceService Server Handler
// ===============================

type referenceServiceServer struct {
	ReferenceService
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
}

// NewReferenceServiceServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewReferenceServiceServer(svc ReferenceService, opts ...interface{}) TwirpServer {
	serverOpts := twirp.ServerOptions{}
	for _, opt := range opts {
		switch o := opt.(type) {
		case twirp.ServerOption:
			o(&serverOpts)
		case *twirp.ServerHooks: // backwards compatibility, allow to specify hooks as an argument
			twirp.WithServerHooks(o)(&serverOpts)
		case nil: // backwards compatibility, allow nil value for the argument
			continue
		default:
			panic(fmt.Sprintf("Invalid option type %T on NewReferenceServiceServer", o))
		}
	}

	return &referenceServiceServer{
		ReferenceService: svc,
		pathPrefix:       serverOpts.PathPrefix(),
		interceptor:      twirp.ChainInterceptors(serverOpts.Interceptors...),
		hooks:            serverOpts.Hooks,
		jsonSkipDefaults: serverOpts.JSONSkipDefaults,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *referenceServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *referenceServiceServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
	if context.Canceled == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.Canceled, "failed to read request: context canceled"))
		return
	}
	if context.DeadlineExceeded == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.DeadlineExceeded, "failed to read request: deadline exceeded"))
		return
	}
	s.writeError(ctx, resp, twirp.WrapError(malformedRequestError(msg), err))
}

// ReferenceServicePathPrefix is a convenience constant that could used to identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// that add a "/twirp" prefix by default, and use CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const ReferenceServicePathPrefix = "/twirp/buf.alpha.registry.v1alpha1.ReferenceService/"

func (s *referenceServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "buf.alpha.registry.v1alpha1")
	ctx = ctxsetters.WithServiceName(ctx, "ReferenceService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	// Verify path format: [<prefix>]/<package>.<Service>/<Method>
	prefix, pkgService, method := parseTwirpPath(req.URL.Path)
	if pkgService != "buf.alpha.registry.v1alpha1.ReferenceService" {
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
	if prefix != s.pathPrefix {
		msg := fmt.Sprintf("invalid path prefix %q, expected %q, on path %q", prefix, s.pathPrefix, req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	switch method {
	case "GetReferenceByName":
		s.serveGetReferenceByName(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *referenceServiceServer) serveGetReferenceByName(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetReferenceByNameJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveGetReferenceByNameProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *referenceServiceServer) serveGetReferenceByNameJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "GetReferenceByName")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(GetReferenceByNameRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.ReferenceService.GetReferenceByName
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *GetReferenceByNameRequest) (*GetReferenceByNameResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*GetReferenceByNameRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*GetReferenceByNameRequest) when calling interceptor")
					}
					return s.ReferenceService.GetReferenceByName(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*GetReferenceByNameResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*GetReferenceByNameResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *GetReferenceByNameResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *GetReferenceByNameResponse and nil error while calling GetReferenceByName. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: true, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *referenceServiceServer) serveGetReferenceByNameProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "GetReferenceByName")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(GetReferenceByNameRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.ReferenceService.GetReferenceByName
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *GetReferenceByNameRequest) (*GetReferenceByNameResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*GetReferenceByNameRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*GetReferenceByNameRequest) when calling interceptor")
					}
					return s.ReferenceService.GetReferenceByName(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*GetReferenceByNameResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*GetReferenceByNameResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *GetReferenceByNameResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *GetReferenceByNameResponse and nil error while calling GetReferenceByName. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *referenceServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor8, 0
}

func (s *referenceServiceServer) ProtocGenTwirpVersion() string {
	return "v8.0.0"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *referenceServiceServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "buf.alpha.registry.v1alpha1", "ReferenceService")
}

var twirpFileDescriptor8 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0xcb, 0xd3, 0x40,
	0x10, 0xc6, 0x1b, 0x5b, 0x0b, 0xd9, 0x42, 0x95, 0xc5, 0x43, 0x8d, 0x07, 0x4b, 0x0e, 0x2a, 0x8a,
	0x59, 0xdb, 0x82, 0x1e, 0x04, 0x0f, 0x51, 0x68, 0x4f, 0x22, 0xab, 0xa7, 0x5e, 0x64, 0x37, 0x4e,
	0xd2, 0x85, 0x66, 0x37, 0x6e, 0x36, 0x95, 0x7e, 0x03, 0x41, 0x3c, 0x7b, 0xf5, 0xf3, 0xf9, 0x29,
	0x24, 0x9b, 0x7f, 0x85, 0xf7, 0x7d, 0xf3, 0xb6, 0xb7, 0xcd, 0xcc, 0xfc, 0x9e, 0x99, 0x67, 0x32,
	0xe8, 0x05, 0x2f, 0x62, 0xc2, 0xf6, 0xd9, 0x8e, 0x11, 0x0d, 0x89, 0xc8, 0x8d, 0x3e, 0x92, 0xc3,
	0xc2, 0x06, 0x16, 0x44, 0x43, 0x0c, 0x1a, 0x64, 0x04, 0x41, 0xa6, 0x95, 0x51, 0xf8, 0x11, 0x2f,
	0xe2, 0xc0, 0xe6, 0x82, 0xa6, 0x38, 0x68, 0x8a, 0xbd, 0x79, 0xa7, 0xc4, 0x32, 0xd1, 0x89, 0xb0,
	0x4c, 0x54, 0xb8, 0xb7, 0xea, 0xef, 0x95, 0xa9, 0x5c, 0x18, 0xa5, 0x8f, 0x5f, 0xb9, 0x66, 0x32,
	0xda, 0x5d, 0x08, 0x45, 0x2a, 0x4d, 0x85, 0xa9, 0xa1, 0x57, 0x67, 0x42, 0x86, 0x25, 0x15, 0xe1,
	0xff, 0x73, 0x90, 0x4b, 0x1b, 0xbb, 0x78, 0x8d, 0xc6, 0xd5, 0x10, 0x33, 0x67, 0xee, 0x3c, 0x9b,
	0x2c, 0x5f, 0x06, 0x3d, 0xce, 0x03, 0xda, 0x0a, 0x86, 0x16, 0xda, 0x0c, 0x68, 0x8d, 0xe3, 0x77,
	0x68, 0x68, 0x58, 0x32, 0xbb, 0x63, 0x55, 0x9e, 0x9f, 0xa9, 0xf2, 0x85, 0x25, 0x9b, 0x01, 0x2d,
	0xc1, 0x72, 0x90, 0xca, 0xd8, 0x6c, 0x78, 0xd1, 0x20, 0xef, 0x2d, 0x54, 0x0e, 0x52, 0xe1, 0xe1,
	0x04, 0xb9, 0xed, 0xdf, 0xf4, 0x25, 0x7a, 0xb8, 0x06, 0xd3, 0xda, 0x0d, 0x8f, 0x1f, 0x59, 0x0a,
	0x14, 0xbe, 0x17, 0x90, 0x1b, 0x8c, 0xd1, 0x48, 0xb2, 0x14, 0xac, 0x73, 0x97, 0xda, 0x37, 0x7e,
	0x80, 0xee, 0xaa, 0x1f, 0x12, 0xb4, 0x35, 0xe2, 0xd2, 0xea, 0x03, 0x3f, 0x45, 0xf7, 0x4e, 0x76,
	0x69, 0xa1, 0xa1, 0xcd, 0x4f, 0xbb, 0x70, 0xa9, 0xec, 0x73, 0xe4, 0x5d, 0xd7, 0x2f, 0xcf, 0x94,
	0xcc, 0x01, 0x7f, 0x38, 0x19, 0xad, 0xde, 0xf7, 0x93, 0x5b, 0x6c, 0xd6, 0xd5, 0xb4, 0x03, 0x97,
	0x7f, 0x1d, 0x74, 0xbf, 0x4d, 0x7c, 0x06, 0x7d, 0x10, 0x11, 0xe0, 0x5f, 0x0e, 0xc2, 0x57, 0x3b,
	0xe3, 0xd7, 0xbd, 0xf2, 0x37, 0xae, 0xc6, 0x7b, 0x73, 0x31, 0x57, 0x59, 0xf4, 0x47, 0x3f, 0xff,
	0xf8, 0x4e, 0xf8, 0xdb, 0x41, 0x8f, 0x23, 0x95, 0xf6, 0x89, 0x84, 0xd3, 0x56, 0xe2, 0x53, 0x79,
	0x97, 0xdb, 0x6d, 0x22, 0xcc, 0xae, 0xe0, 0x41, 0xa4, 0x52, 0xc2, 0x8b, 0x98, 0x17, 0x62, 0xff,
	0xad, 0x7c, 0x10, 0x21, 0x0d, 0x68, 0xc9, 0xf6, 0x24, 0x01, 0x49, 0xec, 0x0d, 0x93, 0x44, 0x91,
	0x9e, 0xbb, 0x7f, 0xdb, 0x44, 0x9a, 0x00, 0x1f, 0x5b, 0x6c, 0xf5, 0x3f, 0x00, 0x00, 0xff, 0xff,
	0x2a, 0xe4, 0xde, 0x9f, 0x04, 0x04, 0x00, 0x00,
}
