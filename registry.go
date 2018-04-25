package static

var (
	modules = make(map[string]Module)
)

// Register makes module available with provided alias.
func Register(alias string, module Module) {
	modules[alias] = module
}

// Modules iterates all registered modules applying provided func.
func Modules(f func(string, Module) bool) {
	for alias, resource := range modules {
		if !f(alias, resource) {
			break
		}
	}
}
