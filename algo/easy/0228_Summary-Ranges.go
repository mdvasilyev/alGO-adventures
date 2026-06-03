package main

import "fmt"

func summaryRanges(nums []int) []string {
    length := len(nums)
    if length == 0 {
        return []string{}
    }
    start := nums[0]
    resInts := [][]int{{start}}
    for i := range nums {
        end := nums[i]
        if start == end {
            continue
        }
        resLen := len(resInts)
        lastInterval := resInts[resLen-1]
        if end == lastInterval[len(lastInterval)-1]+1 {
            resInts[resLen-1] = []int{start, end}
        } else {
            resInts = append(resInts, []int{end})
            start = end
        }
    }
    res := make([]string, 0)
    for i := range resInts {
        val := resInts[i]
        if len(val) == 1 {

            res = append(res, fmt.Sprintf("%d", val[0]))
        } else {
            res = append(res, fmt.Sprintf("%d->%d", val[0], val[1]))
        }
    }
    return res
}
