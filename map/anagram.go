package main

import "fmt"

//比较两个单词是否为异位词
func main() {
	fmt.Println(isAnagram("dsda","dsad"))
}

func isAnagram(s string, t string)bool {
	dic1:=make(map[int]int)
	for _,v:=range []rune(s){
		key1:=int(v)-32
		dic1[key1] += 1
	}
	for _,v2:=range []rune(t){
		key2:=int(v2)-32
		val,ok:=dic1[key2]
		if ok{
			if val>0{
				dic1[key2] -=1
			}else {
				return false;
			}

		}else{
			return  false
		}
	}
	for _,val:=range dic1{
		if val>0{
			return  false
		}

	}
	return  true
}
//大神算法
func isAnagram2(s string, t string) bool {
	var charCount [26]int
	for _, c := range s {
		// charCount[c - 97]++
		charCount[c - 'a']++
	}
	for _, c := range t {
		//charCount[c - 97]--
		charCount[c - 'a']--
	}
	for idx, _ := range charCount {
		if charCount[idx] != 0 {
			return false
		}
	}
	return true
}