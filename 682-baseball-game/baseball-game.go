// что требуется:
//   - проитерироваться по массиву, выполняя заданные операции, после чего вернуть сумму эллементов
// как решить:
//   - очевидно будет удобен switch для ветвлений по операциям + необходим доступ ко всем элементам
// итог:
//   - мб стоит переписать на type Stack sturct с методами push/pop? Базовая/простая задача, тоже можно на собес

func calPoints(operations []string) int {
    s := make([]int, 0, len(operations))

    for _, v := range operations {
        switch v {
            case "D":
                s = append(s, s[len(s)-1] * 2)
            case "C":
                s = s[:len(s)-1]
            case "+":
                s = append(s, s[len(s)-1] + s[len(s)-2])
            default:
                num, _ := strconv.Atoi(v)
                s = append(s, num)
        }
    }

    var res int
    for _, v := range s {
        res += v
    }

    return res
}

// получать 2 последних значения
// получать последние значение
// удалять последние значение