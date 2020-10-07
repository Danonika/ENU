package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	answer, len, x, y int
	seq               []rune
)

func main() {

	st := map[string]int{}
	fmt.Println("Do you wanna single check(1) or multi check(2) or work(3)? ")
	fmt.Scan(&answer)
	j := 1
	if answer == 2 {
		j = 100
	}
	rand.Seed(time.Now().UnixNano())
	fmt.Print("Let's check generator with length - ")
	fmt.Scan(&len)
	fmt.Printf("Please Enter two position which will be XOR(ed) ")
	fmt.Scan(&x, &y)
	for k := 1; k <= j; k++ {
		seq = []rune{}
		for i := 0; i < len; i++ {
			if answer != 3 { // Random Sequence for testing
				tx := rand.Intn(2)
				seq = append(seq, rune(tx+'0'))
			} else {
				seq = append(seq, '1') // Just sequence
			}
		}
		if answer == 1 || answer == 3 {
			fmt.Printf("Random sequence with length - %d is %s\n", len, string(seq))
		}
		if answer == 1 || answer == 2 { // For testing sequences
			i := 1
			mx := 0
			first_seq := []rune{}
			first_seq = append(first_seq, seq...)
			for {
				if i == 900000 { // Decrease this number to increase perfomance
					break
				}
				val, ok := st[string(seq)]
				if !ok {
					st[string(seq)] = i
					val = -1
				}
				if val > mx {
					mx = val
				}
				if answer == 1 {
					fmt.Printf("Iteration %d - %s - First entrance %d\n", i, string(seq), val)
				}
				seq = append(seq, rune(seq[x]^seq[y]+'0'))
				seq = seq[1:]
				i++
			}
			if answer == 2 {
				fmt.Printf("Case #%d: %d-length Sequence %s has %d period\n", k, len, string(first_seq), mx)
			}
		} else {
			answer2 := 1
			answer3 := 1
			fmt.Print("Choose Async(1) or Sync(2)? ")
			fmt.Scan(&answer2)
			fmt.Print("Choose decrypt(1) or encrypt(2)? ")
			fmt.Scan(&answer3)
			fl, _ := os.Open("input.txt")
			if answer3 == 1 {
				fl, _ = os.Open("output.txt")
			}
			scanner := bufio.NewScanner(fl) // Scanner
			output := ""

			for scanner.Scan() { // Scan if u can
				fmt.Println(scanner.Text())   // Get line from scanner
				arr := []rune(scanner.Text()) // Transfer to array of rune
				for _, i := range arr {       // Get rune one by one
					num := fmt.Sprintf("%08b", i) // Transfer rune to binary number
					ans := ""

					for _, j := range num { // Get one by one bit from binary number
						tmp := rune(seq[x] ^ seq[y] + '0')
						if answer2 == 1 { // Async ?
							tmp = rune(tmp ^ j + '0')
							if answer3 == 1 { // Decrypt ?
								seq = append(seq, j)
							} else { // Encrypt ?
								seq = append(seq, tmp)
							}
							seq = seq[1:]
						} else { // Sync
							seq = append(seq, tmp)
							seq = seq[1:] // Cut first
							tmp = rune(tmp ^ j + '0')
						}
						ans += string(tmp)
					}
					ch, _ := strconv.ParseInt(ans, 2, 64) // Binary number to decimal
					output += string(ch)                  // Decimal to ASCII rune
					if ch < 32 {                          // Skip if operational code
						ch = ' '
					}
					fmt.Printf("%c -> %c\n", i, ch) // Print result of transformation
				}
				fl2, _ := os.Create("output.txt")
				fl2.Write([]byte(output))
			}
		}
	}
}
