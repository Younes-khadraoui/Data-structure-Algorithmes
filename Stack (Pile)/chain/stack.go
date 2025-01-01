package main

import "fmt"

type stackElem struct {
	value int
	next  *stackElem
}

func main() {
	stack := setup()
	values := []int{1, 2, 3, 4, 5}

	for _, value := range values {
		stack = push(stack, value)
	}

	i := stack
	for i != nil {
		fmt.Print("|", i.value, "| -> ")
		i = i.next
	}
	fmt.Print("nil ")

	stack, val := pop(stack)
	fmt.Println("\npoped value", val)
	i = stack
	for i != nil {
		fmt.Print("|", i.value, "| -> ")
		i = i.next
	}
	fmt.Print("nil ")

	searchVal := 4
	if search(stack, searchVal) {
		fmt.Println("\nvalue",searchVal,"exists in the stack")
	} else {
		fmt.Println("\nvalue",searchVal,"does not exist in the stack")
	}
}

func setup() *stackElem {
	return nil
}

func push(stack *stackElem, value int) *stackElem {
	return &stackElem{
		value: value,
		next:  stack,
	}
}

func pop(stack *stackElem) (*stackElem, int) {
	if stack == nil {
		fmt.Println("Stack is empty")
		return nil, 0
	}
	return stack.next, stack.value
}

func isEmpty(stack *stackElem) bool {
	return stack == nil
}

func search(stack *stackElem, value int) bool {
	tmpStack := setup()
	found := false

	for stack != nil {
		var val int
		stack, val = pop(stack)

		if val == value {
			found = true
			break
		}

		tmpStack = push(tmpStack, val)
	}

	for tmpStack != nil {
		var tmp int
		tmpStack, tmp = pop(tmpStack)
		stack = push(stack, tmp)
	}

	return found
}
