package transport

import (
	"fmt"
	"os"
	"strings"
)

func (s *Commander) Executor(in string) {
	in = strings.TrimSpace(in)
	blocks := strings.Split(in, " ")

	switch blocks[0] {
	case registration:
		if len(blocks) != 3 {
			fmt.Println("bad registration format")
			return
		}
		login := blocks[1]
		password := blocks[2]

		s.Log.Debug().Msgf("command: %s, login: %s, password: %s", blocks[0], blocks[1], blocks[2])

		err := s.Sender.Registration(login, password)
		if err != nil {
			fmt.Printf("registration failed: %s\n", err.Error())
			return
		}
		fmt.Println("registration is successful")
	case authentication:
		fmt.Println("i input auth")
	case exit:
		fmt.Println("GophKeeper stop")
		os.Exit(0)
		return
	}
}
