package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// with string reverse
func isPalindrome(value string) bool {
	var temp string

	for i := len(value) - 1; i >= 0; i-- {
		temp = temp + string(value[i])
	}

	return value == temp
}

func TestPalinDromeWithReverseString(t *testing.T) {
	t.Run("false test", func(t *testing.T) {
		assert.False(t, isPalindrome("ab"), "Result Should return false")
		assert.False(t, isPalindrome("abab"), "Result Should return false")
		assert.False(t, isPalindrome("kodcok"), "Result Should return false")
		assert.False(t, isPalindrome("moon"), "Result Should return false")
	})

	t.Run("true test", func(t *testing.T) {
		assert.True(t, isPalindrome("a"), "Result should return true")
		assert.True(t, isPalindrome("aba"), "Result should return true")
		assert.True(t, isPalindrome("kodok"), "Result should return true")
		assert.True(t, isPalindrome(""), "Result should return true")
	})
}

// without reverse string
func isPalindromeWithoutReverseString(value string) bool {
	for i := 0; i < len(value)/2; i++ {
		firstIndex := string(value[i])
		lastIndex := string(value[len(value)-i-1])

		if firstIndex != lastIndex {
			return false
		}
	}

	return true
}

func TestPalinDromeWithoutReverseString(t *testing.T) {
	t.Run("false test", func(t *testing.T) {
		assert.False(t, isPalindromeWithoutReverseString("ab"), "Result Should return false")
		assert.False(t, isPalindromeWithoutReverseString("abab"), "Result Should return false")
		assert.False(t, isPalindromeWithoutReverseString("kodcok"), "Result Should return false")
		assert.False(t, isPalindromeWithoutReverseString("andre"), "Result Should return false")
	})

	t.Run("true test", func(t *testing.T) {
		assert.True(t, isPalindromeWithoutReverseString("a"), "Result should return true")
		assert.True(t, isPalindromeWithoutReverseString("aba"), "Result should return true")
		assert.True(t, isPalindromeWithoutReverseString("kodok"), "Result should return true")
		assert.True(t, isPalindromeWithoutReverseString(""), "Result should return true")
	})
}

// recursive function
func isPalingdromeRecursive(value string, i int) bool {
	if i < len(value)/2 {
		if string(value[i]) != string(value[len(value)-i-1]) {
			return false
		} else {
			return isPalingdromeRecursive(value, i+1)
		}
	}
	return true
}

func isPalindromeWithRecursive(value string) bool {
	return isPalingdromeRecursive(value, 0)
}

func TestPalinDromeWithRecursiveFunc(t *testing.T) {
	t.Run("false test", func(t *testing.T) {
		assert.False(t, isPalindromeWithRecursive("ab"), "Result Should return false")
		assert.False(t, isPalindromeWithRecursive("abab"), "Result Should return false")
		assert.False(t, isPalindromeWithRecursive("kodcok"), "Result Should return false")
		assert.False(t, isPalindromeWithRecursive("andre"), "Result Should return false")
	})

	t.Run("true test", func(t *testing.T) {
		assert.True(t, isPalindromeWithRecursive("a"), "Result should return true")
		assert.True(t, isPalindromeWithRecursive("aba"), "Result should return true")
		assert.True(t, isPalindromeWithRecursive("kodok"), "Result should return true")
		assert.True(t, isPalindromeWithRecursive(""), "Result should return true")
	})
}
