package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var str []string = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupAnagrams(str)
}
func groupAnagrams(strs []string) [][]string {
	ans := make(map[string][]string)

	for _, v := range strs {
		str := v
		split := strings.Split(str, "")
		sort.Strings(split)
		r := strings.Join(split, "")
		_, isP := ans[r]
		if !isP {
			ans[r] = []string{}
		}
		ans[r] = append(ans[r], v)

	}
	var lst = [][]string{}
	for _, v := range ans {
		fmt.Println(v)
		lst = append(lst, v)
	}
	return lst

}
