package comparator

import (
	"fmt"
	"testing"

	"github.com/kotoproger/home_work_basic/hw04_struct_comparator/book"
	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	testCase := []struct {
		name     string
		rate1    float64
		rate2    float64
		year1    int
		year2    int
		size1    int
		size2    int
		compares []struct {
			mode   CompareType
			result bool
		}
	}{
		{"book with different params", 1.235, 1.234, 2001, 2000, 20, 19, []struct {
			mode   CompareType
			result bool
		}{{CompareRate, true}, {CompareSize, true}, {CompareYear, true}}},
		{"book with same params", 1.234, 1.234, 2000, 2000, 20, 20, []struct {
			mode   CompareType
			result bool
		}{{CompareRate, false}, {CompareSize, false}, {CompareYear, false}}},
	}

	for _, tt := range testCase {
		for _, ttt := range tt.compares {
			t.Run(fmt.Sprintf("%s/%s/%v", tt.name, ttt.mode, ttt.result), func(t *testing.T) {
				book1 := book.NewBook(1, "no matter", "no matter", tt.year1, tt.size1, tt.rate1)
				book2 := book.NewBook(2, "no matter2", "no matter2", tt.year2, tt.size2, tt.rate2)
				comparator := NewBookComparator(ttt.mode)

				assert.Equal(t, ttt.result, comparator.Compare(book1, book2))
			})
		}
	}
}
