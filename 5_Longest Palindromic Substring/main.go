package main

func main() {

}

func longestPalindrome(s string) string {
	var (
		result string
		supStr []byte
		lens   = len(supStr)
		supStart int
	)

	supStr = make([]byte, 1000)
	supStr = append(supStr, []byte(s)...)

	if lens == 0 || lens == 1 {
		return s
	}

	for i := 0; i < lens; i++ {

		supStart = i
		if start+1 < lens && supStr[supStart] == supStr[supStart+1] {
			var (
				start int = supStart
				end   int
			)
			end = start + 1
			if start >= 0 && end < lens && supStr[start] == supStr[end] {
				if len([]byte(result)) <= len(supStr[start:end+1]) {
					result = string(supStr[start : end+1])
				}
	
				start--
				end++
			}
		} 
		
		
		if start+2 < lens && supStr[start] == supStr[start+2] {
			var (
				start int = supStart
				end   int
			)
			end = start + 2
			if start >= 0 && end < lens && supStr[start] == supStr[end] {
				if len([]byte(result)) <= len(supStr[start:end+1]) {
					result = string(supStr[start : end+1])
				}
	
				start--
				end++
			}
		}
	}

	return result
}
