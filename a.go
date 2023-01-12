package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		a = 1
		b = 0
		c = ""
		d = "something"
	)

	/*
	 * 書き方a
	 */
	if forRun1a(a, c, d) {
		if err := Run1(); err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(-1)
		}
	}
	if forRun2a(b, c, d) {
		if err := Run2(); err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(-1)
		}
	}
	/* ここまで */

	/*
	 * 書き方b
	 */
	if forRun1b(a, c, d) {
		if err := Run1(); err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(-1)
		}
	}
	if forRun2b(b, c, d) {
		if err := Run2(); err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(-1)
		}
	}
	/* ここまで */

	/*
	 * 書き方c
	 */
	if isPositive(a) && eitherNotBlank(c, d) {
		if err := Run1(); err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(-1)
		}
	}
	if isZero(b) && eitherNotBlank(c, d) {
		if err := Run2(); err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(-1)
		}
	}
	/* ここまで */

	os.Exit(0)
}
func forRun1a(a int, c, d string) bool { return a > 0 && (c != "" || d != "") }
func forRun2a(b int, c, d string) bool { return b == 0 && (c != "" || d != "") }
func forRun1b(a int, c, d string) bool { return isPositive(a) && eitherNotBlank(c, d) }
func forRun2b(b int, c, d string) bool { return isZero(b) && eitherNotBlank(c, d) }
func isPositive(p int) bool            { return p > 0 }
func isZero(p int) bool                { return p == 0 }
func isBlank(p string) bool            { return p == "" }
func eitherNotBlank(a, b string) bool  { return !isBlank(a) || !isBlank(b) }

// 処理
func Run1() error {
	fmt.Println("**処理1**")
	return nil
}
func Run2() error {
	fmt.Println("**処理2**")
	return nil
}
