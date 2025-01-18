// ок
// успел забыть, что  append к make([]int, len(nums) будет добавлять элементы "сверху"

func shuffle(nums []int, n int) []int {
    res := make([]int, 0, len(nums))

    for i := 0; i < n; i++ {
        res = append(res, nums[i], nums[n+i])
    }

    return res
}