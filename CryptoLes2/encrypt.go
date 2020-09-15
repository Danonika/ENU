package main

import "fmt"

var a [][]rune

func main() {
	arr := []rune("ЯЛЮБЛЮАНИМЕ")
	key := 3
	cur := 0
	cnt := 0
	a = append(a, []rune{})
	for i := range arr {
		if cur == key {
			a = append(a, []rune{})
			cnt++
			cur = 0
		}
		a[cnt] = append(a[cnt], arr[i])
		cur++

	}
	ans := []rune{}
	for i := 0; i < key; i++ {
		for j := range a {
			if len(a[j]) > i {
				ans = append(ans, a[j][i])
			}
		}
	}
	fmt.Println(string(ans))
}

/*
ЖЕЛСІЗ ТҮНДЕ ЖАРЫҚ АЙ СӘУЛЕСІ СУДА ДІРІЛДЕП АУЫЛДЫҢ ЖАНЫ ТЕРЕҢ САЙ ТАСЫҒАН ӨЗЕН КҮРІЛДЕПЕ
*/
