package sorting
func QuickSort(a []int) []int {
	return quickSort(a, 0, len(a)-1)
}
func quickSort(a []int, l, h int) []int {
	if l < h {
		p := partition(a, l, h)
		quickSort(a, l, p-1)
		quickSort(a, p+1, h)
	}
	return a
}
func partition(a []int, l, h int) int {
	pivot := a[h]
	i := l
	for j:=l;j<h;j++ {
		if a[j] <= pivot {
			a[i],a[j] = a[j],a[i]
		    i++
		}
	}
	a[i],a[h] = a[h],a[i]
	return i
}