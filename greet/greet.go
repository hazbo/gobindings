package greet

//#include "greet.h"
import "C"

/**
 * Wraps out C function to return a
 * Go string
 */
func SayHello(name string) string {
	return C.GoString(C.say_hello(C.CString(name)))
}