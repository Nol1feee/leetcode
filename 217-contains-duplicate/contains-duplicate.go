// что требуется:
//   - есть дубликаты -> true, нет -> false
// как решить:
//   - dict
// итог:
//   - ofc easy

func containsDuplicate(nums []int) bool {
    dict := make(map[int]bool, len(nums))

    for _, v := range nums {
        if dict[v] {
            return true 
        }
        dict[v] = true
    }

    return false
}