package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var nums1, nums2 []int

	fmt.Fscan(in, &nums1, &nums2)
	findMedianSortedArrays(nums1, nums2)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

}
