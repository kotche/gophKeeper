package transport

import (
	"fmt"
	"strconv"
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
		c.CreateData(in)
	case update:
		c.UpdateData(in)
	case delete:
		c.DeleteData(in)
	case read:
		c.ReadData(in)
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

func (c *Commander) CreateData(in string) {
	blocks := strings.Split(in, " ")
	if len(blocks) < 4 {
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
		fmt.Println("create login password successful")
	default:
		fmt.Println(invalidFormat)
	}
}

func (c *Commander) UpdateData(in string) {
	blocks := strings.Split(in, " ")
	if len(blocks) < 5 {
		fmt.Println(invalidFormat)
		return
	}

	indMeta, meta := c.getMetaInfo(in, blocks)
	if indMeta > 0 {
		blocks = blocks[:indMeta]
	}

	switch blocks[1] {
	case loginPassDataType:
		idStr := blocks[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.Log.Err(err).Msgf("commander updateData convert id '%d' to int error", idStr)
			fmt.Printf("id %s is not a number", idStr)
			return
		}
		login := blocks[3]
		password := blocks[4]

		c.Log.Debug().Msgf("update lp: id %d login: %s, password: %s meta: %s", id, login, password, meta)

		err = c.Sender.UpdateLoginPass(id, login, password, meta)
		if err != nil {
			fmt.Printf("update login password failed: %s\n", err.Error())
			return
		}
		fmt.Println("update login password successful")
	default:
		fmt.Println(invalidFormat)
	}
}

func (c *Commander) DeleteData(in string) {
	blocks := strings.Split(in, " ")
	if len(blocks) < 3 {
		fmt.Println(invalidFormat)
		return
	}

	switch blocks[1] {
	case loginPassDataType:
		idStr := blocks[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.Log.Err(err).Msgf("commander deleteData convert id '%d' to int error", idStr)
			fmt.Printf("id %s is not a number", idStr)
			return
		}

		c.Log.Debug().Msgf("delete lp: id %d", id)

		err = c.Sender.DeleteLoginPass(id)
		if err != nil {
			fmt.Printf("delete login password failed: %s\n", err.Error())
			return
		}
		fmt.Println("delete login password successful")
	default:
		fmt.Println(invalidFormat)
	}
}

func (c *Commander) ReadData(in string) {
	blocks := strings.Split(in, " ")
	if len(blocks) != 2 {
		fmt.Println(invalidFormat)
		return
	}

	switch blocks[1] {
	case loginPassDataType:
		c.Log.Debug().Msg("read lp")

		lpPairs, err := c.Sender.ReadLoginPassCache()
		if err != nil {
			fmt.Printf("failed read data login password : %s\n", err.Error())
			return
		}
		if len(lpPairs) == 0 {
			fmt.Println("no data login password")
		}
		for _, v := range lpPairs {
			fmt.Printf("id: %d, login: %s, password: %s, info: %s\n", v.ID, v.Login, v.Password, v.MetaInfo)
		}
	default:
		fmt.Println(invalidFormat)
	}
}
func (c *Commander) getMetaInfo(in string, blocks []string) (int, string) {
	var indMeta int
	if !strings.Contains(in, metaInfo) {
		return -1, ""
	}
	for i, v := range blocks {
		if string(v) == metaInfo {
			indMeta = i
			break
		}
	}

	meta := strings.Join(blocks[indMeta+1:], " ")
	return indMeta, meta
}
