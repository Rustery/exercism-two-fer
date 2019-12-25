// Package twofer — Two-fer or 2-fer is short for two for one. One for you and one for me.
// Just an exercise for Exercism.io
package twofer

// ShareWith take a name as argument and return a string with this name. However, if the name
// is missing, function returns the string with just 'you'.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
