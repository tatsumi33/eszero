package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	s, argexst := codeStr()
	c := 0
	if argexst {
		var err error
		if c, err = toReturnInt(s); err != nil {
			fmt.Fprintln(os.Stderr, "Parse error:", s)
			os.Exit(0)
			return
		}
	}

	fmt.Println(c)
	os.Exit(c)
}

func codeStr() (s string, argexst bool) {
	if len(os.Args) == 1 {
		b := filepath.Base(os.Args[0])
		b = b[:len(b)-len(filepath.Ext(b))]
		if b == "eszero" {
			return b, false
		}

		s = b[:len(b)-len(filepath.Ext(b))]
		if p := strings.LastIndex(s, "_"); p >= 0 {
			s = s[p+1:]
		}
	} else {
		s = os.Args[1]
	}
	return s, true
}

func toReturnInt(s string) (int, error) {
	t, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return 0, err
	}
	return int(t), nil
}
