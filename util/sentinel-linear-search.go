package util

func SentinelLinearSearch(list List, element string) int {
    n := len(list) - 1

    last := list[n]
    list[n] = element
    i := 0
    for ; list[i] != element; i++ {
      // ok
    }
    list[n] = last

    var result int
    if i < n || list[n] == element {
        result = i
    } else {
      result = -1
    }

    return result
}
