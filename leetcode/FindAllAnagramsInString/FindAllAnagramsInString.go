package FindAllAnagramsInString

func findAnagrams(s, p string) (ans []int) {        
    sLen, pLen := len(s), len(p)        
    if sLen < pLen {        
        return        
    }        
        
    var sCount, pCount [26]int        
    for i, ch := range p {        
        sCount[s[i]-'a']++        
        pCount[ch-'a']++        
    }        
    if sCount == pCount {        
        ans = append(ans, 0)        
    }        
        
    for i, ch := range s[:sLen-pLen] { //滑动窗口根据字符串长度依次遍历S的子序列        
        sCount[ch-'a']--        
        sCount[s[i+pLen]-'a']++        
        if sCount == pCount {        
            ans = append(ans, i+1)        
        }        
    }        
    return        
}        
        