package main

import (
	"abcdefg/server"
)

func main() {
	s, err := server.NewServer(".ssh/term_info_ed25519", "localhost", 2324)
	if err != nil {
		return
	}
	s.Start()
}

