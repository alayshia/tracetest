package resourcemanager

import (
	"encoding/json"
	"net/http"

	"github.com/goccy/go-yaml"
)

type encoder interface {
	Marshal(in interface{}) (out []byte, err error)
	Unmarshal(in []byte, out interface{}) (err error)
	Accepts(contentType string) bool
	ResponseContentType() string
}

var encoders = []encoder{
	jsonEncoder,
	yamlEncoder,
}

var defaultEncoder = jsonEncoder

var jsonEncoder = basicEncoder{
	contentType: "application/json",
	marshalFn:   json.Marshal,
	unmarshalFn: json.Unmarshal,
}

var yamlEncoder = basicEncoder{
	contentType: "text/yaml",
	marshalFn:   yaml.Marshal,
	unmarshalFn: yaml.Unmarshal,
}

type basicEncoder struct {
	contentType string
	marshalFn   func(interface{}) ([]byte, error)
	unmarshalFn func([]byte, interface{}) error
}

func (e basicEncoder) ResponseContentType() string {
	return e.contentType
}

func (e basicEncoder) Accepts(contentType string) bool {
	return contentType == e.contentType
}

func (e basicEncoder) Marshal(in interface{}) (out []byte, err error) {
	return e.marshalFn(in)
}

func (e basicEncoder) Unmarshal(in []byte, out interface{}) (err error) {
	return e.unmarshalFn(in, out)
}

func getInputEncoder(r *http.Request) encoder {
	return getEncoderForHeader(r, "Content-Type", "Accept")
}

func getOutputEncoder(r *http.Request) encoder {
	return getEncoderForHeader(r, "Accept", "Content-Type")
}

func getEncoderForHeader(r *http.Request, headerWaterfall ...string) encoder {
	for _, header := range headerWaterfall {
		headerValue := r.Header.Get(header)
		if headerValue == "" {
			continue
		}

		for _, enc := range encoders {
			if enc.Accepts(headerValue) {
				return enc
			}
		}
	}

	return defaultEncoder
}
