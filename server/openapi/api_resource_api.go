/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// ResourceApiApiController binds http requests to an api service and writes the service results to the http response
type ResourceApiApiController struct {
	service      ResourceApiApiServicer
	errorHandler ErrorHandler
}

// ResourceApiApiOption for how the controller is set up.
type ResourceApiApiOption func(*ResourceApiApiController)

// WithResourceApiApiErrorHandler inject ErrorHandler into controller
func WithResourceApiApiErrorHandler(h ErrorHandler) ResourceApiApiOption {
	return func(c *ResourceApiApiController) {
		c.errorHandler = h
	}
}

// NewResourceApiApiController creates a default api controller
func NewResourceApiApiController(s ResourceApiApiServicer, opts ...ResourceApiApiOption) Router {
	controller := &ResourceApiApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ResourceApiApiController
func (c *ResourceApiApiController) Routes() Routes {
	return Routes{
		{
			"CreateDemo",
			strings.ToUpper("Post"),
			"/api/demos",
			c.CreateDemo,
		},
		{
			"CreateEnvironment",
			strings.ToUpper("Post"),
			"/api/environments",
			c.CreateEnvironment,
		},
		{
			"DeleteDataStore",
			strings.ToUpper("Delete"),
			"/api/datastores/{dataStoreId}",
			c.DeleteDataStore,
		},
		{
			"DeleteDemo",
			strings.ToUpper("Delete"),
			"/api/demos/{demoId}",
			c.DeleteDemo,
		},
		{
			"DeleteEnvironment",
			strings.ToUpper("Delete"),
			"/api/environments/{environmentId}",
			c.DeleteEnvironment,
		},
		{
			"GetConfiguration",
			strings.ToUpper("Get"),
			"/api/configs/{configId}",
			c.GetConfiguration,
		},
		{
			"GetDataStore",
			strings.ToUpper("Get"),
			"/api/datastores/{dataStoreId}",
			c.GetDataStore,
		},
		{
			"GetDemo",
			strings.ToUpper("Get"),
			"/api/demos/{demoId}",
			c.GetDemo,
		},
		{
			"GetEnvironment",
			strings.ToUpper("Get"),
			"/api/environments/{environmentId}",
			c.GetEnvironment,
		},
		{
			"GetPollingProfile",
			strings.ToUpper("Get"),
			"/api/pollingprofiles/{pollingProfileId}",
			c.GetPollingProfile,
		},
		{
			"ListConfiguration",
			strings.ToUpper("Get"),
			"/api/configs",
			c.ListConfiguration,
		},
		{
			"ListDataStore",
			strings.ToUpper("Get"),
			"/api/datastores",
			c.ListDataStore,
		},
		{
			"ListDemos",
			strings.ToUpper("Get"),
			"/api/demos",
			c.ListDemos,
		},
		{
			"ListEnvironments",
			strings.ToUpper("Get"),
			"/api/environments",
			c.ListEnvironments,
		},
		{
			"ListPollingProfile",
			strings.ToUpper("Get"),
			"/api/pollingprofiles",
			c.ListPollingProfile,
		},
		{
			"UpdateConfiguration",
			strings.ToUpper("Put"),
			"/api/configs/{configId}",
			c.UpdateConfiguration,
		},
		{
			"UpdateDataStore",
			strings.ToUpper("Put"),
			"/api/datastores/{dataStoreId}",
			c.UpdateDataStore,
		},
		{
			"UpdateDemo",
			strings.ToUpper("Put"),
			"/api/demos/{demoId}",
			c.UpdateDemo,
		},
		{
			"UpdateEnvironment",
			strings.ToUpper("Put"),
			"/api/environments/{environmentId}",
			c.UpdateEnvironment,
		},
		{
			"UpdatePollingProfile",
			strings.ToUpper("Put"),
			"/api/pollingprofiles/{pollingProfileId}",
			c.UpdatePollingProfile,
		},
	}
}

// CreateDemo - Create a Demonstration setting
func (c *ResourceApiApiController) CreateDemo(w http.ResponseWriter, r *http.Request) {
	demoParam := Demo{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&demoParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDemoRequired(demoParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateDemo(r.Context(), demoParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// CreateEnvironment - Create an environment
func (c *ResourceApiApiController) CreateEnvironment(w http.ResponseWriter, r *http.Request) {
	environmentResourceParam := EnvironmentResource{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&environmentResourceParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEnvironmentResourceRequired(environmentResourceParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateEnvironment(r.Context(), environmentResourceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteDataStore - Delete a Data Store
func (c *ResourceApiApiController) DeleteDataStore(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dataStoreIdParam := params["dataStoreId"]

	result, err := c.service.DeleteDataStore(r.Context(), dataStoreIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteDemo - Delete a Demonstration setting
func (c *ResourceApiApiController) DeleteDemo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	demoIdParam := params["demoId"]

	result, err := c.service.DeleteDemo(r.Context(), demoIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteEnvironment - Delete an environment
func (c *ResourceApiApiController) DeleteEnvironment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	environmentIdParam := params["environmentId"]

	result, err := c.service.DeleteEnvironment(r.Context(), environmentIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetConfiguration - Get Tracetest configuration
func (c *ResourceApiApiController) GetConfiguration(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	configIdParam := params["configId"]

	result, err := c.service.GetConfiguration(r.Context(), configIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetDataStore - Get a Data Store
func (c *ResourceApiApiController) GetDataStore(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dataStoreIdParam := params["dataStoreId"]

	result, err := c.service.GetDataStore(r.Context(), dataStoreIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetDemo - Get Demonstration setting
func (c *ResourceApiApiController) GetDemo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	demoIdParam := params["demoId"]

	result, err := c.service.GetDemo(r.Context(), demoIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetEnvironment - Get a specific environment
func (c *ResourceApiApiController) GetEnvironment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	environmentIdParam := params["environmentId"]

	result, err := c.service.GetEnvironment(r.Context(), environmentIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetPollingProfile - Get Polling Profile
func (c *ResourceApiApiController) GetPollingProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pollingProfileIdParam := params["pollingProfileId"]

	result, err := c.service.GetPollingProfile(r.Context(), pollingProfileIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ListConfiguration - List Tracetest configuration
func (c *ResourceApiApiController) ListConfiguration(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	takeParam, err := parseInt32Parameter(query.Get("take"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	skipParam, err := parseInt32Parameter(query.Get("skip"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	sortByParam := query.Get("sortBy")
	sortDirectionParam := query.Get("sortDirection")
	result, err := c.service.ListConfiguration(r.Context(), takeParam, skipParam, sortByParam, sortDirectionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ListDataStore - List Data Store
func (c *ResourceApiApiController) ListDataStore(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	takeParam, err := parseInt32Parameter(query.Get("take"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	skipParam, err := parseInt32Parameter(query.Get("skip"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	sortByParam := query.Get("sortBy")
	sortDirectionParam := query.Get("sortDirection")
	result, err := c.service.ListDataStore(r.Context(), takeParam, skipParam, sortByParam, sortDirectionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ListDemos - List Demonstrations
func (c *ResourceApiApiController) ListDemos(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	takeParam, err := parseInt32Parameter(query.Get("take"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	skipParam, err := parseInt32Parameter(query.Get("skip"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	sortByParam := query.Get("sortBy")
	sortDirectionParam := query.Get("sortDirection")
	result, err := c.service.ListDemos(r.Context(), takeParam, skipParam, sortByParam, sortDirectionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ListEnvironments - List environments
func (c *ResourceApiApiController) ListEnvironments(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	takeParam, err := parseInt32Parameter(query.Get("take"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	skipParam, err := parseInt32Parameter(query.Get("skip"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	sortByParam := query.Get("sortBy")
	sortDirectionParam := query.Get("sortDirection")
	result, err := c.service.ListEnvironments(r.Context(), takeParam, skipParam, sortByParam, sortDirectionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ListPollingProfile - List Polling Profile Configuration
func (c *ResourceApiApiController) ListPollingProfile(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	takeParam, err := parseInt32Parameter(query.Get("take"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	skipParam, err := parseInt32Parameter(query.Get("skip"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	sortByParam := query.Get("sortBy")
	sortDirectionParam := query.Get("sortDirection")
	result, err := c.service.ListPollingProfile(r.Context(), takeParam, skipParam, sortByParam, sortDirectionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateConfiguration - Update Tracetest configuration
func (c *ResourceApiApiController) UpdateConfiguration(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	configIdParam := params["configId"]

	configurationResourceParam := ConfigurationResource{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&configurationResourceParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertConfigurationResourceRequired(configurationResourceParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateConfiguration(r.Context(), configIdParam, configurationResourceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateDataStore - Update a Data Store
func (c *ResourceApiApiController) UpdateDataStore(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dataStoreIdParam := params["dataStoreId"]

	dataStoreParam := DataStore{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&dataStoreParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDataStoreRequired(dataStoreParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateDataStore(r.Context(), dataStoreIdParam, dataStoreParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateDemo - Update a Demonstration setting
func (c *ResourceApiApiController) UpdateDemo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	demoIdParam := params["demoId"]

	demoParam := Demo{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&demoParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDemoRequired(demoParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateDemo(r.Context(), demoIdParam, demoParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateEnvironment - Update an environment
func (c *ResourceApiApiController) UpdateEnvironment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	environmentIdParam := params["environmentId"]

	environmentResourceParam := EnvironmentResource{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&environmentResourceParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEnvironmentResourceRequired(environmentResourceParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateEnvironment(r.Context(), environmentIdParam, environmentResourceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdatePollingProfile - Update a Polling Profile
func (c *ResourceApiApiController) UpdatePollingProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pollingProfileIdParam := params["pollingProfileId"]

	pollingProfileParam := PollingProfile{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&pollingProfileParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPollingProfileRequired(pollingProfileParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdatePollingProfile(r.Context(), pollingProfileIdParam, pollingProfileParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
