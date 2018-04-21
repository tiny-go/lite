package restful

// Module describes a single application module.
type Module interface {
	Name() string
	// Register should add Controller to module resources.
	Register(string, Controller) error
	// Unregister should remove Storage from module resources.
	Unregister(alias string) error
	// Lookup should find a Storage by alias in module resource list.
	Resources() map[string]Controller
	// Init should initialize a module with provided dependencies.
	Init() error
}
