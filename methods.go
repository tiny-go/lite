package static

import "strings"

type methods []string

func (ms *methods) add(method string) {
	*ms = append(*ms, method)
}

func (ms *methods) join() string {
	return strings.Join(*ms, ",")
}

func (ms *methods) empty() bool {
	return len(*ms) == 0
}
