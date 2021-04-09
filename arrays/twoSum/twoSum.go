package arrays

func twoSum(nums []int, target int) []int {
    if len(nums) == 2 {
        return []int{0,1}
    } 

	var sumsMap = make(map[int]int)
    
    for i,val := range nums {
        complement := target - val
        
        if mapVal,exists := sumsMap[complement]; exists {
			test := []int{mapVal,i}
            return test
        }
        
        sumsMap[val] = i
    }
    
    return []int{-1}
}