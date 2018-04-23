package codecs

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Alma-media/restful"
)

var (
	codecs   registry
	codecsMu sync.RWMutex
)

// Registry represents codec registry/list.
type Registry interface {
	Get(mime string) restful.Codec
	// WithDefault()
}

// Global returns global codec registry.
func Global() Registry {
	// return an interface (it is safe / read only)
	return codecs
}

type codecInitFunc func(string) restful.Codec

type registry map[string]codecInitFunc

// Get appropriate codec by provided MimeType, if there is no exact match (that
// is possible using codecs such as Multipart), it tries to detect a codec in an
// opposite way, comparing default codec MimeType to the provided argument.
func (r registry) Get(mime string) restful.Codec {
	codecsMu.Lock()
	defer codecsMu.Unlock()
	// strict match
	if codec, ok := r[mime]; ok {
		return codec(mime)
	}
	// submatch (for example "multipart/form-data; boundary=-----...")
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

// Default allows to set codec as adefault by mime type.
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

// Register makes codec available for provided MimeType.
func Register(mime string, f codecInitFunc) {
	codecsMu.Lock()
	defer codecsMu.Unlock()

	if _, ok := codecs[mime]; ok {
		panic(fmt.Sprintf("codec with MimeType %q already registered", mime))
	}
	codecs[mime] = f
}
