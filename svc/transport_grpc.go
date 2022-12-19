// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 5f7d5bf015
// Version Date: 2021-11-26T09:27:01Z

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/mises-id/mises-websitesvc/proto"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC WebsitesvcServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.WebsitesvcServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// websitesvc

		websitecategorylist: grpctransport.NewServer(
			endpoints.WebsiteCategoryListEndpoint,
			DecodeGRPCWebsiteCategoryListRequest,
			EncodeGRPCWebsiteCategoryListResponse,
			serverOptions...,
		),
		websitepage: grpctransport.NewServer(
			endpoints.WebsitePageEndpoint,
			DecodeGRPCWebsitePageRequest,
			EncodeGRPCWebsitePageResponse,
			serverOptions...,
		),
		websitesearch: grpctransport.NewServer(
			endpoints.WebsiteSearchEndpoint,
			DecodeGRPCWebsiteSearchRequest,
			EncodeGRPCWebsiteSearchResponse,
			serverOptions...,
		),
		websiterecommend: grpctransport.NewServer(
			endpoints.WebsiteRecommendEndpoint,
			DecodeGRPCWebsiteRecommendRequest,
			EncodeGRPCWebsiteRecommendResponse,
			serverOptions...,
		),
		websiteimport: grpctransport.NewServer(
			endpoints.WebsiteImportEndpoint,
			DecodeGRPCWebsiteImportRequest,
			EncodeGRPCWebsiteImportResponse,
			serverOptions...,
		),
		updatemetamaskphishing: grpctransport.NewServer(
			endpoints.UpdateMetaMaskPhishingEndpoint,
			DecodeGRPCUpdateMetaMaskPhishingRequest,
			EncodeGRPCUpdateMetaMaskPhishingResponse,
			serverOptions...,
		),
		updatephishingsiteblackorigin: grpctransport.NewServer(
			endpoints.UpdatePhishingSiteBlackOriginEndpoint,
			DecodeGRPCUpdatePhishingSiteBlackOriginRequest,
			EncodeGRPCUpdatePhishingSiteBlackOriginResponse,
			serverOptions...,
		),
		updatephishingoriginbywebsite: grpctransport.NewServer(
			endpoints.UpdatePhishingOriginByWebSiteEndpoint,
			DecodeGRPCUpdatePhishingOriginByWebSiteRequest,
			EncodeGRPCUpdatePhishingOriginByWebSiteResponse,
			serverOptions...,
		),
		updatephishingsitebywebsite: grpctransport.NewServer(
			endpoints.UpdatePhishingSiteByWebsiteEndpoint,
			DecodeGRPCUpdatePhishingSiteByWebsiteRequest,
			EncodeGRPCUpdatePhishingSiteByWebsiteResponse,
			serverOptions...,
		),
		phishingcheck: grpctransport.NewServer(
			endpoints.PhishingCheckEndpoint,
			DecodeGRPCPhishingCheckRequest,
			EncodeGRPCPhishingCheckResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the WebsitesvcServer interface
type grpcServer struct {
	websitecategorylist           grpctransport.Handler
	websitepage                   grpctransport.Handler
	websitesearch                 grpctransport.Handler
	websiterecommend              grpctransport.Handler
	websiteimport                 grpctransport.Handler
	updatemetamaskphishing        grpctransport.Handler
	updatephishingsiteblackorigin grpctransport.Handler
	updatephishingoriginbywebsite grpctransport.Handler
	updatephishingsitebywebsite   grpctransport.Handler
	phishingcheck                 grpctransport.Handler
}

// Methods for grpcServer to implement WebsitesvcServer interface

func (s *grpcServer) WebsiteCategoryList(ctx context.Context, req *pb.WebsiteCategoryListRequest) (*pb.WebsiteCategoryListResponse, error) {
	_, rep, err := s.websitecategorylist.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.WebsiteCategoryListResponse), nil
}

