package main

import (
	"os"
	"log"
	"math/rand"
	"strconv"
	"github.com/hoisie/web"
	"net/http"
	"time"
)

var (
	valid_chars = []rune{'a', 'b', 'c', 'd', 'e', 'f', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	valid_strings = []string{"a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func awaitQuitKey() {
	var buf [1]byte
	for {
		_, err := os.Stdin.Read(buf[:])
		if err != nil || buf[0] == 'q' {
			return
		}
	}
}

func getPassword(countstr string) string {
	t := time.Now()

	count, _ := strconv.ParseInt(countstr, 10, 64)
	var i int64

	passwd := make([]rune, count)

	for i = 0; i < count; i++ {
		passwd[i] = valid_chars[r.Intn(len(valid_chars))]
	}

	log.Printf("Password generated in: %v", time.Now().Sub(t))
	return string(passwd)
}

func getPassword_B(ctx *web.Context, countstr string) {
	flusher, _ := ctx.ResponseWriter.(http.Flusher)
	flusher.Flush()

	count, _ := strconv.ParseInt(countstr, 10, 64)
	var i int64

	for i = 0; i < count; i++ {
		ctx.WriteString(valid_strings[r.Intn(len(valid_strings))])
		flusher.Flush()
	}
}

func main () {
	w := web.NewServer()

	w.Get("/([0-9]+)", getPassword_B)

	go w.Run(":8002")
	awaitQuitKey()
	log.Print("Server exiting gracefully...")
}
