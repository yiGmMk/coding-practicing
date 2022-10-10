package jzoffer

import (
	"fmt"
	"sort"
)

/*
there are two integer arrays starting,ending,each of size n,that represent=n intervals
where the ith interval is [starting[i],ending[i]]

determine the number of ways three non-overlapping intervals can be selected from the n intervals

example:
n=5
starting=[1,2,4,3,7]
ending=[3,4,6,5,8]
expected output:1
5 intervals {1,3},{2,4},{4,6},{3,5},{7,8},view as line segments 1 through 5
there are only one way to select three non-overlapping intervals,line 1,3 and 5 with ranges {1,3},{4,6},{7,8}

n=4
starting=[5,2,3,7]
ending=[7,2,4,8]
expected output=2
intervals:{2,2},{3,4},{5,7},{7,8}

从n个区间(闭区间)中选出3个不重叠的,计算有多少种选择方式,返回数量.区间都是闭区间

constraints
1<=n<=1e5
|starting|,|ending|=n
1<=starting[i],ending[i]<=1e9
*/

var debug = true

func logf(format string, args ...interface{}) {
	if !debug {
		return
	}
	fmt.Printf(format, args...)
	fmt.Println()
}

func getIntervals(st, ed []int) [][]int {
	res := [][]int{}
	for i, v := range st {
		res = append(res, []int{v, ed[i]})
	}
	return res
}

func equal(intervals [][]int, i, j int) bool {
	return intervals[i][0] == intervals[j][0] && intervals[i][1] == intervals[j][1]
}

// n*n*n
func getThreeNonOverlappingIntervalsTimeout(starting, ending []int) int {
	intervals := getIntervals(starting, ending)
	// 右边界排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })

	n := len(intervals)
	var res int
	for i := 0; i < n; i++ {
		if i > 0 && equal(intervals, i, i-1) {
			continue
		}
		// i.r<j.l
		j := sort.Search(n, func(id int) bool { return id > i && intervals[i][1] < intervals[id][0] })
		if j >= n {
			continue
		}

		for ; j < n; j++ {
			if equal(intervals, j, j-1) {
				continue
			}
			// j.r<k.l
			k := sort.Search(n, func(id int) bool { return id > i && intervals[j][1] < intervals[id][0] })
			if k >= n {
				continue
			}
			for ; k < n; k++ {
				if equal(intervals, k, k-1) {
					continue
				}
				// k.l > j.r
				if intervals[k][0] > intervals[j][1] {
					res++
					logf("intervals:%+v,%+v,%+v", intervals[i], intervals[j], intervals[k])
				}
			}
		}
	}
	return res
}

// n*n*n
func getThreeNonOverlappingIntervalsTimeout2(starting, ending []int) int {
	intervals := getIntervals(starting, ending)
	// 右边界排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })

	n := len(intervals)
	var res int
	for i := 0; i < n; i++ {
		if i > 0 && equal(intervals, i, i-1) {
			continue
		}
		for j := i + 1; j < n; j++ {
			// i.r<j.l
			if !(intervals[j][0] > intervals[i][1]) {
				continue
			}
			if equal(intervals, j, j-1) {
				continue
			}
			// j.r<k.l
			for k := j + 1; k < n; k++ {
				if equal(intervals, k, k-1) {
					continue
				}
				// k.l > j.r
				if intervals[k][0] > intervals[j][1] {
					res++
					logf("intervals:%+v,%+v,%+v", intervals[i], intervals[j], intervals[k])
				}
			}
		}
	}
	return res
}

// https://leetcode.cn/circle/discuss/YJ3bL9/
// 按starting排序
// 再对每一个区间，求左边不与该区间重叠的区间个数 (即左边 ending[l] < starting[i] 的个数)，和右边不重叠的区间个数，乘起来，并累加到答案中。
// 求左边不与该区间重叠的区间个数可以用线段树，或者树状数组。鉴于数据范围，需要先对数据先离散化。
// 作者：Grobyc
// 链接：https://leetcode.cn/circle/discuss/YJ3bL9/view/Rt8mRc/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func getThreeNonOverlappingIntervals(starting, ending []int) int {

	intervals := getIntervals(starting, ending)
	// 右边界排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })

	var (
		res int
		n   = len(intervals)
		// dp  = make(map[int]int, 0)
	)
	// dp[i]=dp[i-1]+
	for i := 0; i < n; i++ {
		k := n - 1
		for j := i + 1; j < k; j++ {

		}
	}

	return res
}

// 可以用线段树统计每个线段相交的线段数量，这样就能求出每条线段i之前有多少线段不交，记录为si，然后处理i+1线段时，
// 对于所有不与i+1相交的线段的求和，就是以i+1为结尾的方案数，最后把所有的方案加起来就成了。
// 作者：adanzl
// 链接：https://leetcode.cn/circle/discuss/YJ3bL9/view/5sy86m/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// from sortedcontainers import SortedList
// inters=sorted(set(zip(starting,ending)))
// left,right=SortedList(),SortedList([b for b,e in inters])
// res=0
// for b,e in inters:
//     right.remove(b)
//     l,r=left.bisect_left(b),len(right)-right.bisect_right(e)
//     res+=l*r
//     left.add(e)
// print(res)
// 输入：
// starting=[1, 2, 4, 3, 7]
// ending=[3, 4, 6, 5, 8]
// 结果：1
// 输入：
// starting=[5, 2, 3, 7]
// ending=[7, 2, 4, 8]
// 结果 ：2
// 思路，去重所有区间，按左端点排序。
// 依次读入每个区间，计算左端点大于左边所有右端点的个数，计算右端点小于右边所有左端点的个数。
// 两者相乘即是该区间作出的贡献，累加得到结果。
// 作者：蒺藜
// 链接：https://leetcode.cn/circle/discuss/YJ3bL9/view/lpz7Vt/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
