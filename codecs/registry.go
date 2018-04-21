package codecs

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Alma-media/restful"
)

var (
	codecs   map[string]codecInitFunc
	codecsMu sync.RWMutex
)

type codecInitFunc func(string) restful.Codec

// Default allows to register provided codec as a default (will be used in case
// Content-Type/Accept are not provided) plus makes the codec available by mime.
func Default(mime string, f codecInitFunc) {
	codecsMu.Lock()
	defer codecsMu.Unlock()

	if _, ok := codecs[""]; ok {
		panic(fmt.Sprintf("cannot register %q as a default codec: default already registered", mime))
	}
	codecs[""] = f
	codecs["*/*"] = f
	codecs[mime] = f
}

// Register makes codec available for provided MimeType.
func Register(mime string, f codecInitFunc) {
	codecsMu.Lock()
	defer codecsMu.Unlock()

	if _, ok := codecs[mime]; ok {
		panic(fmt.Sprintf("codec with MimeType %q already registered", mime))
	}
	codecs[mime] = f
}

// Get appropriate codec by provided MimeType, if there is no exact match (that
// is possible using codecs such as Multipart), it tries to detect a codec in an
// opposite way, comparing default codec MimeType to the provided argument.
func Get(mime string) restful.Codec {
	codecsMu.Lock()
	defer codecsMu.Unlock()
	// strict match
	if codec, ok := codecs[mime]; ok {
		return codec(mime)
	}
	// submatch (for example "multipart/form-data; boundary=-----...")
	for mediatype, codec := range codecs {
		// ignore default codec
		if mediatype != "" && mediatype != "*/*" {
			if strings.Contains(mime, mediatype+";") {
				return codec(mime)
			}
		}
	}
	return nil
}
