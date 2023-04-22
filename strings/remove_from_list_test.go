package strings

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestRemoveElementsFromList(t *testing.T) {
	removeOne := []string{Alfa, Alfa, Bravo, Charlie}
	removeOne = RemoveFromList(removeOne, []string{Alfa})
	assert.Any(t, []string{Bravo, Charlie}, removeOne)

	multiple := []string{Alfa, Alfa, Bravo, Charlie}
	multiple = RemoveFromList(multiple, []string{Alfa, Bravo})
	assert.Any(t, []string{Charlie}, multiple)
}
