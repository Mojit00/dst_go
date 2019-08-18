package sortfunction

import (
	"fmt"
)

type SortUtil struct {
	Sortable []int
}

func NewSortUtil(arr []int) *SortUtil {
	return &SortUtil{
		Sortable: arr,
	}
}

func (s *SortUtil) PopSort() {
	lag := false
	tempArr := s.Sortable
	var temp int = 0
	for i := 0; i < len(s.Sortable)-1; i++ {
		for j := 0; j < len(s.Sortable)-1-i; j++ {
			if tempArr[j] > tempArr[j+1] {
				temp = tempArr[j]
				tempArr[j] = tempArr[j+1]
				tempArr[j+1] = temp
				lag = true
			}
		}
		if lag == true {
			lag = false
		} else {
			break
		}
	}
}

func (s *SortUtil) SelectSort() {
	var min int
	for i := 0; i < len(s.Sortable)-1; i++ {
		min = i
		for j := i + 1; j < len(s.Sortable); j++ {
			if s.Sortable[j] < s.Sortable[min] {
				min = j
			}
		}
		if min != i {
			temp := s.Sortable[i]
			s.Sortable[i] = s.Sortable[min]
			s.Sortable[min] = temp
		}
	}
}

func (s *SortUtil) InsertSort() {
	arr := s.Sortable
	for i := 1; i < len(arr); i++ {
		insertValue := arr[i] //定义待插入的数字
		insertIndex := i - 1  //当前元素前面的下表个数
		//没有找到就需要后移元素
		for insertIndex >= 0 && insertValue < arr[insertIndex] {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertValue
		}
	}
}

func (s *SortUtil) QuickSort() {
	quickSort(s.Sortable, 0, len(s.Sortable)-1)
}

func quickSort(arr []int, start, end int) {
	if start < end {

		standard := arr[start]
		low := start
		high := end

		for low < high {
			for low < high && standard <= arr[high] {
				high--
			}
			arr[low] = high

			for low < high && standard >= arr[low] {
				low++
			}
			arr[high] = arr[low]
		}
		arr[low] = standard

		quickSort(arr, start, low)
		quickSort(arr, low+1, end)
	}
}

func (s *SortUtil) print() {
	fmt.Print("[")
	for _, v := range s.Sortable {
		fmt.Print(v, " ")
	}
	fmt.Println("]")
}
