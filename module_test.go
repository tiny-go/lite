package lite

import "testing"

func Test_BaseModule(t *testing.T) {
	t.Run("Given BaseModule", func(t *testing.T) {
		module := NewBaseModule()
		t.Run("test if BaseModule is properly initialized with constructor func", func(t *testing.T) {
			if module.resources == nil {
				t.Error("the resource map has not been initialized")
			}
		})
		t.Run("test if controller with unique alias was registered successfully", func(t *testing.T) {
			if err := module.Register("controller", nil); err != nil {
				t.Error("the controller should have been registered successfully")
			}
			if _, ok := module.resources["controller"]; !ok {
				t.Error("module should have registered controller")
			}
		})
		t.Run("test if module is able to apply callback to registered controllers", func(t *testing.T) {
			var alias string
			module.Controllers(func(name string, _ Controller) bool {
				alias = name
				return true
			})
			if alias != "controller" {
				t.Error("callback function has not been properly applied")
			}
		})
		t.Run("test if module skips the rest of controllers if callback returns false", func(t *testing.T) {
			var called int
			module.Register("another-controller", nil)
			module.Controllers(func(string, Controller) bool {
				called++
				return false
			})
			if called != 1 {
				t.Error("callback was expected to be called only once")
			}
		})
		t.Run("test if error is returned registering controller with non-unique alias", func(t *testing.T) {
			if err := module.Register("controller", nil); err == nil {
				t.Error("should return an error")
			}
		})
		t.Run("test if available controller was unregistered successfully", func(t *testing.T) {
			if err := module.Unregister("controller"); err != nil {
				t.Error("controller should have been unregistered successfully")
			}
			if _, ok := module.resources["controller"]; ok {
				t.Error("module should not have registered controller")
			}
		})
		t.Run("test if error is returned trying to unregister unavailable controller", func(t *testing.T) {
			if err := module.Unregister("controller"); err == nil {
				t.Error("should return an error")
			}
		})
	})
}
