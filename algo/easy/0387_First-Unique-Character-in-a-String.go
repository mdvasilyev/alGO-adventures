package main

func firstUniqChar(s string) int {
    m := make(map[byte]int)
    for i := range s {
        m[s[i]]++
    }
    for i := range s {
        if m[s[i]] == 1 {
            return i
        }
    }
    return -1
}
