package util

import "testing"

func TestSentinelLinearSearch(t *testing.T) {
  t.Log("Testing Sentinel Linear Search...")

  a := List{"9", "4", "2", "8", "3", "5", "7", "10", "1", "6"}

  i1 := SentinelLinearSearch(a, "6")
  if i1 != 9 {
    t.Error("Expected 9, got", i1)
  }

  i2 := SentinelLinearSearch(a, "9")
  if i2 != 0 {
    t.Error("Expected 0, got", i2)
  }

  i3 := SentinelLinearSearch(a, "8")
  if i3 != 3 {
    t.Error("Expected 0, got", i3)
  }
}
