package static

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/tiny-go/errors"
	"github.com/tiny-go/lite"
	mw "github.com/tiny-go/middleware"
)

// FileController serves static content.
type FileController struct {
	*mw.BaseController
	assetsDir string
}

// NewModule creates a new handler for static content.
// TODO: add an extra parameter "mode" (for example 755) and maybe "group"
func NewModule(assetDir string) lite.Module {
	module := lite.NewBaseModule()
	module.Register("", &FileController{
		BaseController: mw.NewBaseController(),
		assetsDir:      assetDir,
	})
	return module
}

// Get returns file content.
func (c *FileController) Get(w http.ResponseWriter, r *http.Request) {
	target, ok := lite.ParamsFromContext(r.Context())["path"]
	if !ok {
		panic(errors.NewBadRequest(fmt.Errorf("file path has not been provided")))
	}
	info, err := os.Stat(path.Join(c.assetsDir, target))
	if err != nil {
		panic(errors.NewNotFound(fmt.Errorf("file %q not found", target)))
	}
	if info.IsDir() {
		panic(errors.NewForbidden(fmt.Errorf("access to directory is not allowed")))
	}
	file, err := os.Open(path.Join(c.assetsDir, target))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	http.ServeContent(w, r, target, time.Now(), file)
}

// Post uploads file.
func (c *FileController) Post(w http.ResponseWriter, r *http.Request) {
	target, ok := lite.ParamsFromContext(r.Context())["path"]
	if !ok {
		panic(errors.NewBadRequest(fmt.Errorf("file path has not been provided")))
	}
	if _, err := os.Stat(path.Join(c.assetsDir, target)); !os.IsNotExist(err) {
		panic(errors.NewConflict(fmt.Errorf("file %q already exists", target)))
	}
}
