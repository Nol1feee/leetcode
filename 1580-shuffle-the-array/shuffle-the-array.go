// решение максимально в лоб, мне максимально не нравится

func shuffle(nums []int, n int) []int {
    res := make([]int, len(nums))
    a1 := nums[:n]
    a2 := nums[n:]
    index := 0

    for i := 0; i < len(nums); i++ {
        if i % 2 == 0 {
            res[i] = a1[index]
            continue
        }
        res[i] = a2[index]
        index++ 
    }

    return res

}