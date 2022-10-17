// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 5f7d5bf015
// Version Date: 2021-11-26T09:27:01Z

// Package http provides an HTTP client for the Websitesvc service.
package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gogo/protobuf/jsonpb"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pkg/errors"

	// This Service
	pb "github.com/mises-id/mises-websitesvc/proto"
	"github.com/mises-id/mises-websitesvc/svc"
)

var (
	_ = endpoint.Chain
	_ = httptransport.NewClient
	_ = fmt.Sprint
	_ = bytes.Compare
	_ = ioutil.NopCloser
	_ = io.EOF
)

// New returns a service backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string, options ...httptransport.ClientOption) (pb.WebsitesvcServer, error) {

	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	_ = u

	var WebsiteCategoryListZeroEndpoint endpoint.Endpoint
	{
		WebsiteCategoryListZeroEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/website_category/list/"),
			EncodeHTTPWebsiteCategoryListZeroRequest,
			DecodeHTTPWebsiteCategoryListResponse,
			options...,
		).Endpoint()
	}
	var WebsitePageZeroEndpoint endpoint.Endpoint
	{
		WebsitePageZeroEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/website/page/"),
			EncodeHTTPWebsitePageZeroRequest,
			DecodeHTTPWebsitePageResponse,
			options...,
		).Endpoint()
	}
	var WebsiteRecommendZeroEndpoint endpoint.Endpoint
	{
		WebsiteRecommendZeroEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/website/recommend/"),
			EncodeHTTPWebsiteRecommendZeroRequest,
			DecodeHTTPWebsiteRecommendResponse,
			options...,
		).Endpoint()
	}
	var WebsiteImportZeroEndpoint endpoint.Endpoint
	{
		WebsiteImportZeroEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/website/import/"),
			EncodeHTTPWebsiteImportZeroRequest,
			DecodeHTTPWebsiteImportResponse,
			options...,
		).Endpoint()
	}

	return svc.Endpoints{
		WebsiteCategoryListEndpoint: WebsiteCategoryListZeroEndpoint,
		WebsitePageEndpoint:         WebsitePageZeroEndpoint,
		WebsiteRecommendEndpoint:    WebsiteRecommendZeroEndpoint,
		WebsiteImportEndpoint:       WebsiteImportZeroEndpoint,
	}, nil
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}

// CtxValuesToSend configures the http client to pull the specified keys out of
// the context and add them to the http request as headers.  Note that keys
// will have net/http.CanonicalHeaderKey called on them before being send over
// the wire and that is the form they will be available in the server context.
func CtxValuesToSend(keys ...string) httptransport.ClientOption {
	return httptransport.ClientBefore(func(ctx context.Context, r *http.Request) context.Context {
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				r.Header.Set(k, v)
			}
		}
		return ctx
	})
}

// HTTP Client Decode

// DecodeHTTPWebsiteCategoryListResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded WebsiteCategoryListResponse response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPWebsiteCategoryListResponse(_ context.Context, r *http.Response) (interface{}, error) {
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err == io.EOF {
		return nil, errors.New("response http body empty")
	}
	if err != nil {
		return nil, errors.Wrap(err, "cannot read http body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errorDecoder(buf), "status code: '%d'", r.StatusCode)
	}

	var resp pb.WebsiteCategoryListResponse
	if err = jsonpb.UnmarshalString(string(buf), &resp); err != nil {
		return nil, errorDecoder(buf)
	}

	return &resp, nil
}

// DecodeHTTPWebsitePageResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded WebsitePageResponse response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPWebsitePageResponse(_ context.Context, r *http.Response) (interface{}, error) {
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err == io.EOF {
		return nil, errors.New("response http body empty")
	}
	if err != nil {
		return nil, errors.Wrap(err, "cannot read http body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errorDecoder(buf), "status code: '%d'", r.StatusCode)
	}

	var resp pb.WebsitePageResponse
	if err = jsonpb.UnmarshalString(string(buf), &resp); err != nil {
		return nil, errorDecoder(buf)
	}

	return &resp, nil
}

// DecodeHTTPWebsiteRecommendResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded WebsiteRecommendResponse response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPWebsiteRecommendResponse(_ context.Context, r *http.Response) (interface{}, error) {
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err == io.EOF {
		return nil, errors.New("response http body empty")
	}
	if err != nil {
		return nil, errors.Wrap(err, "cannot read http body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errorDecoder(buf), "status code: '%d'", r.StatusCode)
	}

	var resp pb.WebsiteRecommendResponse
	if err = jsonpb.UnmarshalString(string(buf), &resp); err != nil {
		return nil, errorDecoder(buf)
	}

	return &resp, nil
}

// DecodeHTTPWebsiteImportResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded WebsiteImportResponse response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPWebsiteImportResponse(_ context.Context, r *http.Response) (interface{}, error) {
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err == io.EOF {
		return nil, errors.New("response http body empty")
	}
	if err != nil {
		return nil, errors.Wrap(err, "cannot read http body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errorDecoder(buf), "status code: '%d'", r.StatusCode)
	}

	var resp pb.WebsiteImportResponse
	if err = jsonpb.UnmarshalString(string(buf), &resp); err != nil {
		return nil, errorDecoder(buf)
	}

	return &resp, nil
}

// HTTP Client Encode

// EncodeHTTPWebsiteCategoryListZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a websitecategorylist request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPWebsiteCategoryListZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.WebsiteCategoryListRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"website_category",
		"list",
		"",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("type", fmt.Sprint(req.Type))

	values.Add("list_num", fmt.Sprint(req.ListNum))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPWebsiteCategoryListOneRequest is a transport/http.EncodeRequestFunc
// that encodes a websitecategorylist request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPWebsiteCategoryListOneRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.WebsiteCategoryListRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"website_category",
		"list",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("type", fmt.Sprint(req.Type))

	values.Add("list_num", fmt.Sprint(req.ListNum))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPWebsitePageZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a websitepage request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPWebsitePageZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.WebsitePageRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"website",
		"page",
		"",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("type", fmt.Sprint(req.Type))

	values.Add("website_category_id", fmt.Sprint(req.WebsiteCategoryId))

	values.Add("Keywords", fmt.Sprint(req.Keywords))

	tmp, err = json.Marshal(req.Paginator)
	if err != nil {
		return errors.Wrap(err, "failed to marshal req.Paginator")
	}
	strval = string(tmp)
	values.Add("paginator", strval)

	values.Add("subcategory_id", fmt.Sprint(req.SubcategoryId))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPWebsitePageOneRequest is a transport/http.EncodeRequestFunc
// that encodes a websitepage request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPWebsitePageOneRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.WebsitePageRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"website",
		"page",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("type", fmt.Sprint(req.Type))

	values.Add("website_category_id", fmt.Sprint(req.WebsiteCategoryId))

	values.Add("Keywords", fmt.Sprint(req.Keywords))

	tmp, err = json.Marshal(req.Paginator)
	if err != nil {
		return errors.Wrap(err, "failed to marshal req.Paginator")
	}
	strval = string(tmp)
	values.Add("paginator", strval)

	values.Add("subcategory_id", fmt.Sprint(req.SubcategoryId))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPWebsiteRecommendZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a websiterecommend request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPWebsiteRecommendZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.WebsiteRecommendRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"website",
		"recommend",
		"",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("type", fmt.Sprint(req.Type))

	values.Add("list_num", fmt.Sprint(req.ListNum))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPWebsiteRecommendOneRequest is a transport/http.EncodeRequestFunc
// that encodes a websiterecommend request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPWebsiteRecommendOneRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.WebsiteRecommendRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"website",
		"recommend",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("type", fmt.Sprint(req.Type))

	values.Add("list_num", fmt.Sprint(req.ListNum))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPWebsiteImportZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a websiteimport request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPWebsiteImportZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.WebsiteImportRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"website",
		"import",
		"",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("file_path", fmt.Sprint(req.FilePath))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPWebsiteImportOneRequest is a transport/http.EncodeRequestFunc
// that encodes a websiteimport request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPWebsiteImportOneRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.WebsiteImportRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"website",
		"import",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("file_path", fmt.Sprint(req.FilePath))

	r.URL.RawQuery = values.Encode()
	return nil
}

func errorDecoder(buf []byte) error {
	var w errorWrapper
	if err := json.Unmarshal(buf, &w); err != nil {
		const size = 8196
		if len(buf) > size {
			buf = buf[:size]
		}
		return fmt.Errorf("response body '%s': cannot parse non-json request body", buf)
	}

	return errors.New(w.Error)
}

type errorWrapper struct {
	Error string `json:"error"`
}
