package restful

import "io"

// Encoder is responsible for data encoding.
type Encoder interface {
	Encode(v interface{}) error
}

// Decoder is responsible for data decoding.
type Decoder interface {
	Decode(v interface{}) error
}

// Codec can create Encoder(s) and Decoder(s) with provided io.Reader/io.Writer.
type Codec interface {
	// Ecoder instantiates the ecnoder part of this codec.
	Encoder(w io.Writer) Encoder
	// Decoder instantiates the decoder part of this codec.
	Decoder(r io.Reader) Decoder
	// MimeType returns the (main) mime type of this codec.
	MimeType() string
}
