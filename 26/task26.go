package main26

import (
	"fmt"
	"strings"
)

func Main26() {
	m:=make(map[int32]bool)
	str:="abCdefAaf"
	str=fmt.Sprint(strings.ToLower(str))
	for _, sumb := range str {
		if m[sumb]{
			fmt.Println(false)
			return
		}else{
			m[sumb]=true
		}
	}
	fmt.Println(true)
}