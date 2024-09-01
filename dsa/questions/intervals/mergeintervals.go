package intervals

import "sort"

func merge(intervals [][]int) (res [][]int) {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	newInterval := intervals[0]
	res = append(res, newInterval)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= newInterval[1] {
			newInterval[1] = max(intervals[i][1], newInterval[1])
		} else {
			newInterval = intervals[i]
			res = append(res, newInterval)
		}
	}
	return res
}
