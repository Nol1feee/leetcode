func removeDuplicates(nums []int) int {
    uniqueNumIndex := 1
    currNum := nums[0]

    for _, v := range nums {
        if v == currNum {
            continue
        }

        nums[uniqueNumIndex] = v
        uniqueNumIndex++
        currNum = v
    }

    return uniqueNumIndex
}

//