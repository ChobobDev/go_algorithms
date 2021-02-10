/*
문제 설명
문자열 s에 나타나는 문자를 큰것부터 작은 순으로 정렬해 새로운 문자열을 리턴하는 함수, solution을 완성해주세요.
s는 영문 대소문자로만 구성되어 있으며, 대문자는 소문자보다 작은 것으로 간주합니다.

제한 사항
str은 길이 1 이상인 문자열입니다.
*/
import (
    "strings"
    "unicode"
    "sort"
)

func solution(s string) string {
    var splits = strings.Split(s, "")
    var upper_case []string
    var lower_case []string
    for _,word := range splits{
        word_rune := []rune(word)
        if (unicode.IsUpper(word_rune[0])== true){
            upper_case = append(upper_case,word)
        }else{
            lower_case = append(lower_case,word)   
        } 
    }
    sort.Sort(sort.Reverse(sort.StringSlice(lower_case)))
    sort.Sort(sort.Reverse(sort.StringSlice(upper_case)))
    lower_result:=strings.Join(lower_case,"")
    upper_result:=strings.Join(upper_case,"")
    final_result:=lower_result+upper_result
   
    return final_result
}