func (s *grpcServer) WebsitePage(ctx context.Context, req *pb.WebsitePageRequest) (*pb.WebsitePageResponse, error) {
	_, rep, err := s.websitepage.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.WebsitePageResponse), nil
}

func (s *grpcServer) WebsiteSearch(ctx context.Context, req *pb.WebsiteSearchRequest) (*pb.WebsiteSearchResponse, error) {
	_, rep, err := s.websitesearch.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.WebsiteSearchResponse), nil
}

func (s *grpcServer) WebsiteRecommend(ctx context.Context, req *pb.WebsiteRecommendRequest) (*pb.WebsiteRecommendResponse, error) {
	_, rep, err := s.websiterecommend.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.WebsiteRecommendResponse), nil
}

func (s *grpcServer) WebsiteImport(ctx context.Context, req *pb.WebsiteImportRequest) (*pb.WebsiteImportResponse, error) {
	_, rep, err := s.websiteimport.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.WebsiteImportResponse), nil
}

func (s *grpcServer) UpdateMetaMaskPhishing(ctx context.Context, req *pb.UpdateMetaMaskPhishingRequest) (*pb.UpdateMetaMaskPhishingResponse, error) {
	_, rep, err := s.updatemetamaskphishing.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateMetaMaskPhishingResponse), nil
}

func (s *grpcServer) UpdatePhishingSiteBlackOrigin(ctx context.Context, req *pb.UpdatePhishingSiteBlackOriginRequest) (*pb.UpdatePhishingSiteBlackOriginResponse, error) {
	_, rep, err := s.updatephishingsiteblackorigin.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdatePhishingSiteBlackOriginResponse), nil
}

func (s *grpcServer) UpdatePhishingOriginByWebSite(ctx context.Context, req *pb.UpdatePhishingOriginByWebSiteRequest) (*pb.UpdatePhishingOriginByWebSiteResponse, error) {
	_, rep, err := s.updatephishingoriginbywebsite.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdatePhishingOriginByWebSiteResponse), nil
}

func (s *grpcServer) UpdatePhishingSiteByWebsite(ctx context.Context, req *pb.UpdatePhishingSiteByWebsiteRequest) (*pb.UpdatePhishingSiteByWebsiteResponse, error) {
	_, rep, err := s.updatephishingsitebywebsite.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdatePhishingSiteByWebsiteResponse), nil
}

func (s *grpcServer) PhishingCheck(ctx context.Context, req *pb.PhishingCheckRequest) (*pb.PhishingCheckResponse, error) {
	_, rep, err := s.phishingcheck.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.PhishingCheckResponse), nil
}

// Server Decode

// DecodeGRPCWebsiteCategoryListRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC websitecategorylist request to a user-domain websitecategorylist request. Primarily useful in a server.
func DecodeGRPCWebsiteCategoryListRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.WebsiteCategoryListRequest)
	return req, nil
}

// DecodeGRPCWebsitePageRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC websitepage request to a user-domain websitepage request. Primarily useful in a server.
func DecodeGRPCWebsitePageRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.WebsitePageRequest)
	return req, nil
}

// DecodeGRPCWebsiteSearchRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC websitesearch request to a user-domain websitesearch request. Primarily useful in a server.
func DecodeGRPCWebsiteSearchRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.WebsiteSearchRequest)
	return req, nil
}

// DecodeGRPCWebsiteRecommendRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC websiterecommend request to a user-domain websiterecommend request. Primarily useful in a server.
func DecodeGRPCWebsiteRecommendRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.WebsiteRecommendRequest)
	return req, nil
}

// DecodeGRPCWebsiteImportRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC websiteimport request to a user-domain websiteimport request. Primarily useful in a server.
func DecodeGRPCWebsiteImportRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.WebsiteImportRequest)
	return req, nil
}

// DecodeGRPCUpdateMetaMaskPhishingRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC updatemetamaskphishing request to a user-domain updatemetamaskphishing request. Primarily useful in a server.
func DecodeGRPCUpdateMetaMaskPhishingRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UpdateMetaMaskPhishingRequest)
	return req, nil
}

