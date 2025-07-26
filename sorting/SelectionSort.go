package sorting;
func SelectionSort(a []int) []int {
	 n := len(a)
	if n == 0 {
		return a;
	}
	for i:=0;i<n;i++ {
       for j:=i+1;j<n;j++ {
		  if a[i] > a[j] {
			a[i],a[j] = a[j],a[i];
		  }
	   }
	}
	return a;
}