package static

import (
	"net/http"

	"github.com/tiny-go/middleware"

	"github.com/gorilla/mux"
)

// options is responsible for handling OPTIONS request.
func options(methods *methods) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", methods.join())
	}
}

// getSingle handles single GET request on provided resource.
func getSingle(controller SingleGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.Get(r.Context(), mux.Vars(r)["pk"])
		if err != nil {
			panic(err)
		}
		// send the success response
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// getPlural handles plural GET request on provided resource.
func getPlural(controller PluralGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.GetAll(r.Context())
		if err != nil {
			panic(err)
		}
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// postSingle handles single POST request on provided resource.
func postSingle(controller SinglePoster) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.Post(r.Context(), mux.Vars(r)["pk"], func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			panic(err)
		}
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
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// patchSingle handles single PATCH request on provided resource.
func patchSingle(controller SinglePatcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.Patch(r.Context(), mux.Vars(r)["pk"], func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			panic(err)
		}
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// patchPlural handles bulk PATCH request on provided resource.
func patchPlural(controller PluralPatcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.PatchAll(r.Context(), func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			panic(err)
		}
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// putSingle handles single PUT request on provided resource.
func putSingle(controller SinglePutter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.Put(r.Context(), mux.Vars(r)["pk"], func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			panic(err)
		}
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// putPlural handles bulk PUT request for provided resource.
func putPlural(controller PluralPutter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.PutAll(r.Context(), func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			panic(err)
		}
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// deleteSingle handles single DELETE request on provided resource.
func deleteSingle(controller SingleDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// delete model by primary key(s) TODO: primary is missing
		data, err := controller.Delete(r.Context(), mux.Vars(r)["pk"])
		if err != nil {
			panic(err)
		}
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}

// deletePlural handles bulk DELETE request on provided resource.
func deletePlural(controller PluralDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call the controller action
		data, err := controller.DeleteAll(r.Context(), func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			panic(err)
		}
		// send data to the client
		panic(mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data))
	}
}