// DecodeGRPCUpdatePhishingSiteBlackOriginRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC updatephishingsiteblackorigin request to a user-domain updatephishingsiteblackorigin request. Primarily useful in a server.
func DecodeGRPCUpdatePhishingSiteBlackOriginRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UpdatePhishingSiteBlackOriginRequest)
	return req, nil
}

// DecodeGRPCUpdatePhishingOriginByWebSiteRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC updatephishingoriginbywebsite request to a user-domain updatephishingoriginbywebsite request. Primarily useful in a server.
func DecodeGRPCUpdatePhishingOriginByWebSiteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UpdatePhishingOriginByWebSiteRequest)
	return req, nil
}

// DecodeGRPCUpdatePhishingSiteByWebsiteRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC updatephishingsitebywebsite request to a user-domain updatephishingsitebywebsite request. Primarily useful in a server.
func DecodeGRPCUpdatePhishingSiteByWebsiteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UpdatePhishingSiteByWebsiteRequest)
	return req, nil
}

// DecodeGRPCPhishingCheckRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC phishingcheck request to a user-domain phishingcheck request. Primarily useful in a server.
func DecodeGRPCPhishingCheckRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.PhishingCheckRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCWebsiteCategoryListResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain websitecategorylist response to a gRPC websitecategorylist reply. Primarily useful in a server.
func EncodeGRPCWebsiteCategoryListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.WebsiteCategoryListResponse)
	return resp, nil
}

// EncodeGRPCWebsitePageResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain websitepage response to a gRPC websitepage reply. Primarily useful in a server.
func EncodeGRPCWebsitePageResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.WebsitePageResponse)
	return resp, nil
}

// EncodeGRPCWebsiteSearchResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain websitesearch response to a gRPC websitesearch reply. Primarily useful in a server.
func EncodeGRPCWebsiteSearchResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.WebsiteSearchResponse)
	return resp, nil
}

// EncodeGRPCWebsiteRecommendResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain websiterecommend response to a gRPC websiterecommend reply. Primarily useful in a server.
func EncodeGRPCWebsiteRecommendResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.WebsiteRecommendResponse)
	return resp, nil
}

// EncodeGRPCWebsiteImportResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain websiteimport response to a gRPC websiteimport reply. Primarily useful in a server.
func EncodeGRPCWebsiteImportResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.WebsiteImportResponse)
	return resp, nil
}

// EncodeGRPCUpdateMetaMaskPhishingResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain updatemetamaskphishing response to a gRPC updatemetamaskphishing reply. Primarily useful in a server.
func EncodeGRPCUpdateMetaMaskPhishingResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.UpdateMetaMaskPhishingResponse)
	return resp, nil
}

// EncodeGRPCUpdatePhishingSiteBlackOriginResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain updatephishingsiteblackorigin response to a gRPC updatephishingsiteblackorigin reply. Primarily useful in a server.
func EncodeGRPCUpdatePhishingSiteBlackOriginResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.UpdatePhishingSiteBlackOriginResponse)
	return resp, nil
}

// EncodeGRPCUpdatePhishingOriginByWebSiteResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain updatephishingoriginbywebsite response to a gRPC updatephishingoriginbywebsite reply. Primarily useful in a server.
func EncodeGRPCUpdatePhishingOriginByWebSiteResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.UpdatePhishingOriginByWebSiteResponse)
	return resp, nil
}

// EncodeGRPCUpdatePhishingSiteByWebsiteResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain updatephishingsitebywebsite response to a gRPC updatephishingsitebywebsite reply. Primarily useful in a server.
func EncodeGRPCUpdatePhishingSiteByWebsiteResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.UpdatePhishingSiteByWebsiteResponse)
	return resp, nil
}

// EncodeGRPCPhishingCheckResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain phishingcheck response to a gRPC phishingcheck reply. Primarily useful in a server.
func EncodeGRPCPhishingCheckResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.PhishingCheckResponse)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
