package main

import "fmt"

type queueElem struct {
	value int
	next  *queueElem
}

func main() {
	queue := setup()
	values := []int{1, 2, 3, 4, 5}

	for _, value := range values {
		queue = push(queue, value)
	}

	i := queue
	for i != nil {
		fmt.Print("|", i.value, "| -> ")
		i = i.next
	}
	fmt.Print("nil ")

	queue, val := pop(queue)
	fmt.Println("\npoped value", val)
	i = queue
	for i != nil {
		fmt.Print("|", i.value, "| -> ")
		i = i.next
	}
	fmt.Print("nil ")

	searchVal := 2
	if search(queue, searchVal) {
		fmt.Println("\nvalue",searchVal,"exists in the queue")
	} else {
		fmt.Println("\nvalue",searchVal,"does not exist in the queue")
	}
}

func setup() *queueElem {
	return nil
}

func push(queue *queueElem, value int) *queueElem {
	newElem := &queueElem{
		value: value,
		next:  nil,
	}

	if queue == nil {
		return newElem
	}

	current := queue
	for current.next != nil {
		current = current.next
	}

	current.next = newElem

	return queue
}

func pop(queue *queueElem) (*queueElem, int) {
	if queue == nil {
		fmt.Println("Queue is empty")
		return nil, 0
	}
	return queue.next, queue.value
}

func isEmpty(queue *queueElem) bool {
	return queue == nil
}

func search(queue *queueElem, value int) bool {
	tmpQueue := setup()
	found := false

	for queue != nil {
		var val int
		queue, val = pop(queue)

		if val == value {
			found = true
			break
		}

		push(tmpQueue, val)
	}

	for tmpQueue != nil {
		var tmp int
		tmpQueue, tmp = pop(tmpQueue)
		queue = push(queue, tmp)
	}

	return found
}
