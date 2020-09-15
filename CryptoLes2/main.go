// "Осфшнсақмырсртрыөсдовөыаыыіқзіынознднрнмгзбылааыаудәеқіңдрлушетрарсаамыиаіілаироа"
package main

import "fmt"

func calc(x []rune, pos, cnt int) ([]rune, []rune) {
	var cur, tmp []rune
	for i := range x {
		if (i+1)%pos == 0 && cnt != 0 {
			cur = append(cur, x[i])
			cnt--
		} else {
			tmp = append(tmp, x[i])
		}
	}
	return tmp, cur
}

// Дешифратор
func main() {
	arr := []rune("рбоецаиерезтсд")
	for key := 2; key < len(arr); key++ {
		tmp := ""
		x := append([]rune{}, arr...)
		ost := []rune{}
		x, ost = calc(x, len(arr)/(len(arr)/key)+1, len(arr)%key)
		for i := 0; i < key; i++ {
			for j := i; j < len(x); j += (key) {
				tmp += string(x[j])
			}
		}
		tmp += string(ost)
		fmt.Print("key - ")
		fmt.Println((len(arr) / key))
		fmt.Print("text - ")
		fmt.Println(tmp)
	}
}
