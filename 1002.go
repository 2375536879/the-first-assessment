package main
import"fmt"
func main(){
var a[10]int

for i:=range a{
	fmt.Scan(&a[i])
}

var h int
fmt.Scan(&h)
h+=30
ans:=0

for i:=range a{
	if h>=a[i]{
		ans++
	}
        
}

fmt.Println(ans)

}