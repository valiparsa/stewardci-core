package framework

import (
	"context"
	"testing"
)

// setup has been added after finding compilation errors due to this
// missing function.
// TODO check why setup has been missing
func setup(t *testing.T) context.Context {
	t.Fatal("function not implemented")
	return nil
}
