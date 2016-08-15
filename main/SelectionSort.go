package main

import (
	"fmt"
	"time"
)

//in every pass make the currSwapCandidate and magenta point to the first element in the sub-list

func main (){

	fmt.Println("selection sort starting")
	arr:= [5]int{2,1,4,3,5}
	//magentaIndex:=0
	//currSwapCandidateIndex:=0

	fmt.Println("unsorted array: ", arr)
	time.Now()
	//outer loop to iterate on subset of array
	for outerLoop:=0; outerLoop < len(arr); outerLoop++{

		min:=outerLoop // the first element in every outer iteration is assumed minimum

		//inner loop for per-iteration comparisons
		for innerLoop:=outerLoop+1; innerLoop< len(arr); innerLoop++ {
			if arr[innerLoop] < arr[min]{
				min = innerLoop
			}
		}

		swapElements(outerLoop, min, &arr)
	}
	fmt.Println("sorted array: ", arr)
}

func swapElements(magentaIndex, currSwapCandidateIndex int, arr *[5]int){

	temp := arr[currSwapCandidateIndex]
	arr[currSwapCandidateIndex] = arr[magentaIndex]
	arr[magentaIndex] = temp
}
