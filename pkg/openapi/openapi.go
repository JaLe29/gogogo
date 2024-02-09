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

// Activity defines model for Activity.
type Activity struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        *string    `json:"id,omitempty"`
	Ip        *string    `json:"ip,omitempty"`
	ProxyId   *string    `json:"proxyId,omitempty"`
}

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

// NewAllow defines model for NewAllow.
type NewAllow struct {
	Ip string `json:"ip"`
}

// NewBlock defines model for NewBlock.
type NewBlock struct {
	Ip string `json:"ip"`
}

// NewProxy defines model for NewProxy.
type NewProxy struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

// PatchProxy defines model for PatchProxy.
type PatchProxy struct {
	Disable *bool `json:"disable,omitempty"`
}

// Proxy defines model for Proxy.
type Proxy struct {
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

// PatchApiProxyJSONRequestBody defines body for PatchApiProxy for application/json ContentType.
type PatchApiProxyJSONRequestBody = PatchProxy

// PostApiProxyJSONRequestBody defines body for PostApiProxy for application/json ContentType.
type PostApiProxyJSONRequestBody = NewProxy

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all activities
	// (GET /api/activity/{proxyId})
	GetApiActivityProxyId(c *gin.Context, proxyId string)
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

// GetApiActivityProxyId operation middleware
func (siw *ServerInterfaceWrapper) GetApiActivityProxyId(c *gin.Context) {

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

	siw.Handler.GetApiActivityProxyId(c, proxyId)
}

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

	router.GET(options.BaseURL+"/api/activity/:proxyId", wrapper.GetApiActivityProxyId)
	router.DELETE(options.BaseURL+"/api/allow/:proxyId", wrapper.DeleteApiAllowProxyId)
	router.GET(options.BaseURL+"/api/allow/:proxyId", wrapper.GetApiAllowProxyId)
	router.POST(options.BaseURL+"/api/allow/:proxyId", wrapper.PostApiAllowProxyId)
	router.DELETE(options.BaseURL+"/api/block/:proxyId", wrapper.DeleteApiBlockProxyId)
	router.GET(options.BaseURL+"/api/block/:proxyId", wrapper.GetApiBlockProxyId)
	router.POST(options.BaseURL+"/api/block/:proxyId", wrapper.PostApiBlockProxyId)
	router.DELETE(options.BaseURL+"/api/proxy", wrapper.DeleteApiProxy)
	router.GET(options.BaseURL+"/api/proxy", wrapper.GetApiProxy)
	router.PATCH(options.BaseURL+"/api/proxy", wrapper.PatchApiProxy)
	router.POST(options.BaseURL+"/api/proxy", wrapper.PostApiProxy)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xYT2/bPgz9KgZ/v6PXpNvNt3QDhmLYFmzYqciBtZlEm22pktI2CPzdB8p2/sl24q1F",
	"XWC3xBIp8r3HpzgbiGWmZE65NRBtwMRLytB9nMRW3Au75s9KS0XaCnIrsSa0lEwsf5lLnaGFCBK09MaK",
	"jCAEu1YEERirRb6AIgSR8F7/sWp8rLR8XF83hRRFCJM0lQ+Dq+oqlfGvwVX1hR5a4GpMV4Sg6W4lNCUQ",
	"3fCeWZmkpbt+SaZcqp/EyJWOqbE5i3pB9vQZVYptAB84RRsvW45MhMHbdP/MWylTwrxErSXqD7jsOKiV",
	"6CfAQyQQ7pUbegDtKmOovq/imIz5RkbJ3JDfeUbG4MItJGRiLZQVMocIPpcLgZwHdkmBrjOEQI+YKe4c",
	"vn7ykTmqt84/c/CLfC5dj8K6BFdoLOokcFIPUAkI4Z60KUu4vBhfjBkcqSjnxQjeuUchKLRLV/8IlRhh",
	"ZWijTTUzBS9VeB629ZFsgGkaVCGMgsuvkTfwtPGWiRK1SU6rKeQzNWZkSRuIbo7zXieMlEJNuS3bAW4X",
	"IlcqhJBjxh2rbbodSlavKKwMukkBM95cwu96fjseO9HK3FLuekSlUhG7FkY/DRe02csnLGUu8H9Nc4jg",
	"v9HuahhV98JoeykUW0pRa1yXc3PYbL13pwreY1ZZhnrdBrLFBQO3u31mHFTyx0Z2SF5CKdkGWX5wzwMM",
	"XIzHXbnM9PFyL+5Y5XVSR9zdivR6x5zoR1r4mhTSJYxjC2nSA8PWJgaPsa0Q3Hd2qe5J5W2tU9qb5tc8",
	"og6wc+aziw8PV58PJU0DIe/drRNgkNNDy/hNpRkcK3crMvZKJusnm4jtT6/i8LbjcgpPCJeDmcRGAo/Z",
	"rz35ln8X9vRkF9Puye6nZm9PrpP+8+ReSnBgn/TkGtxaBeXbwElPdmFtntyf5lfsySVgZ3hyJx8erj4f",
	"Z3hy8/hVnjwoVp7FkysqBufJncw3EnjMfu3Jqn57PenENWktTjytls+04AMJ/J0Fv7ArusZPumLdb81D",
	"CddJV+Sw9hfKGvPn96TypDM8qRMNvysfDoU2XvqA/FAJdsjQ/X/z8ip8ehPa+2PqLBsajvQ9xhq4Pn0B",
	"tdBdXkA7/T+L+/eA/XIwsDeidwx9URS/AwAA//9cmKTXzRcAAA==",
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
