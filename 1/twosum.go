package main

//两数之和
func main() {
	twoSum([]int{2,7,11,15},9)
}

func twoSum(nums []int,target int)[]int{
   for i:=0;i<len(nums);i++{
   	   for j:=i+1;j<len(nums)-1;j++{
   	  	  if nums[i] + nums[j] == target{
   	  	  	return []int{i,j}
		  }
	   }
   }
   return nil
}