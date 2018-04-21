package json

import (
	"encoding/json"
	"io"

	"github.com/Alma-media/restful"
	"github.com/Alma-media/restful/codecs"
)

const (
	// DataTypeJSON contains JSON codec data type.
	DataTypeJSON = "application/json"
)

// JSON codec.
type JSON struct{}

// Encoder creates JSON encoder.
func (j *JSON) Encoder(w io.Writer) restful.Encoder {
	return json.NewEncoder(w)
}

// Decoder creates JSON decoder.
func (j *JSON) Decoder(r io.Reader) restful.Decoder {
	return json.NewDecoder(r)
}

// MimeType returns the mime type of this codec (application/json).
func (j *JSON) MimeType() string {
	return DataTypeJSON
}

func init() {
	// since JSON is a stateless codec (no need to set boundary per request) we can
	// always return the same instance
	codec := &JSON{}
	// register as default codec
	codecs.Default(func(string) restful.Codec {
		return codec
	})
}
