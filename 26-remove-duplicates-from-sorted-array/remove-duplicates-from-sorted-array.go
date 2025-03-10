// что требуется: 
//     1. модифицировать nums, чтобы подряд располагались только уникальные числа 
//     2. вернуть кол-во уникальных чисел
// как решить:
//     1. проставить "ключи", по которым slices.DeleteFunc сделает грязную работу
//     2. в новый массив записывать уникальные ключи, после чего пройтись по nums и пересетить их
//     3. самостоятельно написать slices.DeleteFunc, т.к. функционал clear мне не нужеж 
// итог:
//     - выбираю 3, т.к. функция clear в slices.DeleteFunc мне не нужна + на литкоде хочу писать ручками всё

// мне неинтересны все последующие вхождения одинаковых чисел
func removeDuplicates(nums []int) int {
    currNum, indexToModify := nums[0], 1

    for _, v := range nums {
        if v == currNum {
            continue
        }

        currNum = v
        nums[indexToModify] = v
        indexToModify++
    }

    return indexToModify
}