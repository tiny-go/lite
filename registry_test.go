package lite

import "testing"

func Test_Register(t *testing.T) {
	t.Run("Given global module registry", func(t *testing.T) {
		t.Run("Register func should add module to the list if its alias is unique", func(t *testing.T) {
			defer func() {
				modules = make(map[string]Module)
				if r := recover(); r != nil {
					t.Error("should not panic")
				}
			}()
			moduleA := NewBaseModule()
			Register("moduleA", moduleA)
			if module, ok := modules["moduleA"]; !ok || module != moduleA {
				t.Error("module has not been registered")
			}
		})
		t.Run("Register func should panic registering module if its alias is not unique", func(t *testing.T) {
			defer func() {
				modules = make(map[string]Module)
				if r := recover(); r == nil {
					t.Error("should have panicked")
				}
			}()
			Register("moduleA", NewBaseModule())
			Register("moduleA", NewBaseModule())
		})
	})
}

func Test_Modules(t *testing.T) {
	type module struct {
		*BaseModule
		name string
	}
	moduleA := &module{name: "A"}
	moduleB := &module{name: "B"}
	moduleC := &module{name: "C"}

	t.Run("Given global module registry", func(t *testing.T) {
		defer func() { modules = make(map[string]Module) }()
		Register("moduleA", moduleA)
		Register("moduleB", moduleB)
		Register("moduleC", moduleC)
		t.Run("Modules should iterate all registered modules applying provided func", func(t *testing.T) {
			var out string
			Modules(func(_ string, m Module) bool { out += m.(*module).name; return true })
			if len(out) != 3 {
				t.Errorf("out was expected to contain all (3) elements but had %d", len(out))
			}
		})
		t.Run("Modules should stop iterations if provided func returns false", func(t *testing.T) {
			var out string
			Modules(func(_ string, m Module) bool { out += m.(*module).name; return false })
			if len(out) != 1 {
				t.Errorf("out was expected to contain only one element but had %d", len(out))
			}
		})
	})
}
