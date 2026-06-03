package main

func longestSubarray(nums []int) int {
    length := len(nums)
    if length == 1 {
        return 0
    }
    res := 0
    m := map[int]int{0: 0, 1: 0}
    left, right := 0, 1
    m[nums[left]]++
    for right < length {
        val := nums[right]
        if zeroCnt, _ := m[0]; zeroCnt == 0 || val == 1 {
            m[val]++
            zeroCnt, _ = m[0]
            onesCnt, _ := m[1]
            res = max(res, zeroCnt+onesCnt-1)
            right++
        } else {
            m[nums[left]]--
            left++
        }
    }
    return res
}
