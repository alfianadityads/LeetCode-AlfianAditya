func runningSum(nums []int) []int {
    
    runningSum := make([]int, len(nums))
    
    sum := 0
    
    for i, val := range nums {
        sum += val
        runningSum[i] = sum
    }
    
    return runningSum
}