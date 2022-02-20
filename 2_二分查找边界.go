//算法278

package main
import "fmt"

func isbad(a int) bool{
    if a>5{
		return true
	}else {
		return false
	}
    

}


func findbad(n int) int{
   var left,right,mid int
   left =1
   right=n

   for left <right {
	   mid=left+(right-left)/2
       
	   if isbad(mid) == true{
           right =mid
         
	   }else if isbad(mid) == false{
		 left =mid +1
	
	}

   } 
   return left
}
func main(){
var a=24
 
 fmt.Println(findbad(a))


}