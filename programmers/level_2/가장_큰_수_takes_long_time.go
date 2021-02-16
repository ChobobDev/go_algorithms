import "strconv"
import "sort"

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}
	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func solution(numbers []int) string {
	var glued_nums []int
	var permus [][]int
	permus = permutations(numbers)
	perm := make([][]int, len(permus))
	for i, _ := range perm {
		var glued_num string
		perm[i] = make([]int, len(permus[i]))
		for j := range perm[i] {
			glued_num += strconv.Itoa(permus[i][j])
		}
		result_int, _ := strconv.Atoi(glued_num)
		glued_nums = append(glued_nums, result_int)
	}
	sort.Ints(glued_nums)
	desired_index := len(glued_nums) - 1
	answer := strconv.Itoa(glued_nums[desired_index])

	return answer
}