//  滑动窗口
// 时间复杂度：O(2n) = O(n)，最坏的情况是 left 和 right 都遍历了一遍字符串
// 空间复杂度：O(n)
//思路：每次都把遍历的元素放到一个map中

//使用make分配内存    不使用的话需要初始化{}   
package main
import "fmt"

func lengthOfLongestSubstring(s string) int {
        
	var  window  = make(map[byte]bool)               //定义是否重复字母
	var  left,right = 0,0                            //左右窗口
	var  lens  int                                    
	for right < len(s)    {
		   r_index := s[right]

		  for window[r_index] == true {          //如果含有相同元素，左边界一直向右移动，直到不包括
			delete(window,s[left])               //删除原左边元素
			left++                               //左边界移动						 
		  }

		  if right-left+1 >=lens {                //赋值长度
				lens = right-left+1
		  }
		  window[r_index] = true
		  right++                                //移动右边界
		

	    }
	    return lens
}

func main() {
    var n = lengthOfLongestSubstring("pwwkew")
     fmt.Println(n)


}