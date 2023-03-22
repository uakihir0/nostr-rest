// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// Paging defines model for Paging.
type Paging struct {
	// Specify when to get future posts from this result
	FutureSinceTime string `json:"future_since_time"`

	// Specify when to get past posts from this result
	PastUntileTime string `json:"past_untile_time"`
}

// Post defines model for Post.
type Post struct {
	// Raw text of post content.
	Content string `json:"content"`

	// Time post was created (UTC)
	CreatedAt string `json:"created_at"`
	Id        string `json:"id"`
	User      User   `json:"user"`
}

// Posts defines model for Posts.
type Posts struct {
	// Number of list
	Count int `json:"count"`

	// Post list
	List []Post `json:"list"`
}

// User's PublicKeys
type Pubkeys = []string

// User defines model for User.
type User struct {
	// User description
	About *string `json:"about,omitempty"`

	// User profile banner image url
	Banner *string `json:"banner,omitempty"`

	// User display name
	DisplayName *string `json:"display_name,omitempty"`

	// User name
	Name *string `json:"name,omitempty"`

	// User icon image url
	Picture *string `json:"picture,omitempty"`

	// User public key (user idenitifier)
	Pubkey string `json:"pubkey"`

	// User website url
	Website *string `json:"website,omitempty"`
}

// Users defines model for Users.
type Users struct {
	// Number of list
	Count int `json:"count"`

	// User list
	List []User `json:"list"`
}

// MaxResultsParameter defines model for MaxResultsParameter.
type MaxResultsParameter = int

// PubkeyParameter defines model for PubkeyParameter.
type PubkeyParameter = string

// SinceTimeParameter defines model for SinceTimeParameter.
type SinceTimeParameter = string

// UntilTimePatameter defines model for UntilTimePatameter.
type UntilTimePatameter = string

// PubKeysResponse defines model for PubKeysResponse.
type PubKeysResponse struct {
	// Number of pubkeys
	Count int `json:"count"`

	// User's PublicKeys
	Pubkeys Pubkeys `json:"pubkeys"`
}

// UserResponse defines model for UserResponse.
type UserResponse = User

// UsersResponse defines model for UsersResponse.
type UsersResponse = Users

// UsersTimelineResponse defines model for UsersTimelineResponse.
type UsersTimelineResponse struct {
	Paging *Paging `json:"paging,omitempty"`
	Posts  Posts   `json:"posts"`

	// User's PublicKeys
	Pubkeys Pubkeys `json:"pubkeys"`
}

// UsersPubKeyRequest defines model for UsersPubKeyRequest.
type UsersPubKeyRequest struct {
	// Public key of the user to retrieve
	Pubkeys []string `json:"pubkeys"`
}

// GetV1TimelinesHomeParams defines parameters for GetV1TimelinesHome.
type GetV1TimelinesHomeParams struct {
	// Public key of the user
	Pubkey PubkeyParameter `form:"pubkey" json:"pubkey"`

	// Specifies the number of Posts to try and retrieve (default 20)
	MaxResults *MaxResultsParameter `form:"max_results,omitempty" json:"max_results,omitempty"`

	// Get posts after that time (include)
	SinceTime *SinceTimeParameter `form:"since_time,omitempty" json:"since_time,omitempty"`

	// Get posts up to that time (exclude)
	UntilTime *UntilTimePatameter `form:"until_time,omitempty" json:"until_time,omitempty"`
}

// GetV1TimelinesUserParams defines parameters for GetV1TimelinesUser.
type GetV1TimelinesUserParams struct {
	// Public key of the user
	Pubkey PubkeyParameter `form:"pubkey" json:"pubkey"`

	// Specifies the number of Posts to try and retrieve (default 20)
	MaxResults *MaxResultsParameter `form:"max_results,omitempty" json:"max_results,omitempty"`

	// Get posts after that time (include)
	SinceTime *SinceTimeParameter `form:"since_time,omitempty" json:"since_time,omitempty"`

	// Get posts up to that time (exclude)
	UntilTime *UntilTimePatameter `form:"until_time,omitempty" json:"until_time,omitempty"`
}

// GetV1UsersParams defines parameters for GetV1Users.
type GetV1UsersParams struct {
	// Public key of the user to retrieve
	Pubkey string `form:"pubkey" json:"pubkey"`
}

// GetV1UsersFollowersParams defines parameters for GetV1UsersFollowers.
type GetV1UsersFollowersParams struct {
	// Public key of the user
	Pubkey PubkeyParameter `form:"pubkey" json:"pubkey"`
}

// GetV1UsersFollowingParams defines parameters for GetV1UsersFollowing.
type GetV1UsersFollowingParams struct {
	// Public key of the user
	Pubkey PubkeyParameter `form:"pubkey" json:"pubkey"`
}

// GetV1UsersFollowingPubkeysParams defines parameters for GetV1UsersFollowingPubkeys.
type GetV1UsersFollowingPubkeysParams struct {
	// Public key of the user
	Pubkey PubkeyParameter `form:"pubkey" json:"pubkey"`
}

