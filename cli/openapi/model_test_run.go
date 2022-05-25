/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// TestRun struct for TestRun
type TestRun struct {
	Id      *string `json:"id,omitempty"`
	TraceId *string `json:"traceId,omitempty"`
	SpanId  *string `json:"spanId,omitempty"`
	// Test version used when running this test run
	TestVersion *int32 `json:"testVersion,omitempty"`
	// Current execution state
	State *string `json:"state,omitempty"`
	// Details of the cause for the last `FAILED` state
	LastErrorState *string           `json:"lastErrorState,omitempty"`
	CreatedAt      *time.Time        `json:"createdAt,omitempty"`
	CompletedAt    *time.Time        `json:"completedAt,omitempty"`
	Request        *HTTPRequest      `json:"request,omitempty"`
	Response       *HTTPResponse     `json:"response,omitempty"`
	Trace          *Trace            `json:"trace,omitempty"`
	Result         *AssertionResults `json:"result,omitempty"`
}

// NewTestRun instantiates a new TestRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRun() *TestRun {
	this := TestRun{}
	return &this
}

// NewTestRunWithDefaults instantiates a new TestRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRunWithDefaults() *TestRun {
	this := TestRun{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TestRun) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRun) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TestRun) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TestRun) SetId(v string) {
	o.Id = &v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *TestRun) GetTraceId() string {
	if o == nil || o.TraceId == nil {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRun) GetTraceIdOk() (*string, bool) {
	if o == nil || o.TraceId == nil {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *TestRun) HasTraceId() bool {
	if o != nil && o.TraceId != nil {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *TestRun) SetTraceId(v string) {
	o.TraceId = &v
}

// GetSpanId returns the SpanId field value if set, zero value otherwise.
func (o *TestRun) GetSpanId() string {
	if o == nil || o.SpanId == nil {
		var ret string
		return ret
	}
	return *o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRun) GetSpanIdOk() (*string, bool) {
	if o == nil || o.SpanId == nil {
		return nil, false
	}
	return o.SpanId, true
}

// HasSpanId returns a boolean if a field has been set.
func (o *TestRun) HasSpanId() bool {
	if o != nil && o.SpanId != nil {
		return true
	}

	return false
}

// SetSpanId gets a reference to the given string and assigns it to the SpanId field.
func (o *TestRun) SetSpanId(v string) {
	o.SpanId = &v
}

// GetTestVersion returns the TestVersion field value if set, zero value otherwise.
func (o *TestRun) GetTestVersion() int32 {
	if o == nil || o.TestVersion == nil {
		var ret int32
		return ret
	}
	return *o.TestVersion
}

// GetTestVersionOk returns a tuple with the TestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRun) GetTestVersionOk() (*int32, bool) {
	if o == nil || o.TestVersion == nil {
		return nil, false
	}
	return o.TestVersion, true
}

// HasTestVersion returns a boolean if a field has been set.
func (o *TestRun) HasTestVersion() bool {
	if o != nil && o.TestVersion != nil {
		return true
	}

	return false
}

// SetTestVersion gets a reference to the given int32 and assigns it to the TestVersion field.
func (o *TestRun) SetTestVersion(v int32) {
	o.TestVersion = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TestRun) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRun) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TestRun) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *TestRun) SetState(v string) {
	o.State = &v
}

// GetLastErrorState returns the LastErrorState field value if set, zero value otherwise.
func (o *TestRun) GetLastErrorState() string {
	if o == nil || o.LastErrorState == nil {
		var ret string
		return ret
	}
	return *o.LastErrorState
}

// GetLastErrorStateOk returns a tuple with the LastErrorState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRun) GetLastErrorStateOk() (*string, bool) {
	if o == nil || o.LastErrorState == nil {
		return nil, false
	}
	return o.LastErrorState, true
}

// HasLastErrorState returns a boolean if a field has been set.
func (o *TestRun) HasLastErrorState() bool {
	if o != nil && o.LastErrorState != nil {
		return true
	}

	return false
}

// SetLastErrorState gets a reference to the given string and assigns it to the LastErrorState field.
func (o *TestRun) SetLastErrorState(v string) {
	o.LastErrorState = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TestRun) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRun) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TestRun) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TestRun) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *TestRun) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRun) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil || o.CompletedAt == nil {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *TestRun) HasCompletedAt() bool {
	if o != nil && o.CompletedAt != nil {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given time.Time and assigns it to the CompletedAt field.
func (o *TestRun) SetCompletedAt(v time.Time) {
	o.CompletedAt = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *TestRun) GetRequest() HTTPRequest {
	if o == nil || o.Request == nil {
		var ret HTTPRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRun) GetRequestOk() (*HTTPRequest, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *TestRun) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given HTTPRequest and assigns it to the Request field.
func (o *TestRun) SetRequest(v HTTPRequest) {
	o.Request = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *TestRun) GetResponse() HTTPResponse {
	if o == nil || o.Response == nil {
		var ret HTTPResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRun) GetResponseOk() (*HTTPResponse, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *TestRun) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given HTTPResponse and assigns it to the Response field.
func (o *TestRun) SetResponse(v HTTPResponse) {
	o.Response = &v
}

// GetTrace returns the Trace field value if set, zero value otherwise.
func (o *TestRun) GetTrace() Trace {
	if o == nil || o.Trace == nil {
		var ret Trace
		return ret
	}
	return *o.Trace
}

// GetTraceOk returns a tuple with the Trace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRun) GetTraceOk() (*Trace, bool) {
	if o == nil || o.Trace == nil {
		return nil, false
	}
	return o.Trace, true
}

// HasTrace returns a boolean if a field has been set.
func (o *TestRun) HasTrace() bool {
	if o != nil && o.Trace != nil {
		return true
	}

	return false
}

// SetTrace gets a reference to the given Trace and assigns it to the Trace field.
func (o *TestRun) SetTrace(v Trace) {
	o.Trace = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *TestRun) GetResult() AssertionResults {
	if o == nil || o.Result == nil {
		var ret AssertionResults
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRun) GetResultOk() (*AssertionResults, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *TestRun) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given AssertionResults and assigns it to the Result field.
func (o *TestRun) SetResult(v AssertionResults) {
	o.Result = &v
}

func (o TestRun) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TraceId != nil {
		toSerialize["traceId"] = o.TraceId
	}
	if o.SpanId != nil {
		toSerialize["spanId"] = o.SpanId
	}
	if o.TestVersion != nil {
		toSerialize["testVersion"] = o.TestVersion
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.LastErrorState != nil {
		toSerialize["lastErrorState"] = o.LastErrorState
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CompletedAt != nil {
		toSerialize["completedAt"] = o.CompletedAt
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if o.Trace != nil {
		toSerialize["trace"] = o.Trace
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableTestRun struct {
	value *TestRun
	isSet bool
}

func (v NullableTestRun) Get() *TestRun {
	return v.value
}

func (v *NullableTestRun) Set(val *TestRun) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRun) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRun(val *TestRun) *NullableTestRun {
	return &NullableTestRun{value: val, isSet: true}
}

func (v NullableTestRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}