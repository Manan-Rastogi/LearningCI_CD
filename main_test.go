package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T){
	// Positive Nums
	t.Run("Add Positive Numbers", func(t *testing.T) {
		a := rand.Intn(100) + 1
		b := rand.Intn(100) + 1

		expectedSum := a + b

		actualSum := add(a, b)

		require.Equal(t, expectedSum, actualSum, "Sum of %v and %v should be %v", a, b, expectedSum)
		require.GreaterOrEqual(t, actualSum, a, "Sum(%v) should be greater than individual number a(%v).", actualSum, a)
		require.GreaterOrEqual(t, actualSum, b, "Sum(%v) should be greater than individual number b(%v).", actualSum, b)
		require.NotEmpty(t, actualSum, "Sum of Positive Numbers can not be empty.")
	})

	// Negative Nums
	t.Run("Add Negative Numbers", func(t *testing.T) {
		a := -rand.Intn(100) - 1
		b := -rand.Intn(100) - 1

		expectedSum := a + b

		actualSum := add(a, b)

		require.Equal(t, expectedSum, actualSum, "Sum of %v and %v should be %v", a, b, expectedSum)
		require.LessOrEqual(t, actualSum, a, "Sum(%v) should be less than or equal to individual number a(%v).", actualSum, a)
		require.LessOrEqual(t, actualSum, b, "Sum(%v) should be less than or equal to individual number b(%v).", actualSum, b)
		require.NotEmpty(t, actualSum, "Sum of Negative Numbers should not be empty.")
	})

}