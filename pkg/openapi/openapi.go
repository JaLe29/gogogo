// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
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
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// Allow defines model for Allow.
type Allow struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        *string    `json:"id,omitempty"`
	Ip        *string    `json:"ip,omitempty"`
	ProxyId   *string    `json:"proxyId,omitempty"`
}

// Block defines model for Block.
type Block struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        *string    `json:"id,omitempty"`
	Ip        *string    `json:"ip,omitempty"`
	ProxyId   *string    `json:"proxyId,omitempty"`
}

// Guard defines model for Guard.
type Guard struct {
	Email string `json:"email"`
	Id    string `json:"id"`
}

// GuardExluce defines model for GuardExluce.
type GuardExluce struct {
	Id   string `json:"id"`
	Path string `json:"path"`
}

// NewAllow defines model for NewAllow.
type NewAllow struct {
	Ip string `json:"ip"`
}

// NewBlock defines model for NewBlock.
type NewBlock struct {
	Ip string `json:"ip"`
}

// NewGuard defines model for NewGuard.
type NewGuard struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewGuardExluce defines model for NewGuardExluce.
type NewGuardExluce struct {
	Path string `json:"path"`
}

// NewProxy defines model for NewProxy.
type NewProxy struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

// PatchProxy defines model for PatchProxy.
type PatchProxy struct {
	Cache   bool `json:"cache"`
	Disable bool `json:"disable"`
}

// Proxy defines model for Proxy.
type Proxy struct {
	Cache     bool      `json:"cache"`
	CreatedAt time.Time `json:"createdAt"`
	Disable   bool      `json:"disable"`
	Id        string    `json:"id"`
	Source    string    `json:"source"`
	Target    string    `json:"target"`
}

// SuccessResponse defines model for SuccessResponse.
type SuccessResponse struct {
	// Message Message of the response
	Message string `json:"message"`
}

// DeleteApiAllowProxyIdParams defines parameters for DeleteApiAllowProxyId.
type DeleteApiAllowProxyIdParams struct {
	// Id Id of the allow
	Id string `form:"id" json:"id"`
}

// DeleteApiBlockProxyIdParams defines parameters for DeleteApiBlockProxyId.
type DeleteApiBlockProxyIdParams struct {
	// Id Id of the block
	Id string `form:"id" json:"id"`
}

// DeleteApiGuardExcludeGuardIdParams defines parameters for DeleteApiGuardExcludeGuardId.
type DeleteApiGuardExcludeGuardIdParams struct {
	// Id Id of the block
	Id string `form:"id" json:"id"`
}

// DeleteApiGuardProxyIdParams defines parameters for DeleteApiGuardProxyId.
type DeleteApiGuardProxyIdParams struct {
	// Id Id of the block
	Id string `form:"id" json:"id"`
}

// DeleteApiProxyParams defines parameters for DeleteApiProxy.
type DeleteApiProxyParams struct {
	// Id Id of the proxy
	Id string `form:"id" json:"id"`
}

// PatchApiProxyParams defines parameters for PatchApiProxy.
type PatchApiProxyParams struct {
	// Id Id of the proxy
	Id string `form:"id" json:"id"`
}

// PostApiAllowProxyIdJSONRequestBody defines body for PostApiAllowProxyId for application/json ContentType.
type PostApiAllowProxyIdJSONRequestBody = NewAllow

// PostApiBlockProxyIdJSONRequestBody defines body for PostApiBlockProxyId for application/json ContentType.
type PostApiBlockProxyIdJSONRequestBody = NewBlock

// PostApiGuardExcludeGuardIdJSONRequestBody defines body for PostApiGuardExcludeGuardId for application/json ContentType.
type PostApiGuardExcludeGuardIdJSONRequestBody = NewGuardExluce

// PostApiGuardProxyIdJSONRequestBody defines body for PostApiGuardProxyId for application/json ContentType.
type PostApiGuardProxyIdJSONRequestBody = NewGuard

// PatchApiProxyJSONRequestBody defines body for PatchApiProxy for application/json ContentType.
type PatchApiProxyJSONRequestBody = PatchProxy

