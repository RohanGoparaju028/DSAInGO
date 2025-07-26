package sorting;
func BubbleSort(a []int) []int {
	n := len(a)
	if n == 0 {
		return a;
	}
	for i:=0;i<n;i++ {
		for j:=0;j<n-1-i;j++ {
			if a[j] > a[j+1] {
				a[j],a[j+1] = a[j+1],a[j];
			}
		}
	}
	return a;
}