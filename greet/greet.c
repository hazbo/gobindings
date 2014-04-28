#include <stdlib.h>
#include <string.h>
#include "greet.h"

/**
 * A fairly horrible function to
 * strcat our name with a
 * greeting
 */
char *say_hello(char *name)
{
	// Assign 80 bytes for the full
	// greeting string
	char *greeting = malloc(80);

	char message[] = "Hello ";

	// Concat the strings together
	strcat(greeting, message);
	strcat(greeting, name);
	strcat(greeting, "!");

	return greeting;
}