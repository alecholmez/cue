package mail_test

import (
	"testing"

	"cuelang.org/go/pkg/internal/builtintest"
)

func TestBuiltin(t *testing.T) {
	builtintest.Run("mail", t)
}
