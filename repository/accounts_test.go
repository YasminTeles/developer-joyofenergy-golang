package repository_test

import (
	"joi-energy-golang/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPriceIdForSmartMeterId(t *testing.T) {
	t.Parallel()

	accounts := repository.NewAccounts(map[string]string{
		"home-sweet-home": "test-plan",
	})

	plan, err := accounts.PricePlanIdForSmartMeterId("home-sweet-home")

	assert.NoError(t, err)
	assert.Equal(t, "test-plan", plan)
}

func TestNotFoundPriceIdForSmartMeterId(t *testing.T) {
	t.Parallel()

	accounts := repository.NewAccounts(map[string]string{
		"home-sweet-home": "test-plan",
	})

	plan, err := accounts.PricePlanIdForSmartMeterId("smartMeterId")

	assert.Error(t, err)
	assert.Equal(t, "", plan)
}
