package portainer

/*
	Setup pre-conditions:
	1. The size should be always an even number
	2. We can assume that only accepted symbols are: [{()}] otherwise it will return invalid
*/

// ValidateSetup checks if a setup configuration is built properly
func ValidateSetup(setup string) bool {
	if len(setup)%2 != 0 {
		return false
	}

	closerStack := []rune{}
	openBrackets := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for _, r := range setup {

		// If current character(bracket) is an open bracket stack it's closer pair into the stack
		if closerBracket, ok := openBrackets[r]; ok {
			closerStack = append(closerStack, closerBracket)
			continue
		}

		// If current character(bracket) is a close bracket compare with the top stack element
		l := len(closerStack) - 1
		if l < 0 || r != closerStack[l] {
			return false
		}

		// Pop stack element
		closerStack = closerStack[:l]
	}

	// Return if all stack elements are consumed
	return len(closerStack) == 0
}
