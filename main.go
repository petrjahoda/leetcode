package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	//fmt.Println(isPalindrome(121))
	//fmt.Println(reverse(121))
	//fmt.Println(reverse(911))
	//fmt.Println(reverse(-878))
	//fmt.Println(reverse(1534236469))
	//fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	//fmt.Println(longestCommonPrefix([]string{"jahoda", "jahodi", "jahaman"}))
	fmt.Println(longestCommonPrefix([]string{"cir", "car"}))
	fmt.Println(longestCommonPrefix([]string{"abab", "aba", ""}))
}

func longestCommonPrefix(strs []string) string {
	var result string
	if len(strs) == 0 {
		return result
	}
	if len(strs) == 1 {
		return strs[0]
	}
	maxCharacters := len(strs[0])
	for _, str := range strs {
		if len(str) < maxCharacters {
			maxCharacters = len(str)
		}
	}
	for i := 0; i < maxCharacters; i++ {
		var temp uint8
		var match int
		for index, str := range strs {
			if index == 0 {
				temp = str[i]
				match++
				continue
			}
			if temp == str[i] {
				match++
			}
		}
		if match == len(strs) {
			result += string(temp)
		} else {
			break
		}
	}
	return result
}

func isPalindrome(x int) bool {
	temp := x
	if x < 0 {
		return false
	}
	var result int
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}
	if result == temp {
		return true
	}
	return false
}

func reverse(x int) int {
	var result int
	for x != 0 {
		result = result*10 + x%10
		if result > 2147483647 || result < -2147483648 {
			return 0
		}
		x /= 10
	}
	return result
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
