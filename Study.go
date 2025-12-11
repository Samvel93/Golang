package main

import "fmt"

func MaskLink(b string) string {
	out := make([]byte, 0, len(b))
	for i := 0; i < len(b); i++ {
		if i+7 <= len(b) && b[i:i+7] == "http://" {
			out = append(out, "HTTP"...)
			i += 6
		} else {
			out = append(out, b[i])
		}
	}
	return string(out)
}

func main() {
	fmt.Println(MaskLink("Here's my spammy page: http://hehefouls.netHAHAHA see you."))
}
