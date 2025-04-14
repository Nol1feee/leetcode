// что требуется:
//   - есть дубликаты -> true, нет -> false
// как решить:
//   - dict
// итог:
//   - ofc easy

func containsDuplicate(nums []int) bool {
    dict := make(map[int]struct{}, len(nums))

    for _, v := range nums {
        if _, ok := dict[v]; ok {
            return true 
        }
        dict[v] = struct{}{}
    }

    return false
}