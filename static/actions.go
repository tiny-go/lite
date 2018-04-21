package static

import (
	"net/http"
	"strings"

	mw "github.com/Alma-media/restful/middleware"

	"github.com/gorilla/mux"
)

type actionFunc func(w http.ResponseWriter, r *http.Request) error

// options is responsible for handling OPTIONS request.
func options(methods []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	}
}

// getPlural handles plural GET request on provided resource.
func getSingle(controller SingleGetter) actionFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		// call the controller action
		data, err := controller.Get(r.Context(), mux.Vars(r)["pk"])
		if err != nil {
			return err
		}
		// send the success response
		return mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data)
	}
}

// getPlural handles plural GET request on provided resource.
func getPlural(controller PluralGetter) actionFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		// call the controller action
		data, err := controller.GetAll(r.Context())
		if err != nil {
			return err
		}
		// send data to the client
		return mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data)
	}
}

// postSingle handles single model POST request for provided resource.
func postSingle(controller SinglePoster) actionFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		// call the controller action
		data, err := controller.Post(r.Context(), mux.Vars(r)["pk"], func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			return err
		}
		// send data to the client
		return mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data)
	}
}

// postPlural handles plural model POST request for provided resource.
func postPlural(controller PluralPoster) actionFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		// call the controller action
		data, err := controller.PostAll(r.Context(), func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			return err
		}
		// send data to the client
		return mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data)
	}
}

// patchSingle handles single model PATCH request for provided resource.
func patchSingle(controller SinglePatcher) actionFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		// call the controller action
		data, err := controller.Patch(r.Context(), mux.Vars(r)["pk"], func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			return err
		}
		// send data to the client
		return mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data)
	}
}

// putSingle handles single model PUT request for provided resource.
func putSingle(controller SinglePutter) actionFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		// call the controller action
		data, err := controller.Put(r.Context(), mux.Vars(r)["pk"], func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			return err
		}
		// send data to the client
		return mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data)
	}
}

// deleteSingle handles single DELETE model request for provided resource.
func deleteSingle(controller SingleDeleter) actionFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		// delete model by primary key(s) TODO: primary is missing
		data, err := controller.Delete(r.Context(), mux.Vars(r)["pk"])
		if err != nil {
			return err
		}
		// send data to the client
		return mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data)
	}
}

// deletePlural handles bulk DELETE request for provided resource.
func deletePlural(controller PluralDeleter) actionFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		// call the controller action
		data, err := controller.DeleteAll(r.Context(), func(v interface{}) error {
			return mw.RequestCodecFromContext(r.Context()).Decoder(r.Body).Decode(v)
		})
		if err != nil {
			return err
		}
		// send data to the client
		return mw.ResponseCodecFromContext(r.Context()).Encoder(w).Encode(data)
	}
}

func checkError(handler actionFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call handler
		panic(handler(w, r))
	}
}
