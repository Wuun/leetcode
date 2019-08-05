package main

import "fmt"

//Given a string, find the length of the longest substring without repeating characters.
// Example 1:
// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Example 2:
// Input: "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:
// Input: "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
//              Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

func main() {
	str := "au"
	fmt.Print(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	var (
		supMap    = make(map[byte]int)
		sourceStr = []byte(s)
		lens      = len(sourceStr)
		result    int
		count     int
	)

	for j := 0; j < lens; j++ {
		if v, ok := supMap[sourceStr[j]]; !ok {
			supMap[sourceStr[j]] = j
		} else {
			j = v
			supMap = make(map[byte]int)
		}

		count = len(supMap)
		if result < count {
			result = count
		}
	}
	return result
}
