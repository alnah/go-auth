package testhelper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	testCases := []struct {
		name     string
		length   uint
		expected int
	}{
		{"generate a string corresponding to the specified length", 10, 10},
		{"generate an empty string when the specified length is zero", 0, 0},
		{"generate a string of a very large length", 10000, 10000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Len(t, RandomString(tc.length), tc.expected)
		})
	}
}
