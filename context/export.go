package context

/*
#cgo LDFLAGS: -lsass -lstdc++ -lm
#cgo CFLAGS:

#include <stdlib.h>
#include <stdio.h>
#include "sass_context.h"
#include "sass_functions.h"

extern int customHandler();

union Sass_Value* CallSassFunction( union Sass_Value* s_args, void* cookie ) {
    // printf("callback yo\n");
    // union Sass_Value* sass_value = NULL;
    int a;
    a = customHandler();
    return sass_make_boolean(false);
}

*/
import "C"
