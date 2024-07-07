package list

import (
	"reflect"
	"testing"
)

func TestFromString(t *testing.T) {
	t.Run("Test case 1: Valid input with pruning quotes", func(t *testing.T) {
		str := " 'one' , 'two' , 'three' "
		expected := []string{"one", "two", "three"}
		result := FromString(str, ",", true)
		if !reflect.DeepEqual(result, expected) {
			t.Fatalf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Test case 2: Valid input without pruning quotes", func(t *testing.T) {
		str := " 'one' , 'two' , 'three' "
		expected := []string{" 'one' ", " 'two' ", " 'three' "}
		result := FromString(str, ",", false)
		if !reflect.DeepEqual(result, expected) {
			t.Fatalf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Test case 3: Empty input string", func(t *testing.T) {
		str := ""
		var expected []string
		result := FromString(str, ",", true)
		if len(result) != len(expected) {
			t.Fatalf("size mismatch")
		}
		if !reflect.DeepEqual(result, expected) {
			t.Fatalf("Expected %v, but got %v", expected, result)
		}
	})
}
