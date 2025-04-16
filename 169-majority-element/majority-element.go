// что требуется:
//   - вернуть элемент слайса, который встречается наибольшее количетсво раз за O(1)
// как решить:
//   - за O(n) понятно как, просто map[val]count, а вот O(1)..?
//   - можно идти с 2 концов, но все равно не то
// итог:
//   - 

func majorityElement(nums []int) int {
    dict := make(map[int]int)

    for _, v := range nums {
        dict[v]++
    }

    res, maxCount := 0, 0


    for k, v := range dict {
        fmt.Println(k, v)
        if v > maxCount {
            maxCount = v
            res = k
        }
    }

    return res
}