package main22

import (
	"fmt"
	"math/big"
)

func Main22()  {
	bi1,bi2:=new(big.Int),new(big.Int)
	res:=new(big.Int)
	var strbi1, strbi2 string
	fmt.Scan(&strbi1,&strbi2)
	_,ok1:=bi1.SetString(strbi1,10)
	_,ok2:=bi2.SetString(strbi2,10)
	if!(ok1 && ok2){
		fmt.Println("Введено некорректное число")
		return
	}
	fmt.Printf("Cумма чисел: %s + %s = %s \n",bi1,bi2,res.Add(bi1,bi2))
	fmt.Printf("Разность чисел: %s - %s = %s \n",bi1,bi2,res.Sub(bi1,bi2))
	fmt.Printf("Произведение чисел: %s * %s = %s \n",bi1,bi2,res.Mul(bi1,bi2))
	fmt.Printf("Деление чисел: %s / %s = %s \n",bi1,bi2,res.Div(bi1,bi2))
}