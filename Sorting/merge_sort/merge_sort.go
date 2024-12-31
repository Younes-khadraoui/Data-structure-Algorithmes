package main

import "fmt"

func main() {
	array := []int{1, 4, 2, 3}
	sorted_array := merge_sort(array)
	fmt.Println(sorted_array)
}

func merge_sort(array []int) []int {
	N := len(array)
	if  N == 1 {
		return array
	}

	arrayOne := array[0:N/2]
	arrayTwo := array[N/2:]

	merge_sort(arrayOne)
	merge_sort(arrayTwo)

	return merge(arrayOne,arrayTwo)
}

func merge(arrayOne []int, arrayTwo []int) []int {
	arrayThree := []int{}
	N1 := len(arrayOne)
	N2 := len(arrayTwo)
	i,j := 0, 0

	for i<N1 && j<N2 {
		if arrayOne[i] < arrayTwo[j] {
			arrayThree = append(arrayThree, arrayOne[i])
			i ++
		} else {
			arrayThree = append(arrayThree, arrayTwo[j])
			j ++
		}
	}

	if i<N1 {
		arrayThree = append(arrayThree, arrayOne[i:]...)
	}else {
		arrayThree = append(arrayThree, arrayTwo[j:]...)
	}

	return arrayThree
}
