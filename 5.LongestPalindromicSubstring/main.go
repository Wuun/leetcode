package main

import "fmt"

func main() {
	restlt := longestPalindrome("aabaaaaaaaaaaaaa")
	fmt.Print(restlt)
}

//在这里还是可以有一个优化就是，当我们在进行下一个数的回文的搜索的时候，并不需要逐个计算
//先把它的前后的二分之一以前记录的最长回文的长度，看看有没有可能是回文，不是就跳过。
func longestPalindrome(s string) string {
	var (
		supSlice = []byte(s)
		sliceResult []byte
		lens = len(s)
		result string
	)

	if lens < 2 {
		return s
	}
	for i := 0;i < lens;i++ {
		before := i-1
		after := i+1
		before2 := i
		after2 := i+1
		for before >= 0 && after < lens && supSlice[before] == supSlice[after]{
			if len(sliceResult) <= len(supSlice[before:after+1]){
				sliceResult = supSlice[before:after+1]
			}

			before-=1
			after+=1
		}

		for  before2 >= 0 && after2 < lens && supSlice[before2] == supSlice[after2]{
			if len(sliceResult) <= len(supSlice[before2:after2+1]){
				sliceResult = supSlice[before2:after2+1]
			}
			before2-=1
			after2+=1
		}

		if len(sliceResult) == 0{
			sliceResult = supSlice[i:i+1]
		}
	}

	result = string(sliceResult)
	return result
}