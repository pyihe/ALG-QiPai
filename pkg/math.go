package pkg

func factorial(m int) (n int) {
	n = 1
	for i := 2; i <= m; i++ {
		n *= i
	}
	return
}

// combination (n,m)的组合情况(从n个数中取m个进行组合)
func combination(n, m int) (data int) {
	data = factorial(n) / (factorial(n-m) * factorial(m))
	return
}

// Combination 从src中选取length个元素出来
// 组合是一个基本的数学问题，本程序的目标是输出从n个元素中取m个的所有组合。
// 例如从[1,2,3]中取出2个数，一共有3中组合：[1,2],[1,3],[2,3]。（组合不考虑顺序，即[1,2]和[2,1]属同一个组合）
// 从n中选取m个组合
// 例如: 从make([]int, 10)中选取5个元素的所有组合，此时n=10, m=5
// 返回的combination为为每个元素的索引组成的二维切片
func Combination(n, m int) (combs [][]int) {
	//（1）创建有n个元素数组，数组元素的值为1表示选中，为0则没选中。
	//（2）初始化，将数组前m个元素置1，表示第一个组合为前m个数。
	//（3）从左到右扫描数组元素值的“10”组合，找到第一个“10”组合后将其变为“01”组合，同时将其左边的所有“1”全部移动到数组的最左端。
	//（4）当某次循环没有找到“10“组合时，说明得到了最后一个组合，循环结束。

	if n < m || m < 1 {
		return
	}
	// cap直接根据公式计算得出
	combs = make([][]int, 0, combination(n, m))
	// 初始化n个元素，前m个置为1，其余置为0
	idx := make([]int, n)
	for i := 0; i < m; i++ {
		idx[i] = 1
	}

	// 第一种情况添加进结果集合
	add(&combs, idx)

	for {
		flag := false
		for i := 0; i < n-1; i++ {
			if idx[i] == 1 && idx[i+1] == 0 {
				flag = true
				idx[i], idx[i+1] = idx[i+1], idx[i]
				if i > 1 {
					moveToLeft(idx[:i])
				}
				add(&combs, idx)
				break
			}
		}
		if !flag {
			break
		}
	}
	return
}

func add(dst *[][]int, ele []int) {
	data := make([]int, len(ele))
	copy(data, ele)
	*dst = append(*dst, data)
}

func moveToLeft(s []int) {
	sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 1 {
			sum++
		}
	}

	for i := 0; i < len(s); i++ {
		if i < sum {
			s[i] = 1
		} else {
			s[i] = 0
		}
	}
}
