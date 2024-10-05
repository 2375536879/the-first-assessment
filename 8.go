package main
import(
    "fmt"
)
func twoSum(nums []int, target int) []int {
    m:=make(map[int]int)
    for i,v:=range nums{
        if _,exist :=m[target-v];exist{
            return []int{i,m[target-v]}
        }
        m[v]=i
    }
    return nil
    }
    
func reverseSlice(nums []int) []int {
    reversed := make([]int, len(nums))
    for i, v := range nums {
        reversed[len(nums)-1-i] = v
    }
    return reversed
}
func main(){
    //输入数组长度
    var n int 
    fmt.Scan(&n)
    nums:=make([]int,n)

    for i:=0;i<n;i++{
        fmt.Scan(&nums[i])
    }

    var target int
    fmt.Scan(&target)

    result:=twoSum(nums,target)
    result=reverseSlice(result)
    fmt.Printf("%v",result)
    
}