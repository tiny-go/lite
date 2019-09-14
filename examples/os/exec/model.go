package exec

// Commands represents a list of executable commands.
type Commands []string

// Lookup returns true if provided command exists in the list.
func (cs Commands) Lookup(cmd string) bool {
	for _, curr := range cs {
		if curr == cmd {
			return true
		}
	}
	return false
}
