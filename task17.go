package main

//func main(){
//	arr:=make([]int,1<<6)
//	for i := 0; i < len(arr); i++ {
//		arr[i]=rand.Int()%(1<<8)
//	}
//	val:=arr[rand.Int()%(len(arr)-1)]
//	sort.Slice(arr, func(i, j int) bool {
//		return arr[i] < arr[j]
//	})
//	fmt.Println(arr)
//	i:=binarySearch(val,arr)
//	fmt.Println("искомое значение: ",val)
//	fmt.Println("найденная позиция: ",i)
//	fmt.Println("значение на данной позиции: ",arr[i])
//}
func binarySearch(val int, arr[]int) int{
	l,r:=0,len(arr)-1
	if arr[l] > val || arr[r] < val{
		return -1
	}
	for l < r {
		m:=(l+r)/2
		if arr[m] < val {
			l=m
		}else if arr[m] > val {
			r=m
		}else{
			return m
		}
	}
	return -1
}