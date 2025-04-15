// что требуется:
//   - вернуть индекс значений из массива, сумма значений которых равна 'target' 
// как решить:
//   - самое простое - перебор, но сложность в таком случае nˆ2
//   - target - currNum -> targetForUs; for nums {map[val]index} return
//     - проблемы с одинаковыми значениями на разных индексах
//       - можно вообще за 1 итерацию с мапой
// итог:
//   - изи, можно на собесе стажеру/джуну дать

func twoSum(nums []int, target int) []int {
    dict := make(map[int]int)
    for i, v := range nums {
        if index, ok := dict[target - v]; ok {
            return []int{index, i}
        }
        dict[v] = i
    }

    return []int{}
}