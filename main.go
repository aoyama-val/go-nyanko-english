package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	words := LoadWords("words.txt")

	sc := bufio.NewScanner(os.Stdin)

	for {
		goto_next := false
		correct, masked := MakeProblem(words)
		Say("「" + masked + "」に当てはまる英単語を答えるニャ")
		for sc.Scan() {
			answer := sc.Text()
			if answer == "正解は？" {
				Say("しょうがにゃいニャ～")
				Say(correct)
				goto_next = true
				break
			} else if answer == correct {
				Say("正解だニャ")
				goto_next = true
				break
			} else {
				Say("不正解だニャ")
			}
		}
		if goto_next == false {
			break
		}
	}
	fmt.Print("bye!\n")
}

func Check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
}

func LoadWords(filename string) []string {
	f, err := os.Open(filename)
	Check(err)
	defer f.Close()
	ret := make([]string, 0, 3000)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		ret = append(ret, sc.Text())
	}
	return ret
}

// 0..(max-1)の配列をシャッフルする
func Shuffle(numbers []int) {
	max := len(numbers)
	for i := 0; i < max-1; i++ {
		j := i + 1 + rand.Intn(max-1-i)
		tmp := numbers[i]
		numbers[i] = numbers[j]
		numbers[j] = tmp
	}
}

func MaskWord(word string) string {
	count := 0
	wordlen := len(word)
	if wordlen <= 5 {
		count = 1
	} else if wordlen <= 8 {
		count = 2
	} else {
		count = 3
	}
	_ = count

	positions := make([]int, wordlen)
	for i := 0; i < wordlen; i++ {
		positions[i] = i
	}
	Shuffle(positions)

	for i := 0; i < count; i++ {
		pos := positions[i]
		word = word[:pos] + "*" + word[pos+1:]
	}

	return word
}

func MakeProblem(words []string) (string, string) {
	i := rand.Intn(len(words))
	w := words[i]
	masked := MaskWord(w)
	return words[i], masked
}

func Say(msg string) {
	fmt.Println("ニャンコ:", msg)
}