// PostApiProxyJSONRequestBody defines body for PostApiProxy for application/json ContentType.
type PostApiProxyJSONRequestBody = NewProxy

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Delete a allow
	// (DELETE /api/allow/{proxyId})
	DeleteApiAllowProxyId(c *gin.Context, proxyId string, params DeleteApiAllowProxyIdParams)
	// Get all allows
	// (GET /api/allow/{proxyId})
	GetApiAllowProxyId(c *gin.Context, proxyId string)
	// Create a new allow
	// (POST /api/allow/{proxyId})
	PostApiAllowProxyId(c *gin.Context, proxyId string)
	// Delete a block
	// (DELETE /api/block/{proxyId})
	DeleteApiBlockProxyId(c *gin.Context, proxyId string, params DeleteApiBlockProxyIdParams)
	// Get all blocks
	// (GET /api/block/{proxyId})
	GetApiBlockProxyId(c *gin.Context, proxyId string)
	// Create a new block
	// (POST /api/block/{proxyId})
	PostApiBlockProxyId(c *gin.Context, proxyId string)
	// Disable guard exclude
	// (DELETE /api/guard-exclude/{guardId})
	DeleteApiGuardExcludeGuardId(c *gin.Context, guardId string, params DeleteApiGuardExcludeGuardIdParams)
	// Get all guard exclude
	// (GET /api/guard-exclude/{guardId})
	GetApiGuardExcludeGuardId(c *gin.Context, guardId string)
	// Enable guard exclude
	// (POST /api/guard-exclude/{guardId})
	PostApiGuardExcludeGuardId(c *gin.Context, guardId string)
	// Disable guard
	// (DELETE /api/guard/{proxyId})
	DeleteApiGuardProxyId(c *gin.Context, proxyId string, params DeleteApiGuardProxyIdParams)
	// Get all guards
	// (GET /api/guard/{proxyId})
	GetApiGuardProxyId(c *gin.Context, proxyId string)
	// Enable guard
	// (POST /api/guard/{proxyId})
	PostApiGuardProxyId(c *gin.Context, proxyId string)
	// Delete a proxy
	// (DELETE /api/proxy)
	DeleteApiProxy(c *gin.Context, params DeleteApiProxyParams)
	// Get all proxies
	// (GET /api/proxy)
	GetApiProxy(c *gin.Context)
	// Update a proxy
	// (PATCH /api/proxy)
	PatchApiProxy(c *gin.Context, params PatchApiProxyParams)
	// Create a new proxy
	// (POST /api/proxy)
	PostApiProxy(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// DeleteApiAllowProxyId operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiAllowProxyId(c *gin.Context) {

	var err error

	// ------------- Path parameter "proxyId" -------------
	var proxyId string

	err = runtime.BindStyledParameterWithOptions("simple", "proxyId", c.Param("proxyId"), &proxyId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter proxyId: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteApiAllowProxyIdParams

	// ------------- Required query parameter "id" -------------

	if paramValue := c.Query("id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "id", c.Request.URL.Query(), &params.Id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteApiAllowProxyId(c, proxyId, params)
}

// GetApiAllowProxyId operation middleware
func (siw *ServerInterfaceWrapper) GetApiAllowProxyId(c *gin.Context) {

	var err error

	// ------------- Path parameter "proxyId" -------------
	var proxyId string

	err = runtime.BindStyledParameterWithOptions("simple", "proxyId", c.Param("proxyId"), &proxyId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter proxyId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiAllowProxyId(c, proxyId)
}

// PostApiAllowProxyId operation middleware
func (siw *ServerInterfaceWrapper) PostApiAllowProxyId(c *gin.Context) {

	var err error

	// ------------- Path parameter "proxyId" -------------
	var proxyId string

	err = runtime.BindStyledParameterWithOptions("simple", "proxyId", c.Param("proxyId"), &proxyId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter proxyId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiAllowProxyId(c, proxyId)
}

// DeleteApiBlockProxyId operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiBlockProxyId(c *gin.Context) {

	var err error

	// ------------- Path parameter "proxyId" -------------
	var proxyId string

	err = runtime.BindStyledParameterWithOptions("simple", "proxyId", c.Param("proxyId"), &proxyId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter proxyId: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteApiBlockProxyIdParams

	// ------------- Required query parameter "id" -------------

	if paramValue := c.Query("id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "id", c.Request.URL.Query(), &params.Id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteApiBlockProxyId(c, proxyId, params)
}

// GetApiBlockProxyId operation middleware
func (siw *ServerInterfaceWrapper) GetApiBlockProxyId(c *gin.Context) {

	var err error

	// ------------- Path parameter "proxyId" -------------
	var proxyId string

	err = runtime.BindStyledParameterWithOptions("simple", "proxyId", c.Param("proxyId"), &proxyId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter proxyId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiBlockProxyId(c, proxyId)
}

// PostApiBlockProxyId operation middleware
func (siw *ServerInterfaceWrapper) PostApiBlockProxyId(c *gin.Context) {

	var err error

	// ------------- Path parameter "proxyId" -------------
	var proxyId string

	err = runtime.BindStyledParameterWithOptions("simple", "proxyId", c.Param("proxyId"), &proxyId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter proxyId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiBlockProxyId(c, proxyId)
}

// DeleteApiGuardExcludeGuardId operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiGuardExcludeGuardId(c *gin.Context) {

	var err error

	// ------------- Path parameter "guardId" -------------
	var guardId string

	err = runtime.BindStyledParameterWithOptions("simple", "guardId", c.Param("guardId"), &guardId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter guardId: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteApiGuardExcludeGuardIdParams

	// ------------- Required query parameter "id" -------------

	if paramValue := c.Query("id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "id", c.Request.URL.Query(), &params.Id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteApiGuardExcludeGuardId(c, guardId, params)
}

// GetApiGuardExcludeGuardId operation middleware
func (siw *ServerInterfaceWrapper) GetApiGuardExcludeGuardId(c *gin.Context) {

	var err error

	// ------------- Path parameter "guardId" -------------
	var guardId string

	err = runtime.BindStyledParameterWithOptions("simple", "guardId", c.Param("guardId"), &guardId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter guardId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiGuardExcludeGuardId(c, guardId)
}

// PostApiGuardExcludeGuardId operation middleware
func (siw *ServerInterfaceWrapper) PostApiGuardExcludeGuardId(c *gin.Context) {

	var err error

	// ------------- Path parameter "guardId" -------------
	var guardId string

	err = runtime.BindStyledParameterWithOptions("simple", "guardId", c.Param("guardId"), &guardId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter guardId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiGuardExcludeGuardId(c, guardId)
}

// DeleteApiGuardProxyId operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiGuardProxyId(c *gin.Context) {

	var err error

	// ------------- Path parameter "proxyId" -------------
	var proxyId string

	err = runtime.BindStyledParameterWithOptions("simple", "proxyId", c.Param("proxyId"), &proxyId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter proxyId: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteApiGuardProxyIdParams

	// ------------- Required query parameter "id" -------------

	if paramValue := c.Query("id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "id", c.Request.URL.Query(), &params.Id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteApiGuardProxyId(c, proxyId, params)
}

// GetApiGuardProxyId operation middleware
func (siw *ServerInterfaceWrapper) GetApiGuardProxyId(c *gin.Context) {

	var err error

	// ------------- Path parameter "proxyId" -------------
	var proxyId string

	err = runtime.BindStyledParameterWithOptions("simple", "proxyId", c.Param("proxyId"), &proxyId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter proxyId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiGuardProxyId(c, proxyId)
}

// PostApiGuardProxyId operation middleware
func (siw *ServerInterfaceWrapper) PostApiGuardProxyId(c *gin.Context) {

	var err error

	// ------------- Path parameter "proxyId" -------------
	var proxyId string

	err = runtime.BindStyledParameterWithOptions("simple", "proxyId", c.Param("proxyId"), &proxyId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter proxyId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiGuardProxyId(c, proxyId)
}

// DeleteApiProxy operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiProxy(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteApiProxyParams

	// ------------- Required query parameter "id" -------------

	if paramValue := c.Query("id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "id", c.Request.URL.Query(), &params.Id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteApiProxy(c, params)
}

// GetApiProxy operation middleware
func (siw *ServerInterfaceWrapper) GetApiProxy(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiProxy(c)
}

// PatchApiProxy operation middleware
func (siw *ServerInterfaceWrapper) PatchApiProxy(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PatchApiProxyParams

	// ------------- Required query parameter "id" -------------

	if paramValue := c.Query("id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "id", c.Request.URL.Query(), &params.Id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PatchApiProxy(c, params)
}

// PostApiProxy operation middleware
func (siw *ServerInterfaceWrapper) PostApiProxy(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiProxy(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.DELETE(options.BaseURL+"/api/allow/:proxyId", wrapper.DeleteApiAllowProxyId)
	router.GET(options.BaseURL+"/api/allow/:proxyId", wrapper.GetApiAllowProxyId)
	router.POST(options.BaseURL+"/api/allow/:proxyId", wrapper.PostApiAllowProxyId)
	router.DELETE(options.BaseURL+"/api/block/:proxyId", wrapper.DeleteApiBlockProxyId)
	router.GET(options.BaseURL+"/api/block/:proxyId", wrapper.GetApiBlockProxyId)
	router.POST(options.BaseURL+"/api/block/:proxyId", wrapper.PostApiBlockProxyId)
	router.DELETE(options.BaseURL+"/api/guard-exclude/:guardId", wrapper.DeleteApiGuardExcludeGuardId)
	router.GET(options.BaseURL+"/api/guard-exclude/:guardId", wrapper.GetApiGuardExcludeGuardId)
	router.POST(options.BaseURL+"/api/guard-exclude/:guardId", wrapper.PostApiGuardExcludeGuardId)
	router.DELETE(options.BaseURL+"/api/guard/:proxyId", wrapper.DeleteApiGuardProxyId)
	router.GET(options.BaseURL+"/api/guard/:proxyId", wrapper.GetApiGuardProxyId)
	router.POST(options.BaseURL+"/api/guard/:proxyId", wrapper.PostApiGuardProxyId)
	router.DELETE(options.BaseURL+"/api/proxy", wrapper.DeleteApiProxy)
	router.GET(options.BaseURL+"/api/proxy", wrapper.GetApiProxy)
	router.PATCH(options.BaseURL+"/api/proxy", wrapper.PatchApiProxy)
	router.POST(options.BaseURL+"/api/proxy", wrapper.PostApiProxy)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+yZTW/jNhCG/4rA9qhdJ+3Nt6RdBIuirdGip0UOjDSx2UoiQ1JNDEP/veCQ8odI6sON",
	"sUqbmy2RM+T7zjyi5R3JeCl4BZVWZLkjKttASfHjTVHwZ/NBSC5AagZ4OZNANeQ32nx55LKkmixJTjV8",
	"0KwEkhK9FUCWRGnJqjVpUsJyM9a/LIKXheQv28+hKU2TktuCZ3/NblV3NZW5vyooKSvCWYKRUiLhqWYS",
	"crL84mbj2Ps2x6eXos7AzxTZjKB6M5yI5cQNNXl+geeI9UFpurFEGyTi1LQgk4UVVKlnLifIu59xnDKm",
	"8zhBj7VcmcLxAyleS5vA24Kmcg16OIsLsZ9gEq6ozjaRlBnNNscZHzgvgFYmZc4UfSiCNztJbZDDDEw6",
	"Od8Z7dqzxGgvv4LI2BuH5aae6oeVpW7HRpLf6ywDpX4DJXilAoVUglJ0jTdyUJlkQjNekSX52d5I+GOi",
	"N5DINkJK4IWWwihAfv3JV6iz7jb+PfKJVY8c98o0BrilSlOZJ0i1hApGUvI3SGWXcP3x6uOVEYkLqMzN",
	"JfkeL1lK4PoXVLAFNZxY7BwbG7uZAnRgWz/i9YQmOIdgbEnNTQNVd/tGMETPysHW5JO0BA1SkeWXbszP",
	"eatSG5SZy081yC1JSUVLs1d08CCNljWk7jkXtD+cRVAJlbaCtYmwzfd5xH7R45Pdm8HWYFT1u6sr7Bxe",
	"aaiwOqkQBctQqMWfyixodxTvWwmPZEm+WRye4gv3CF90SxAL4XRrKPahxMwAVZcllduQY5qujQvuYGCq",
	"3HXQadA70GaKnaY8p+9An2XzPA1gGko15IQVrNm3LJWSbif74enq+yG4ChjyA9IroUkFz5H2W3E1O1ee",
	"alD6lufbV+uI/cmmOaWlWU7jFcL1bDoxaGDX/Sa1TH4wx66JTMY5cSbjSW4yk9ug70yeVAko9iCTW3Hb",
	"KrCH7UEm47QYk6fb/IaZbAUbweRePzxdfT9GMDncfo7Js3LlIkx2VsyOyb3OBw3sut8yeW1+Un6Al6yo",
	"c1js8OsQm+2PigTHJm5qHNHuRyuOurPhZ0hq3Ey4JNf7Rb8ZUqPOUVJH/GtLxL7YGAR2v/2W22d7P09X",
	"RuH7+CXNCIj3WhXT2rcqzPJP1YhOdTSfo1cXgfqJQWPQPp/WjfjZLYcTuI88cB9DYQDm7+ftuVH8THqr",
	"Pmz/X47bVrBXJbU6E9G9aP5PH7SdCW+Xxj0UFu1/EIMvO1q7IvBdudsjqXti/r+j7lcGIW588MVDu9/W",
	"ByvXIAnNNAYxFLaaX55DNtMIDvWq4e/Kl0NQnW18Qf4QOe0pQ/z/7utX4evj5+iPyfkBqNdsz7GA18Pv",
	"eCJ220fPof4vwv0Jsl/PRvagel3pm6b5JwAA//8hHuiiwiIAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
