/*
문제 설명
함수 solution은 정수 n을 매개변수로 입력받습니다. n의 각 자릿수를 큰것부터 작은 순으로 정렬한 새로운 정수를 리턴해주세요. 예를들어 n이 118372면 873211을 리턴하면 됩니다.

제한 조건
n은 1이상 8000000000 이하인 자연수입니다.
*/

import "strconv"
import "sort"
import "strings"
func solution(n int64) int64 {
    var result []string
    s := strconv.Itoa(int(n))
    var splits = strings.Split(s, "")
    for _,num := range splits {
        result=append(result,num)
    }
    sort.Sort(sort.Reverse(sort.StringSlice(result)))
    result_str:=strings.Join(result,"")
    result_int,_ := strconv.Atoi(result_str)
    return int64(result_int)
}
