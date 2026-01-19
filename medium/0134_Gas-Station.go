package main

func canCompleteCircuit(gas []int, cost []int) int {
	if total(gas) < total(cost) {
		return -1
	}
	idx := 0
	tank := 0
	for i := range gas {
		tank += gas[i] - cost[i]
		if tank < 0 {
			tank = 0
			idx = i + 1
		}
	}
	return idx
}

func total(nums []int) int {
	res := 0
	for i := range nums {
		res += nums[i]
	}
	return res
}
