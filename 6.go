package main
import(
    "fmt"
    "os"
    "bufio"
)

func main(){
    file,err:=os.OpenFile("d:/ninenine.txt",os.O_RDWR|os.O_APPEND|os.O_CREATE,114514)
    
    if(err!=nil){
        fmt.Printf("文件打开出错 err:",err)
    }

    defer file.Close()

   writer:=bufio.NewWriter(file)


   for i:=1;i<=9;i++{
    for j:=1;j<=i;j++{
        str:=fmt.Sprintf("  %v * %v = %v  ",j,i,i*j)
        writer.WriteString(str)
    }
     writer.WriteString("\n")
   }
   writer.Flush()
}