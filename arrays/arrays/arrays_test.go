package arrays

import (
	"testing"
)

func TestSliceOperations(t *testing.T) {
	// Initialize the prices array as in the main function
	prices := [4]float64{14.99, 198.99, 54.99, 20}

	// Test case: Demonstrate len function
	t.Run("GetLen", func(t *testing.T) {
		length := len(prices)
		if length != 4 {
			t.Errorf("Expected length [%d] got [%d].\n", 4, length)
		}
	})

	// Test case 1: Verify slice prices[1:3] (2nd through 3rd elements)
	t.Run("SliceSecondToThird", func(t *testing.T) {
		slice := prices[1:3]
		expected := []float64{198.99, 54.99}
		if len(slice) != 2 {
			t.Errorf("Expected slice length 2, got %d", len(slice))
		}
		for i, v := range expected {
			if slice[i] != v {
				t.Errorf("Expected slice[%d] = %f, got %f", i, v, slice[i])
			}
		}
	})

	// Test case 2: Verify slice prices[:3] (1st through 3rd elements)
	t.Run("SliceFirstToThird", func(t *testing.T) {
		slice := prices[:3]
		expected := []float64{14.99, 198.99, 54.99}
		if len(slice) != 3 {
			t.Errorf("Expected slice length 3, got %d", len(slice))
		}
		for i, v := range expected {
			if slice[i] != v {
				t.Errorf("Expected slice[%d] = %f, got %f", i, v, slice[i])
			}
		}
	})

	// Test case 3: Verify slice prices[1:] (2nd element through end)
	t.Run("SliceSecondToEnd", func(t *testing.T) {
		slice := prices[1:]
		expected := []float64{198.99, 54.99, 20}
		if len(slice) != 3 {
			t.Errorf("Expected slice length 3, got %d", len(slice))
		}
		for i, v := range expected {
			if slice[i] != v {
				t.Errorf("Expected slice[%d] = %f, got %f", i, v, slice[i])
			}
		}
	})

	// Test case 4: Verify modification through slices
	t.Run("SliceModification", func(t *testing.T) {
		featuredPrices := prices[1:]
		highlightedPrices := featuredPrices[:1]
		originalPrice := prices[1]
		highlightedPrices[0] += 1
		if prices[1] != originalPrice+1 {
			t.Errorf("Expected modified price %f, got %f", originalPrice+1, prices[1])
		}
		if highlightedPrices[0] != originalPrice+1 {
			t.Errorf("Expected highlightedPrices[0] = %f, got %f", originalPrice+1, highlightedPrices[0])
		}
	})

	// Test case 5: Verify slice bounds and capacity
	t.Run("SliceBoundsAndCapacity", func(t *testing.T) {
		slice := prices[1:3]
		if len(slice) != 2 {
			t.Errorf("Expected slice length 2, got %d", len(slice))
		}
		if cap(slice) != 3 {
			t.Errorf("Expected slice capacity 3, got %d", cap(slice))
		}
	})

	// Test case 6: Reslice an array to show referential nature of slices
	testArray := [4]int{1, 2, 3, 4}
	t.Run("ExperimentsWithSlices", func(t *testing.T) {

		sliceA := testArray[1:3]
		if len(sliceA) != 2 {
			t.Errorf("Expected slice length 2, got %d", len(sliceA))
		}
		if len(sliceA) == len(testArray) {
			t.Errorf("Expected slice and orig array to have different lengths, got %d for slice and %d for orig array", len(sliceA), len(testArray))
		}
		if !(sliceA[0] == 2 && sliceA[1] == 3) {
			t.Errorf("Expected values 2 and 3 in slice, got %d and %d", sliceA[0], sliceA[1])
		}
		if len(sliceA[:3]) != 3 {
			t.Errorf("Expected resliced slice length 3, got %d", len(sliceA[:4]))
		}
	})
}
