package jzoffer

// 给你一个正整数数组 nums ，对 nums 所有元素求积之后，
// 找出并返回乘积中 不同质因数 的数目。

func distinctPrimeFactors(nums []int) int {
	prime := func(x int) []int {
		vs := []int{}
		for i := 2; i*i <= x; i++ { // 查询边界,最多到sqrt(x)
			if x%i != 0 {
				continue
			}
			for x /= i; x%i == 0; x /= i {
			}
			vs = append(vs, i)
		}
		if x > 1 {
			vs = append(vs, x)
		}
		return vs
	}

	res := make(map[int]struct{})

	for _, v := range nums {
		vs := prime(v)
		for _, v := range vs {
			res[v] = struct{}{}
		}
	}
	return len(res)
}

// 效率差很多,这里需要i可能循环到x,上面那种写法最多到sqrt(x)
func distinctPrimeFactors1(nums []int) int {
	prime := func(x int) []int {
		vs := []int{}
		for i := 2; x > 1; i++ {
			for x%i == 0 {
				vs = append(vs, i)
				x /= i
			}
		}
		return vs
	}

	res := make(map[int]struct{})

	for _, v := range nums {
		vs := prime(v)
		for _, v := range vs {
			res[v] = struct{}{}
		}
	}
	return len(res)
}
