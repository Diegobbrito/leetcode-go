func arraySign(nums []int) int {
    i := 1
    for _, num := range (nums){
        if(num == 0) {
            return 0
        } 
        if(num < 0){
            i *= -1
        } 
    }
    return i
}
