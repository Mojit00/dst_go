package sortfunction

import "testing"

func TestSortUtil_PopSort(t *testing.T) {
	arr := []int{1, 5, 3, 7, 9, 8, 6, 2, 4}
	sortUtil := NewSortUtil(arr)
	sortUtil.PopSort()
	sortUtil.print()
}

func TestSortUtil_SelectSort(t *testing.T) {
	arr := []int{1, 5, 3, 7, 9, 8, 6, 2, 4}
	sortUtil := NewSortUtil(arr)
	sortUtil.SelectSort()
	sortUtil.print()
}

func TestSortUtil_InsertSort(t *testing.T) {
	arr := []int{1, 5, 3, 7, 9, 8, 6, 2, 4}
	sortUtil := NewSortUtil(arr)
	sortUtil.InsertSort()
	sortUtil.print()
}

func TestSortUtil_QuickSort(t *testing.T) {
	arr := []int{1, 5, 3, 7, 9, 8, 6, 2, 4}
	sortUtil := NewSortUtil(arr)
	sortUtil.QuickSort()
	sortUtil.print()
}
