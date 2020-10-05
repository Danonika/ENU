package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("En(1) or Ru(2) ? Press 3 to see instructions")
	answer := 1
	fmt.Scan(&answer)
	if answer == 3 {
		fmt.Println("text.txt - our input. On tetx.txt first line is key.\nOther lines are message or text which will be encrypted.\noutput.txt - is final answer or encrypted text")
		os.Exit(0)
	}
	if answer == 1 {
		fmt.Println("You choose English text encryption:")
	}
	if answer == 2 {
		fmt.Println("You choose Russian text encryption:")
	}
	st := map[rune]int{
		'А': 0,
		'Б': 1,
		'В': 2,
		'Г': 3,
		'Д': 4,
		'Е': 5,
		'Ж': 6,
		'З': 7,
		'И': 8,
		'Й': 9,
		'К': 10,
		'Л': 11,
		'М': 12,
		'Н': 13,
		'О': 14,
		'П': 15,
		'Р': 16,
		'С': 17,
		'Т': 18,
		'У': 19,
		'Ф': 20,
		'Х': 21,
		'Ц': 22,
		'Ч': 23,
		'Ш': 24,
		'Щ': 25,
		'Ъ': 26,
		'Ы': 27,
		'Ь': 28,
		'Э': 29,
		'Ю': 30,
		'Я': 31,
	}
	pt := map[int]rune{}
	for i, j := range st {
		pt[j] = i
	}
	f, _ := os.Open("text.txt")
	f2, _ := os.Create("output.txt")
	sc := bufio.NewScanner(f)
	sc.Scan()
	key := sc.Text()
	j := 0
	key = strings.ToUpper(key)
	arr := []rune(key)
	for sc.Scan() {
		i := []rune(sc.Text())
		s := ""
		for _, x := range i {
			tmp := x
			if unicode.IsLetter(x) {
				var cnt int
				if answer == 1 {
					if 'A' <= x && x <= 'Z' {
						cnt = int(x - 'A' + arr[j] - 'A')
						if cnt > 25 {
							cnt -= 26
						}
						tmp = rune('A' + cnt)
					} else {
						cnt = int(x - 'a' + unicode.ToLower(arr[j]) - 'a')
						if cnt > 25 {
							cnt -= 26
						}
						tmp = rune('a' + cnt)
					}
				} else {
					if unicode.IsUpper(x) {
						cnt = int(st[x] + st[arr[j]] + 1)
						if cnt > 31 {
							cnt -= 32
						}
						tmp = pt[cnt]
					} else {
						cnt = int(st[unicode.ToUpper(x)] + st[arr[j]] + 1)
						if cnt > 31 {
							cnt -= 32
						}
						tmp = unicode.ToLower(pt[cnt])
					}
				}

				fmt.Printf("%d -> %c - %c = %c\n", j, arr[j], x, tmp)
				j++
				if j == len(arr) {
					j = 0
				}
			}
			s += string(tmp)
		}
		s += "\n"
		f2.Write([]byte(s))

	}
	fmt.Println("Please open output.txt to see full encrypted text")
}
