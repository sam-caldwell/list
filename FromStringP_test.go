package list

import (
	"reflect"
	"testing"
)

func TestFromStringP(t *testing.T) {
	t.Run("Test case 1: Valid input with pruning quotes", func(t *testing.T) {
		str := " 'one' , 'two' , 'three' "
		expected := []string{"one", "two", "three"}
		result := FromStringP(&str, ",", true)
		if !reflect.DeepEqual(result, expected) {
			t.Fatalf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Test case 2: Valid input without pruning quotes", func(t *testing.T) {
		str := " 'one' , 'two' , 'three' "
		expected := []string{" 'one' ", " 'two' ", " 'three' "}
		result := FromStringP(&str, ",", false)
		if !reflect.DeepEqual(result, expected) {
			t.Fatalf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Test case 3: Empty input string pointer", func(t *testing.T) {
		str := ""
		var expected []string
		result := FromStringP(&str, ",", true)
		if !reflect.DeepEqual(result, expected) {
			t.Fatalf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Test case 4: Nil input string pointer", func(t *testing.T) {
		var nilStr *string
		var expected []string
		result := FromStringP(nilStr, ",", true)
		if !reflect.DeepEqual(result, expected) {
			t.Fatalf("Expected %v, but got %v", expected, result)
		}
	})
}
