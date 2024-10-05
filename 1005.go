package main
import"fmt"

func deleteSlice(s[]int)[]int{
	for i:=0;i<len(s);i++{
		if s[i]%3==0 {
			s=append(s[:i],s[i+1:]...)
			i--
		}
	}
	return s;
}

func main(){
	var intarr[100]int
	for i:=range intarr{
		intarr[i]=i
	}
	slice:=intarr[1:51]
	fmt.Println("slice:",slice)
	slice=deleteSlice(slice)
	slice=append(slice,114514)
	fmt.Println("slice:",slice)
	
}
