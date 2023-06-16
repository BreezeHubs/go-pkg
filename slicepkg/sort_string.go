package slicepkg

import "sort"

func ReverseString(s []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
}
