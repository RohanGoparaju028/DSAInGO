package sorting;
func InsertionSort(a []int) []int {
	n := len(a)
	if n == 0 {
		return a
	}
	for i:=1;i<n;i++ {
		key := a[i]
		j := i-1
		for j>=0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key 
	}
	return a
}