package sorting;
func MergeSort(a []int) []int {
	if len(a) <= 1 {
		return a;
	}
	mid := len(a)/2
	left := MergeSort(a[:mid])
	right := MergeSort(a[mid:])
	return merge(a,left,right);
}
func merge(a,left,right []int) []int {
	leftSize,rightSize := len(left),len(right)
	i,j,k := 0,0,0
	for i<leftSize && j<rightSize {
		if left[i] <= right[j] {
			a[k] = left[i]
			i++
		} else {
			a[k] = right[j]
			j++
		}
		k++
	}
	for i < leftSize {
		a[k] = left[i]
		i++
		k++
	}
	for j < rightSize {
		a[k] = right[j]
		j++
		k++
	}
	return a;
}