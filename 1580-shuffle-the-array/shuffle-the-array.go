// уже лучше, но недостаточно хорошо 

func shuffle(nums []int, n int) []int {
    a1, a2 := make([]int, n), make([]int, n)
    copy(a1, nums[:n])
    copy(a2, nums[n:])
    index := 0

    for i := 0; i < len(nums); i++ {
        if i % 2 == 0 {
            nums[i] = a1[index]
            continue
        }
        nums[i] = a2[index]
        index++ 
    }

    return nums
}