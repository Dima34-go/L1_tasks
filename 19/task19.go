package main19

import (
	"fmt"
	"math/rand"
)

func Main19(){
	//strR:=[]rune("olkekcheburek")
	strR:=make([]rune,1<<5)
	for i := 0; i < len(strR); i++ {
		strR[i]=int32(50+rand.Int()%27)
	}
	fmt.Println(string(strR))
	//симметрично отражаем
	for i,j:=0,len(strR)-1;i<j;i,j=i+1,j-1{
		strR[i],strR[j]=strR[j],strR[i]
	}
	fmt.Println(string(strR))
}