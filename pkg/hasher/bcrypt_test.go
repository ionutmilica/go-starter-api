package hasher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBCryptHasher_Make(t *testing.T) {
	hasher := BCryptHasher{Cost: 4}

	hashed, err := hasher.Make("dev")

	assert.NoError(t, err, "no error was expected")
	assert.Len(t, hashed, 60)
}

func TestBCryptHasher_Check(t *testing.T) {
	t.Run("true when hash and plain are equal", func(t *testing.T) {
		h := BCryptHasher{Cost: 4}

		result := h.Check(
			"$2a$04$Ci1pK4uaNRtagNXf7u3FNeaMnbr72kij6u6xXRhRVcxhgR2Az6o4e",
			"dev",
		)

		assert.True(t, result)
	})

	t.Run("false when hash and plain are not equal", func(t *testing.T) {
		h := BCryptHasher{Cost: 4}

		result := h.Check(
			"$2a$04$Ci1pK4uaNRtagNXf7u3FNeaMnbr72kij6u6xXRhRVcxhgR2Az6o4e",
			"devx",
		)

		assert.False(t, result)
	})
}
