// there can't be functions from different packages in the same directory
package cube

// package-private functions, types and global variables start with lowercase
// they can be only invoked/referenced inside the same package where they
// are defined
func pow3(x float64) float64 {
	return x * x * x
}

// public stuff starts with Uppercase
func Volume(edge float64) float64 {
	return pow3(edge)
}
