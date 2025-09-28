package main

import (
	"fmt"
	"sort"
)

/*
*

	以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	// 1. 按区间起始值排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 2. 初始化结果集
	result := [][]int{intervals[0]}
	fmt.Println(result)

	// 3. 遍历合并
	for _, interval := range intervals[1:] {
		last := &result[len(result)-1] // 引用结果集最后一个区间
		if interval[0] <= (*last)[1] {
			// 重叠：更新结束值为较大者
			if interval[1] > (*last)[1] {
				(*last)[1] = interval[1]
			}
		} else {
			// 不重叠叠：直接加入
			result = append(result, interval)
		}
	}

	return result
}

/*
*

	给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
	考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
	更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
	返回 k 。
*/
func removeDuplicates(nums []int) int {
	m := make(map[int]int)
	n := []int{}
	for _, k := range nums {
		_, exist := m[k]
		if exist {
			continue
		} else {
			n = append(n, k)
			m[k] = 1
		}
	}
	fmt.Println(n)
	return len(n)
}

/*
*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
将大整数加 1，并返回结果的数字数组。
*/
func plusOne(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		digits[i]++

		if digits[i] < 10 {
			return digits
		}

		digits[i] = 0
	}

	return append([]int{1}, digits...)
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
