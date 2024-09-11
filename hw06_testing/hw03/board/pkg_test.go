package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaint(t *testing.T) {
	paintCases := []struct {
		name  string
		size  int
		black string
		white string
		out   string
	}{
		{"size 1", 1, "#", " ", "___\n|#|\n⎺⎺⎺\n"},
		{"size 0", 0, "#", " ", "__\n⎺⎺\n"},
		{"size -1", -1, "#", " ", "__\n⎺⎺\n"},
		{"size 1 reverse blocks", 1, " ", "#", "___\n| |\n⎺⎺⎺\n"},
		{"size 2 reverse blocks", 2, " ", "#", "____\n| #|\n|# |\n⎺⎺⎺⎺\n"},
	}

	for _, tt := range paintCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, Paint(tt.size, tt.black, tt.white))
		})
	}
}
