//函数数组中copy函数的使用
package main
import "fmt"

func rotate(nums []int, k int)  {
	var  index int
	a := len(nums)
	b := make([]int,a)
	for i:=0;i<a;i++{
	  
	index = (i+k)%a
	b[index]=nums[i]
	}
	copy(nums,b)

}

func main() {

   b :=[]int{1,2,3,4,5,6,7}
   rotate(b,3)
   fmt.Println(b)

}
