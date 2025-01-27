func calPoints(operations []string) int {
    res := make([]int, 0, len(operations))
    for _, v := range operations {
        res = push(res, v, len(res))
    }

    var resSum int
    for _, v := range res {
        resSum += v
    }

    return resSum
}

func push(stack []int, option string, l int) []int{
    switch option {
        case "+":
            stack = append(stack, stack[l-1] + stack[l - 2])
        case "D":
            stack = append(stack, stack[l-1] * 2)
        case "C":
            stack = stack[:l - 1]
        default:
            i, _ :=  strconv.Atoi(option)
            stack = append(stack, i)
    }

    return stack
}