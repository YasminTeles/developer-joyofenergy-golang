package priceplans

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessfulValidation(t *testing.T) {
	t.Parallel()

	err := validateSmartMeterId("smart-meter-0")
	assert.NoError(t, err)
}

func TestValidationFailureWithMissingID(t *testing.T) {
	t.Parallel()

	err := validateSmartMeterId("")

	assert.Error(t, err)
	assert.EqualError(t, err, "cannot be blank")
}
