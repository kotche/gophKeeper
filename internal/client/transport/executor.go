package transport

import (
	"fmt"
	"os"
	"strings"
)

func (c *Commander) Executor(in string) {
	in = strings.TrimSpace(in)
	blocks := strings.Split(in, " ")

	switch blocks[0] {
	case registration:
		c.UserLogin(blocks)
	case authentication:
		c.UserAuthentication(blocks)
	case create:
		c.CreateDataType(in)
	case update:
		//update
	case delete:
		//delete
	case read:
		//read
	case exit:
		fmt.Println("GophKeeper stop")
		os.Exit(0)
		return
	default:
		fmt.Println(invalidFormat)
	}
}

func (c *Commander) UserLogin(blocks []string) {
	if len(blocks) != 3 {
		fmt.Println("bad registration format")
		return
	}
	username := blocks[1]
	password := blocks[2]

	c.Log.Debug().Msgf("user reg command: %s, username: %s, password: %s", blocks[0], blocks[1], blocks[2])

	err := c.Sender.Login(username, password)
	if err != nil {
		fmt.Printf("registration failed: %s\n", err.Error())
		return
	}
	fmt.Println("registration is successful")
}

func (c *Commander) UserAuthentication(blocks []string) {
	if len(blocks) != 3 {
		fmt.Println("bad authentication format")
		return
	}
	username := blocks[1]
	password := blocks[2]

	c.Log.Debug().Msgf("user auth command: %s, login: %s, password: %s", blocks[0], blocks[1], blocks[2])

	err := c.Sender.Authentication(username, password)
	if err != nil {
		fmt.Printf("authentication failed: %s\n", err.Error())
		return
	}
	fmt.Println("authentication is successful")
}

func (c *Commander) CreateDataType(in string) {
	blocks := strings.Split(in, " ")
	if len(blocks) < 3 {
		fmt.Println(invalidFormat)
		return
	}

	indMeta, meta := c.getMetaInfo(in, blocks)
	if indMeta > 0 {
		blocks = blocks[:indMeta]
	}

	switch blocks[1] {
	case loginPassDataType:
		login := blocks[2]
		password := blocks[3]

		c.Log.Debug().Msgf("create lp: login: %s, password: %s meta: %s", login, password, meta)

		err := c.Sender.CreateLoginPass(login, password, meta)
		if err != nil {
			fmt.Printf("create login password failed: %s\n", err.Error())
			return
		}

	default:
		fmt.Println(invalidFormat)
	}
}

func (c *Commander) getMetaInfo(in string, blocks []string) (int, string) {
	var indMeta int
	if strings.Contains(in, metaInfo) {
		for i, v := range blocks {
			if string(v) == metaInfo {
				indMeta = i
				break
			}
		}
	}
	meta := strings.Join(blocks[indMeta+1:], " ")
	return indMeta, meta
}
