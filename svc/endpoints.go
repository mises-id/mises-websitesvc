// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 5f7d5bf015
// Version Date: Fri Nov 26 09:27:01 UTC 2021

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	pb "github.com/mises-id/mises-websitesvc/proto"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	WebsiteCategoryListEndpoint endpoint.Endpoint
	WebsitePageEndpoint         endpoint.Endpoint
	WebsiteRecommendEndpoint    endpoint.Endpoint
}

// Endpoints

func (e Endpoints) WebsiteCategoryList(ctx context.Context, in *pb.WebsiteCategoryListRequest) (*pb.WebsiteCategoryListResponse, error) {
	response, err := e.WebsiteCategoryListEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.WebsiteCategoryListResponse), nil
}

func (e Endpoints) WebsitePage(ctx context.Context, in *pb.WebsitePageRequest) (*pb.WebsitePageResponse, error) {
	response, err := e.WebsitePageEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.WebsitePageResponse), nil
}

func (e Endpoints) WebsiteRecommend(ctx context.Context, in *pb.WebsiteRecommendRequest) (*pb.WebsiteRecommendResponse, error) {
	response, err := e.WebsiteRecommendEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.WebsiteRecommendResponse), nil
}

// Make Endpoints

func MakeWebsiteCategoryListEndpoint(s pb.WebsitesvcServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.WebsiteCategoryListRequest)
		v, err := s.WebsiteCategoryList(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeWebsitePageEndpoint(s pb.WebsitesvcServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.WebsitePageRequest)
		v, err := s.WebsitePage(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeWebsiteRecommendEndpoint(s pb.WebsitesvcServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.WebsiteRecommendRequest)
		v, err := s.WebsiteRecommend(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAllExcept wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllExcept(middleware, "Status", "Ping")
func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"WebsiteCategoryList": {},
		"WebsitePage":         {},
		"WebsiteRecommend":    {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "WebsiteCategoryList" {
			e.WebsiteCategoryListEndpoint = middleware(e.WebsiteCategoryListEndpoint)
		}
		if inc == "WebsitePage" {
			e.WebsitePageEndpoint = middleware(e.WebsitePageEndpoint)
		}
		if inc == "WebsiteRecommend" {
			e.WebsiteRecommendEndpoint = middleware(e.WebsiteRecommendEndpoint)
		}
	}
}

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(string, endpoint.Endpoint) endpoint.Endpoint

// WrapAllLabeledExcept wraps each Endpoint field of struct Endpoints with a
// LabeledMiddleware, which will receive the name of the endpoint. See
// LabeldMiddleware. See method WrapAllExept for details on excluded
// functionality.
func (e *Endpoints) WrapAllLabeledExcept(middleware func(string, endpoint.Endpoint) endpoint.Endpoint, excluded ...string) {
	included := map[string]struct{}{
		"WebsiteCategoryList": {},
		"WebsitePage":         {},
		"WebsiteRecommend":    {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "WebsiteCategoryList" {
			e.WebsiteCategoryListEndpoint = middleware("WebsiteCategoryList", e.WebsiteCategoryListEndpoint)
		}
		if inc == "WebsitePage" {
			e.WebsitePageEndpoint = middleware("WebsitePage", e.WebsitePageEndpoint)
		}
		if inc == "WebsiteRecommend" {
			e.WebsiteRecommendEndpoint = middleware("WebsiteRecommend", e.WebsiteRecommendEndpoint)
		}
	}
}