// PostV1UsersJSONRequestBody defines body for PostV1Users for application/json ContentType.
type PostV1UsersJSONRequestBody UsersPubKeyRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get Home Timeline
	// (GET /v1/timelines/home)
	GetV1TimelinesHome(ctx echo.Context, params GetV1TimelinesHomeParams) error
	// Get User Timeline
	// (GET /v1/timelines/user)
	GetV1TimelinesUser(ctx echo.Context, params GetV1TimelinesUserParams) error
	// GET User Profiles
	// (GET /v1/users)
	GetV1Users(ctx echo.Context, params GetV1UsersParams) error
	// GET Users Profiles
	// (POST /v1/users)
	PostV1Users(ctx echo.Context) error
	// Get User's Followers
	// (GET /v1/users/followers)
	GetV1UsersFollowers(ctx echo.Context, params GetV1UsersFollowersParams) error
	// Get Following Users
	// (GET /v1/users/following)
	GetV1UsersFollowing(ctx echo.Context, params GetV1UsersFollowingParams) error
	// Get Following User's PublicKeys
	// (GET /v1/users/following/pubkeys)
	GetV1UsersFollowingPubkeys(ctx echo.Context, params GetV1UsersFollowingPubkeysParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetV1TimelinesHome converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1TimelinesHome(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetV1TimelinesHomeParams
	// ------------- Required query parameter "pubkey" -------------

	err = runtime.BindQueryParameter("form", true, true, "pubkey", ctx.QueryParams(), &params.Pubkey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pubkey: %s", err))
	}

	// ------------- Optional query parameter "max_results" -------------

	err = runtime.BindQueryParameter("form", true, false, "max_results", ctx.QueryParams(), &params.MaxResults)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max_results: %s", err))
	}

	// ------------- Optional query parameter "since_time" -------------

	err = runtime.BindQueryParameter("form", true, false, "since_time", ctx.QueryParams(), &params.SinceTime)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter since_time: %s", err))
	}

	// ------------- Optional query parameter "until_time" -------------

	err = runtime.BindQueryParameter("form", true, false, "until_time", ctx.QueryParams(), &params.UntilTime)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter until_time: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetV1TimelinesHome(ctx, params)
	return err
}

// GetV1TimelinesUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1TimelinesUser(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetV1TimelinesUserParams
	// ------------- Required query parameter "pubkey" -------------

	err = runtime.BindQueryParameter("form", true, true, "pubkey", ctx.QueryParams(), &params.Pubkey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pubkey: %s", err))
	}

	// ------------- Optional query parameter "max_results" -------------

	err = runtime.BindQueryParameter("form", true, false, "max_results", ctx.QueryParams(), &params.MaxResults)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max_results: %s", err))
	}

	// ------------- Optional query parameter "since_time" -------------

	err = runtime.BindQueryParameter("form", true, false, "since_time", ctx.QueryParams(), &params.SinceTime)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter since_time: %s", err))
	}

	// ------------- Optional query parameter "until_time" -------------

	err = runtime.BindQueryParameter("form", true, false, "until_time", ctx.QueryParams(), &params.UntilTime)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter until_time: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetV1TimelinesUser(ctx, params)
	return err
}

// GetV1Users converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1Users(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetV1UsersParams
	// ------------- Required query parameter "pubkey" -------------

	err = runtime.BindQueryParameter("form", true, true, "pubkey", ctx.QueryParams(), &params.Pubkey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pubkey: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetV1Users(ctx, params)
	return err
}

// PostV1Users converts echo context to params.
func (w *ServerInterfaceWrapper) PostV1Users(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostV1Users(ctx)
	return err
}

// GetV1UsersFollowers converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1UsersFollowers(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetV1UsersFollowersParams
	// ------------- Required query parameter "pubkey" -------------

	err = runtime.BindQueryParameter("form", true, true, "pubkey", ctx.QueryParams(), &params.Pubkey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pubkey: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetV1UsersFollowers(ctx, params)
	return err
}

// GetV1UsersFollowing converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1UsersFollowing(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetV1UsersFollowingParams
	// ------------- Required query parameter "pubkey" -------------

	err = runtime.BindQueryParameter("form", true, true, "pubkey", ctx.QueryParams(), &params.Pubkey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pubkey: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetV1UsersFollowing(ctx, params)
	return err
}

// GetV1UsersFollowingPubkeys converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1UsersFollowingPubkeys(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetV1UsersFollowingPubkeysParams
	// ------------- Required query parameter "pubkey" -------------

	err = runtime.BindQueryParameter("form", true, true, "pubkey", ctx.QueryParams(), &params.Pubkey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pubkey: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetV1UsersFollowingPubkeys(ctx, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/v1/timelines/home", wrapper.GetV1TimelinesHome)
	router.GET(baseURL+"/v1/timelines/user", wrapper.GetV1TimelinesUser)
	router.GET(baseURL+"/v1/users", wrapper.GetV1Users)
	router.POST(baseURL+"/v1/users", wrapper.PostV1Users)
	router.GET(baseURL+"/v1/users/followers", wrapper.GetV1UsersFollowers)
	router.GET(baseURL+"/v1/users/following", wrapper.GetV1UsersFollowing)
	router.GET(baseURL+"/v1/users/following/pubkeys", wrapper.GetV1UsersFollowingPubkeys)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
