/*
문제 설명
주어진 숫자 중 3개의 수를 더했을 때 소수가 되는 경우의 개수를 구하려고 합니다. 숫자들이 들어있는 배열 nums가 매개변수로 주어질 때, nums에 있는 숫자들 중 서로 다른 3개를 골라 더했을 때 소수가 되는 경우의 개수를 return 하도록 solution 함수를 완성해주세요.

제한사항
nums에 들어있는 숫자의 개수는 3개 이상 50개 이하입니다.
nums의 각 원소는 1 이상 1,000 이하의 자연수이며, 중복된 숫자가 들어있지 않습니다.

*/
import "math/big"

func sum(array []int) int {  
 result := 0  
 for _, v := range array {  
  result += v  
 }  
 return result  
}  

func rPool(p int, n []int, c []int, cc [][]int) [][]int {
    if len(n) == 0 || p <= 0 {
        return cc
    }
    p--
    for i := range n {
        r := make([]int, len(c)+1)
        copy(r, c)
        r[len(r)-1] = n[i]
        if p == 0 {
            cc = append(cc, r)
        }
        cc = rPool(p, n[i+1:], r, cc)
    }
    return cc
}

func Pool(p int, n []int) [][]int {
    return rPool(p, n, nil, nil)
}

func solution(nums []int) int {
    answer := 0
    p := Pool(3, nums)
    for i := range p {
        sum_of_combination :=sum(p[i])
        if big.NewInt(int64(sum_of_combination)).ProbablyPrime(0) {
            answer ++
        } 
    }
    // [실행] 버튼을 누르면 출력 값을 볼 수 있습니다.

    return answer
}