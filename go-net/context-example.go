package main

import (
	"context"
	"fmt"
)

// START OMIT
func main() {
	// Associate a value with the context
	ctx := context.WithValue(context.Background(), "key", 1234)
	doSomething(ctx)
}

// doSomething demonstrates using a context as key/value storage.
// This is almost always a bad idea.
func doSomething(ctx context.Context) {
	key := ctx.Value("key").(int)
	fmt.Println("key =", key)
}

// END OMIT
