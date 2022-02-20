//1.快慢指针
// 定义slow，fast 头部添加dummy指针，俩者相差n，也就是我们要删除的倒数第几个元素
package main

type  list struct  {
    val int
	next *list
}

func remove_1(head *list,n int) *list {
    dummy := &list{0,head}                       // 添加虚拟头节点
	slow,fast := dummy,dummy                     // 快慢指针赋值

	for i:=0;i<=n;i++ {                          //快指针先走N+1步
    fast = fast.next		
	} 


	for  fast!= nil {                           //快慢指针均往前移动，直到fast到达末尾nil
 
      slow = slow.next
	  fast = fast.next
	}

	delnode     := slow.next                           //待删除节点
    slow.next    = delnode.next                        //改变指向
	delnode.next = nil                                 //删除节点


	return   dummy.next

}

func main() {
}




//  2.采用栈的方法    先进后出


func remove_2(head *list,n int) *list {

        stack  := []*list{}                //定义一个栈           
        dummy := &list{0,head}             //头部虚指针

		for   node:=dummy;node!=nil;node = node.next {    //压入栈内
			stack = append(stack,node)
		}

		// 处理节点
		pre_delcode := stack[len(stack)-n-1]             //找到待删除节点的前一个节点
		delcode     := pre_delcode.next
		pre_delcode.next = delcode.next
		delcode.next = nil





		return  dummy.next



}