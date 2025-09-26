package main

import (
	"fmt"
)

func main() {
	//singleNumber([]int{2, 2, 1})
	//println(isValid("()[]{}"))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

/*
*
 14. 最长公共前缀
    编写一个函数来查找字符串数组中的最长公共前缀。
    如果不存在公共前缀，返回空字符串 ""。
*/
func longestCommonPrefix(strs []string) string {
	perfix := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(perfix) {
			perfix = perfix[:len(strs[i])]
		}
		for j := range strs[i] {
			if len(perfix) <= j || perfix[j] != strs[i][j] {
				perfix = perfix[:j]
				return perfix
			}
		}
	}
	return perfix
}

/*
*
*
 20. 有效的括号
    给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
    有效字符串需满足：
    左括号必须用相同类型的右括号闭合。
    左括号必须以正确的顺序闭合。
    每个右括号都有一个对应的相同类型的左括号。
*/
func isValid(s string) bool {
	var sli []rune
	for _, v := range s {
		if v == '(' {
			sli = append(sli, ')')
		} else if v == '{' {
			sli = append(sli, '}')
		} else if v == '[' {
			sli = append(sli, ']')
		} else if len(sli) == 0 || sli[len(sli)-1] != v {
			return false
		} else {
			sli = sli[:len(sli)-1]
		}
	}
	return len(sli) == 0
}

/*
*
 136. 只出现一次的数字
    给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
    你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
*/
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		value, exist := m[v]
		if exist {
			m[v] = value + 1
		} else {
			m[v] = 1
		}
	}
	var one int
	for k, v := range m {
		if v == 1 {
			one = k
		}
	}
	fmt.Println(one)
	return one
}
