package exec

import (
	"context"
	"fmt"
	"io"
	"os/exec"

	"github.com/tiny-go/errors"
	"github.com/tiny-go/lite"
	"github.com/tiny-go/lite/examples/os/config"
	mw "github.com/tiny-go/middleware"
)

var _ lite.SingleGetter = &Controller{}

// Controller is responsible for user AUTH operations.
type Controller struct {
	// inherit BaseController functionality (related to middleware)
	*mw.BaseController
	// controller dependencies
	Config   *config.Config `inject:"t"`
	Commands Commands       `inject:"t"`
}

// Init user controller (TODO: add middleware for available methods).
func (c *Controller) Init() error { return nil }

// Get handles os command request.
func (c *Controller) Get(_ context.Context, command string) (interface{}, error) {
	if !c.Config.AllowAll && !c.Commands.Lookup(command) {
		return nil, errors.NewForbidden(fmt.Errorf("command %q is not allowed", command))
	}

	cmd := exec.Command("sh", "-c", command) // "sh", "-c", "cd .. && ls -la"
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return err, nil
	}

	return out, nil
}
