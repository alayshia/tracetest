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
	"time"
)

type DataStore struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name"`

	Type SupportedDataStores `json:"type"`

	Default bool `json:"default,omitempty"`

	Jaeger GrpcClientSettings `json:"jaeger,omitempty"`

	Tempo BaseClient `json:"tempo,omitempty"`

	Opensearch ElasticSearch `json:"opensearch,omitempty"`

	Elasticapm ElasticSearch `json:"elasticapm,omitempty"`

	Signalfx SignalFx `json:"signalfx,omitempty"`

	Awsxray AwsXRay `json:"awsxray,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// AssertDataStoreRequired checks if the required fields are not zero-ed
func AssertDataStoreRequired(obj DataStore) error {
	elements := map[string]interface{}{
		"name": obj.Name,
		"type": obj.Type,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertGrpcClientSettingsRequired(obj.Jaeger); err != nil {
		return err
	}
	if err := AssertBaseClientRequired(obj.Tempo); err != nil {
		return err
	}
	if err := AssertElasticSearchRequired(obj.Opensearch); err != nil {
		return err
	}
	if err := AssertElasticSearchRequired(obj.Elasticapm); err != nil {
		return err
	}
	if err := AssertSignalFxRequired(obj.Signalfx); err != nil {
		return err
	}
	if err := AssertAwsXRayRequired(obj.Awsxray); err != nil {
		return err
	}
	return nil
}

// AssertRecurseDataStoreRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DataStore (e.g. [][]DataStore), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDataStoreRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDataStore, ok := obj.(DataStore)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDataStoreRequired(aDataStore)
	})
}
