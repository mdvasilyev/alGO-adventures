package main

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	keyToElem map[int]int
	keyList   []int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Constructorr() RandomizedSet {
	return RandomizedSet{
		keyToElem: make(map[int]int),
		keyList:   make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.keyToElem[val]; !ok {
		this.keyList = append(this.keyList, val)
		this.keyToElem[val] = len(this.keyList) - 1
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if res, ok := this.keyToElem[val]; !ok {
		return false
	} else {
		length := len(this.keyList)
		lastElem := this.keyList[length-1]
		this.keyToElem[lastElem] = res
		this.keyList[res], this.keyList[length-1] = this.keyList[length-1], this.keyList[res]
		this.keyList = this.keyList[:length-1]
		delete(this.keyToElem, val)
	}
	return true
}

func (this *RandomizedSet) GetRandom() int {
	randIdx := rand.Intn(len(this.keyList))
	return this.keyList[randIdx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
