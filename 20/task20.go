package main20

import (
	"fmt"
	"strings"
)

func Main20() {
	str:="snow dog sun booo"
	fmt.Println("start str:",str)
	//разделили строки на массив строк
	strS:=strings.Split(str," ")
	for i,j:=0,len(strS)-1;i<j;i,j=i+1,j-1 {
		strS[i],strS[j]=strS[j],strS[i]
	}
	str=strings.Join(strS," ")
	fmt.Println("finish str:",str)
}