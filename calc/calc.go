package calc

//#include "calc.h"
import "C"

/**
 * Wraps up our C add function
 */
func GoAdd(a int, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}

/**
 * Wraps up our C subtract function
 */
func GoSubtract(a int, b int) int {
	return int(C.subtract(C.int(a), C.int(b)))
}