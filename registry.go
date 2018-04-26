package static

import "sync"

var (
	modulesMu sync.RWMutex
	modules   = make(map[string]Module)
)

// Register makes module available with provided alias.
func Register(alias string, module Module) {
	modulesMu.Lock()
	defer modulesMu.Unlock()

	modules[alias] = module
}

// Modules iterates all registered modules applying provided func.
func Modules(f func(alias string, module Module) bool) {
	modulesMu.Lock()
	defer modulesMu.Unlock()

	for alias, resource := range modules {
		if !f(alias, resource) {
			break
		}
	}
}
