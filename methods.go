package lite

import "strings"

// Methods is a collection of HTTP methods.
type Methods []string

// Add method to the list of supported actions (ignoring duplicates).
func (ms *Methods) Add(method string) { *ms = append(*ms, method) }

// Join converts list of methods/actions to a string (separated by comma).
func (ms *Methods) Join() string { return strings.Join(*ms, ",") }

// Empty returns true if list of methods is empty.
func (ms *Methods) Empty() bool { return len(*ms) == 0 }
