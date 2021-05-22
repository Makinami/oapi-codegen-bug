package app

import (
	"net/url"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestBindRequiredDateParam(t *testing.T) {
	t.Run("minimal", func(t *testing.T) {
		day := types.Date{} // NOT a pointer
		queryParams := url.Values{}

		// required parameter "day" is not provided so BindQueryParameter should return an error
		err := runtime.BindQueryParameter("form", true, true, "day", queryParams, &day)
		assert.NotNil(t, err)
	})

	t.Run("generated struct", func(t *testing.T) {
		params := GetPeopleBornOnDayParams{}
		queryParams := url.Values{}

		// required parameter "day" is not provided so BindQueryParameter should return an error
		err := runtime.BindQueryParameter("form", true, true, "day", queryParams, &params.Day)
		assert.NotNil(t, err)
	})
}
