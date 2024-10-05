package main
import"fmt"

func isLeapYear(y int)bool{
	return (y%4==0&&y%100!=0)||(y%400==0)
}


func main(){
	var x,y int
	var cnt int
	fmt.Scan(&x,&y)
	cnt=0
	for i:=x;i<=y;i++{
		if(isLeapYear(i)){
            cnt++;
		}
	}
	fmt.Print(cnt)
	fmt.Println()
	for i:=x;i<=y;i++{
		if(isLeapYear(i)){
			fmt.Print(i," ")		
		}
	}
}