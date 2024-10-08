package word

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWords(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output map[string]int
	}{
		{"empty string", "", make(map[string]int)},
		{"string without words", " #$%#$~@#%", make(map[string]int)},
		{"string is word", "sdfdgwe", map[string]int{
			"sdfdgwe": 1,
		}},
		{"string is word surrounded separators", ",sdfdgwe}", map[string]int{
			"sdfdgwe": 1,
		}},
		{"some words", "булок,sdfdgwe}французских,français:булок,Булок/Ta4ka, k8s, log4j", map[string]int{
			"булок":       2,
			"sdfdgwe":     1,
			"французских": 1,
			"français":    1,
			"Булок":       1,
			"Ta4ka":       1,
			"k8s":         1,
			"log4j":       1,
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.output, Words(tc.input))
		})
	}
}
