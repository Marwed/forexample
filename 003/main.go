package main

import (
	"fmt"
)

func main() {
	ii := []int{}
	var target int
	var temp int
	fmt.Scanln(&target)
	for temp != 9999 {
		fmt.Scanln(&temp)
		if temp == 9999 {
			break
		}
		ii = append(ii, temp)
	}

	yy := TwoNumberSum(ii, target)
	fmt.Println(yy)

}

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	zero_tar := []int{}
	neg := []int{}
	abov := []int{}
	res := []int{}
	for _, v := range array {
		if v < target && v > 0 {
			zero_tar = append(zero_tar, v)

		} else if v < 0 {
			neg = append(neg, v)
		} else {
			abov = append(abov, v)

		}

	}

	if len(abov) == 0 || len(neg) == 0 {

		for i, v := range zero_tar {

			for j := i + 1; j < len(zero_tar); j++ {

				if zero_tar[j] != v && (target-(v+zero_tar[j])) == 0 {
					res = append(res, v, zero_tar[j])
					return res
				}
			}
		}
	} else if len(abov) != 0 || len(neg) != 0 {
		for _, above_v := range abov {
			for _, neg_v := range neg {
				if neg_v+above_v == target {
					res = append(res, neg_v, above_v)
					return res
				}
			}
		}
	} else {
		found_tar := false
		found_z := false

		for _, vv := range array {
			if vv == target {

				found_tar = true
			} else if vv == 0 {
				found_z = false
			}
		}
		if found_z && found_tar {
			res = append(res, target, 0)
			return res
		}
	}

	return nil
}
