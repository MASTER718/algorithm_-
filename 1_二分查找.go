//二分查找找到需要的数字


package main
import"fmt"
func search(nums []int,target int)int{
  
    var left,right int
	left =0
	right=len(nums)-1

	for left<=right {
      var mid int
      mid =(left+right)/2

	  if nums[mid]==target {
 
           return mid        

	  } else if nums[mid]<target {
               left = mid+1
        
	  } else if nums[mid]>target{

           right = mid -1

	  }
	  
	}
	return -1
	

}


func main(){
fmt.Println(search([]int{3,5,7,9,23},23))
 




}
