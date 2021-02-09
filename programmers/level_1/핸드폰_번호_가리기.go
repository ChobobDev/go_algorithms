func solution(phone_number string) string {
    back_number:=phone_number[len(phone_number)-4:]
    var result string
    for i := 0; i < (len(phone_number)-4); i++ {
          result +="*"
    }
    result+=back_number
    return result
}