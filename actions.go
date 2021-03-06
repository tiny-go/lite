package lite

import (
	"net/http"

	mw "github.com/tiny-go/middleware"
)

// options is responsible for handling OPTIONS request.
func options(methods *Methods) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", methods.Join())
	}
}

// getSingle handles single GET request on provided resource.
func getSingle(controller SingleGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.Get(r.Context(), ParamsFromContext(r.Context())["pk"])
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", mw.ResponseCodecFromContext(r.Context()).MimeType())
		// send the success response
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// getPlural handles plural GET request on provided resource.
func getPlural(controller PluralGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.GetAll(r.Context(), r.URL.Query())
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", mw.ResponseCodecFromContext(r.Context()).MimeType())
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// postSingle handles single POST request on provided resource.
func postSingle(controller SinglePoster) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.Post(r.Context(), func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", mw.ResponseCodecFromContext(r.Context()).MimeType())
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// postPlural handles bulk POST request on provided resource.
func postPlural(controller PluralPoster) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.PostAll(r.Context(), func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", mw.ResponseCodecFromContext(r.Context()).MimeType())
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// patchSingle handles single PATCH request on provided resource.
func patchSingle(controller SinglePatcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.Patch(r.Context(), ParamsFromContext(r.Context())["pk"], func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", mw.ResponseCodecFromContext(r.Context()).MimeType())
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// patchPlural handles bulk PATCH request on provided resource.
func patchPlural(controller PluralPatcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.PatchAll(r.Context(), r.URL.Query(), func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", mw.ResponseCodecFromContext(r.Context()).MimeType())
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// putSingle handles single PUT request on provided resource.
func putSingle(controller SinglePutter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.Put(r.Context(), ParamsFromContext(r.Context())["pk"], func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", mw.ResponseCodecFromContext(r.Context()).MimeType())
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// putPlural handles bulk PUT request for provided resource.
func putPlural(controller PluralPutter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.PutAll(r.Context(), r.URL.Query(), func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", mw.ResponseCodecFromContext(r.Context()).MimeType())
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// deleteSingle handles single DELETE request on provided resource.
func deleteSingle(controller SingleDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// delete model by primary key(s) TODO: primary is missing
		data, err := controller.Delete(r.Context(), ParamsFromContext(r.Context())["pk"])
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", mw.ResponseCodecFromContext(r.Context()).MimeType())
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// deletePlural handles bulk DELETE request on provided resource.
func deletePlural(controller PluralDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.DeleteAll(r.Context(), r.URL.Query())
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", mw.ResponseCodecFromContext(r.Context()).MimeType())
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}
