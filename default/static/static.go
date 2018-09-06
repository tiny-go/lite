package static

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/tiny-go/lite"
	mw "github.com/tiny-go/middleware"
)

// FileController serves static content.
type FileController struct {
	*mw.BaseController
	assetsDir string
}

// NewModule creates a new handler for static content.
func NewModule(assetDir string) lite.Module {
	module := lite.NewBaseModule()
	module.Register("", &FileController{
		BaseController: mw.NewBaseController(),
		assetsDir:      assetDir,
	})
	return module
}

// Get is a classic http.HandlerFunc for GET requests.
func (c *FileController) Get(w http.ResponseWriter, r *http.Request) {
	target, ok := lite.ParamsFromContext(r.Context())["path"]
	if !ok {
		http.Error(w, "file path has not been provided", http.StatusBadRequest)
		return
	}
	file, err := os.Open(path.Join(c.assetsDir, target))
	if err != nil {
		http.Error(w, fmt.Sprintf("file %q not found", target), http.StatusNotFound)
		return
	}
	defer file.Close()
	http.ServeContent(w, r, target, time.Now(), file)
}
