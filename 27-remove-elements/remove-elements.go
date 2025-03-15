// что требуется:
//   - удалить из слайса все вхождения target
//   - расположить оставшиеся элементы слайса относительно левого края; порядок не важен
//   - вернуть len(input) - len(deletedTarget)
// как решить:
//   - проходишься по массиву, если v == target -> меняешь на <flag> -> slices.DeleteFunc() {if flag -> true}
//     - O(n), подходит, но можно лучше
//   - завести переменную indexToModify и менять значения по индексу currVal
// итог:
//   - базовая задачка, над которой можно подумать без знания алгосов -> для алго-собесов супер

func removeElement(nums []int, target int) int {
	indexToModify := 0

	for i, v := range nums {
		if v != target {
			nums[indexToModify] = nums[i]
			indexToModify++
		}
	}

	return indexToModify
}
