package codecs

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Alma-media/restful"
)

var (
	codecs   globalRegistry
	codecsMu sync.RWMutex
)

// Global returns global codec registry.
func Global() Registry {
	// return an interface (it is safe / read only)
	return codecs
}

// Register makes codec available for provided mime type.
func Register(mime string, f codecInitFunc) {
	codecsMu.Lock()
	defer codecsMu.Unlock()

	if _, ok := codecs[mime]; ok {
		panic(fmt.Sprintf("codec with MimeType %q already registered", mime))
	}
	codecs[mime] = f
}

// Default allows setting Codec as a default by mime type. If codec was not found
// function returns an error. This function can be called multiple times overriding
// the previous values.
func Default(mime string) error {
	codecsMu.Lock()
	defer codecsMu.Unlock()

	f, ok := codecs[mime]
	if !ok {
		return fmt.Errorf("codec %q is not registered", mime)
	}
	codecs[""] = f
	codecs["*/*"] = f
	return nil
}

// Registry represents codec registry/list.
type Registry interface {
	Get(mime string) restful.Codec
}

// codecInitFunc is a codec constructor which can either return a singleton instance
// or initialize a new codec for every request (for example multipart, because each
// request has its own unique boundary).
type codecInitFunc func(string) restful.Codec

// globalRegistry is a clobal codec registry that is used to store available codec
// and retrieve them by mime type.
type globalRegistry map[string]codecInitFunc

// Get appropriate codec by provided MimeType, if there is no exact match (that
// is possible using codecs such as Multipart), it tries to detect a codec in an
// opposite way, comparing default codec MimeType to the provided argument.
func (r globalRegistry) Get(mime string) restful.Codec {
	codecsMu.Lock()
	defer codecsMu.Unlock()
	// find strict match
	if codec, ok := r[mime]; ok {
		return codec(mime)
	}
	// find submatch (for example "multipart/form-data; boundary=-----...")
	for mediatype, codec := range r {
		// ignore default codec
		if mediatype != "" && mediatype != "*/*" {
			if strings.Contains(mime, mediatype+";") {
				return codec(mime)
			}
		}
	}
	return nil
}
