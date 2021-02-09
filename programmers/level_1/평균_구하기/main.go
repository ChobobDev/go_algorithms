func solution(arr []int) float64 {
    sum := 0  
    for _, num := range arr {  
        sum += num
    }
    res := (float64(sum)) / (float64(len(arr))) 
    
    return float64(res)
}