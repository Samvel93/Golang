package main

import "fmt"

func MaskLink(b string) string {
	src := []byte(b)
	out := make([]byte, 0, len(src))
	http := []byte("http://")
	for i := 0; i < len(src); i++ {
		if i+len(http) <= len(src) {
			match := true
			for k := 0; k < len(http); k++ {
				if src[i+k] != http[k] {
					match = false
					break
				}
			}
			if match {
				out = append(out, []byte("HTTP")...)
				i += len(http) - 1
				continue
			}
		}
		out = append(out, src[i])
	}
	return string(out)
}

func main() {
	fmt.Println(MaskLink("Here's my spammy page: http://hehefouls.netHAHAHA see you."))
}
