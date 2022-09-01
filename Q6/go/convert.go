package convert

// 8ms (77.56%), 7MB (49.57%), 2022/3/11 16:54
func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    
    var(
        jump  []int = []int{(numRows - 1) * 2, 0}
        length  int = len(s)
        ans  string = ""
    )
    
    for i := 0; i < numRows; i++ {
        for j, index := i, 0; j < length; index = (index + 1) % 2 {
            if jump[index] != 0 {
                ans = ans + string(s[j])
                j = j + jump[index]
            }
        }
        
        jump[0] = jump[0] - 2
        jump[1] = jump[1] + 2
    }
    
    return ans
}
