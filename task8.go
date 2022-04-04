package main

func setBit(val *int,n int, bit int) {
	if bit != 0 && bit != 1{
		return
	}
	if *val & 1<<n == bit<<n{//ничего менять не надо
		return
	}else if bit==1{
		*val=*val-1<<n
	}else{
		*val=*val+1<<n
	}
}
//func main(){
//	var val = 270
//	fmt.Printf("%b \n",val)
//	setBit(&val,3,1)
//	fmt.Printf("%b \n",val)
//}