// не лучшее решение, конечно

type MinStack struct {
	stack    []int
	minIndex []int
}

func Constructor() MinStack {
	return MinStack{stack: make([]int, 0), minIndex: make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.minIndex) == 0 {
		this.minIndex = append(this.minIndex, len(this.stack)-1)
	} else if val <= this.stack[this.minIndex[len(this.minIndex)-1]] {
		this.minIndex = append(this.minIndex, len(this.stack)-1)
	}
}

func (this *MinStack) Pop() {

	if len(this.minIndex) > 0 && len(this.stack)-1 == this.minIndex[len(this.minIndex)-1] {
		this.minIndex = this.minIndex[:len(this.minIndex)-1]
	}

	this.stack = this.stack[:len(this.stack)-1]

}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}

	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minIndex) > 0 {
		return this.stack[this.minIndex[len(this.minIndex)-1]]
	}

	return 0
}

