package readings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateSmartMeterIdEmpty(t *testing.T) {
	t.Parallel()

	err := validateSmartMeterId("")

	assert.Error(t, err)
	assert.EqualError(t, err, "cannot be blank")
}

func TestValidateSmartMeterId(t *testing.T) {
	t.Parallel()

	err := validateSmartMeterId("smartMeterId")

	assert.NoError(t, err)
}
