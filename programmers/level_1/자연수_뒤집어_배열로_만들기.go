/*
문제 설명
자연수 n을 뒤집어 각 자리 숫자를 원소로 가지는 배열 형태로 리턴해주세요. 예를들어 n이 12345이면 [5,4,3,2,1]을 리턴합니다.

제한 조건
n은 10,000,000,000이하인 자연수입니다.
*/

import "strconv"
import "strings"
func solution(n int64) []int {
    var result []int
    s := strconv.Itoa(int(n))
    var splits = strings.Split(s, "")
    for i := 1; i <= len(s); i++{
        num,_ :=strconv.Atoi(splits[len(splits)-i])
        result = append(result,num)  
    } 
    return result
}