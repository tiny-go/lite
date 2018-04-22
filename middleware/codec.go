package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Alma-media/restful"
	"github.com/Alma-media/restful/codecs"
	"github.com/Alma-media/restful/errors"
)

// TODO:
// - provide a list of available codecs per action/method with an ability to
// set default
// - provide error handler func (to decide whether panic or send an error directly)
// - provide several constructors to instantiate middleware with panic/http/etc
// error handlers
// - implement smart/independent codec registry

type codecKey struct{ kind string }

// Codec middleware puts the correct codecs into the context.
func Codec(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reqCodec, resCodec restful.Codec
		// get request codec
		if reqCodec = codecs.Get(r.Header.Get("Content-Type")); reqCodec == nil {
			panic(errors.NewBadRequest(fmt.Sprintf("unsupported codec: %q", r.Header.Get("Content-Type"))))
		}
		r = r.WithContext(context.WithValue(r.Context(), codecKey{"req"}, reqCodec))
		// get response codec
		if resCodec = codecs.Get(r.Header.Get("Accept")); reqCodec == nil {
			panic(errors.NewBadRequest(fmt.Sprintf("unsupported codec: %q", r.Header.Get("Accept"))))
		}
		r = r.WithContext(context.WithValue(r.Context(), codecKey{"res"}, resCodec))
		// call the next handler
		next.ServeHTTP(w, r)
	})
}

// RequestCodecFromContext pulls the Codec from a request context or or returns nil.
func RequestCodecFromContext(ctx context.Context) restful.Codec {
	codec, _ := ctx.Value(codecKey{"req"}).(restful.Codec)
	return codec
}

// ResponseCodecFromContext pulls the Codec from a request context or returns nil.
func ResponseCodecFromContext(ctx context.Context) restful.Codec {
	codec, _ := ctx.Value(codecKey{"res"}).(restful.Codec)
	return codec
}
