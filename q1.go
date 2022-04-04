package main

//func main() {
//	N:=1<<25
//	arr:=make([]int,N)
//	for i:=range arr{
//		arr[i]=i
//	}
//	//эффективная конкатенация
//	ts := time.Now()
//	var b strings.Builder
//	//без Grow работает с такой же скоростью как через string.Join
//	b.Grow(10*N)
//	for i:=range arr {
//		fmt.Fprintf(&b," %d",arr[i])
//	}
//	str:=b.String()
//	fmt.Println(str)
//	td1:=time.Now().Sub(ts)
//	////
//	// конкатенация через strings.Join
//	ts = time.Now()
//	sa:=make([]string,N)
//	for i:=range sa{
//		sa[i]=fmt.Sprintf(" %d",arr[i])
//	}
//	str1:=strings.Join(sa,"")
//	fmt.Println(str1)
//	td2:=time.Now().Sub(ts)
//	//  конкатенация через Buffer
//	ts = time.Now()
//	var buf bytes.Buffer
//	for i:=range arr {
//		fmt.Fprintf(&buf," %d",arr[i])
//	}
//	str2:=buf.String()
//	fmt.Println(str2)
//	td3:=time.Now().Sub(ts)
//	fmt.Println("Эффективная конкатенация:",td1.Milliseconds())
//	fmt.Println("Конкатенация через Join:",td2.Milliseconds())
//	fmt.Println("Конкатенация через Buffer:",td3.Milliseconds())
//}