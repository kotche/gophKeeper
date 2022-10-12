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
	if len(blocks) < 3 {
		fmt.Println(invalidFormat)
		return
	}

	var indEnd int
	indMeta, meta := c.getMetaInfo(in, blocks)
	if indMeta > 0 {
		blocks = blocks[:indMeta]
		indEnd = indMeta
	} else {
		indEnd = len(blocks)
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
	case textDataType:
		text := strings.Join(blocks[2:indEnd], " ")

		c.Log.Debug().Msgf("create text: text: %s, meta: %s", text, meta)

		err := c.Sender.CreateText(text, meta)
		if err != nil {
			fmt.Printf("create text data failed: %s\n", err.Error())
			return
		}
		fmt.Println("create text data successful")
	case binaryDataType:
		binary := strings.Join(blocks[2:indEnd], " ")

		c.Log.Debug().Msgf("create binary: binary: %s, meta: %s", binary, meta)

		err := c.Sender.CreateBinary(binary, meta)
		if err != nil {
			fmt.Printf("create binary data failed: %s\n", err.Error())
			return
		}
		fmt.Println("create binary data successful")
	case bankCardDataType:
		number := blocks[2]

		c.Log.Debug().Msgf("create bank card: number: %s, meta: %s", number, meta)

		err := c.Sender.CreateBankCard(number, meta)
		if err != nil {
			fmt.Printf("create bank card failed: %s\n", err.Error())
			return
		}
		fmt.Println("create bank card successful")
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

	var indEnd int
	indMeta, meta := c.getMetaInfo(in, blocks)
	if indMeta > 0 {
		blocks = blocks[:indMeta]
		indEnd = indMeta
	} else {
		indEnd = len(blocks)
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

		c.Log.Debug().Msgf("update lp: id '%d' login: '%s' password: '%s' meta: '%s'", id, login, password, meta)

		err = c.Sender.UpdateLoginPass(id, login, password, meta)
		if err != nil {
			fmt.Printf("update login password failed: %s\n", err.Error())
			return
		}
		fmt.Println("update login password successful")
	case textDataType:
		idStr := blocks[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.Log.Err(err).Msgf("commander updateData convert id '%d' to int error", idStr)
			fmt.Printf("id %s is not a number", idStr)
			return
		}
		text := strings.Join(blocks[2:indEnd], " ")

		c.Log.Debug().Msgf("update text: id %d text: %s meta: %s", id, text, meta)

		err = c.Sender.UpdateText(id, text, meta)
		if err != nil {
			fmt.Printf("update text data failed: %s\n", err.Error())
			return
		}
		fmt.Println("update text data successful")
	case binaryDataType:
		idStr := blocks[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.Log.Err(err).Msgf("commander updateData convert id '%d' to int error", idStr)
			fmt.Printf("id %s is not a number", idStr)
			return
		}
		binary := strings.Join(blocks[2:indEnd], " ")

		c.Log.Debug().Msgf("update binary: id %d text: %s meta: %s", id, binary, meta)

		err = c.Sender.UpdateBinary(id, binary, meta)
		if err != nil {
			fmt.Printf("update binary data failed: %s\n", err.Error())
			return
		}
		fmt.Println("update binary data successful")
	case bankCardDataType:
		idStr := blocks[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.Log.Err(err).Msgf("commander updateData convert id '%d' to int error", idStr)
			fmt.Printf("id %s is not a number", idStr)
			return
		}
		number := blocks[3]

		c.Log.Debug().Msgf("update bank card: id %d text: %s meta: %s", id, number, meta)

		err = c.Sender.UpdateBankCard(id, number, meta)
		if err != nil {
			fmt.Printf("update bank card data failed: %s\n", err.Error())
			return
		}
		fmt.Println("update bank card data successful")
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
	case textDataType:
		idStr := blocks[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.Log.Err(err).Msgf("commander deleteData convert id '%d' to int error", idStr)
			fmt.Printf("id %s is not a number", idStr)
			return
		}

		c.Log.Debug().Msgf("delete text data: id %d", id)

		err = c.Sender.DeleteText(id)
		if err != nil {
			fmt.Printf("delete text data failed: %s\n", err.Error())
			return
		}
		fmt.Println("delete text data successful")
	case binaryDataType:
		idStr := blocks[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.Log.Err(err).Msgf("commander deleteData convert id '%d' to int error", idStr)
			fmt.Printf("id %s is not a number", idStr)
			return
		}

		c.Log.Debug().Msgf("delete binary data: id %d", id)

		err = c.Sender.DeleteBinary(id)
		if err != nil {
			fmt.Printf("delete binary data failed: %s\n", err.Error())
			return
		}
		fmt.Println("delete binary data successful")
	case bankCardDataType:
		idStr := blocks[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.Log.Err(err).Msgf("commander deleteData convert id '%d' to int error", idStr)
			fmt.Printf("id %s is not a number", idStr)
			return
		}

		c.Log.Debug().Msgf("delete bank card : id %d", id)

		err = c.Sender.DeleteBankCard(id)
		if err != nil {
			fmt.Printf("delete bank card failed: %s\n", err.Error())
			return
		}
		fmt.Println("delete bank card successful")
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

		data, err := c.Sender.ReadLoginPassCache()
		if err != nil {
			fmt.Printf("failed read data login password : %s\n", err.Error())
			return
		}
		if len(data) == 0 {
			fmt.Println("no data login password")
		}
		for _, v := range data {
			fmt.Printf("id: %d, login: %s, password: %s, info: %s\n", v.ID, v.Login, v.Password, v.MetaInfo)
		}
	case textDataType:
		c.Log.Debug().Msg("read text")

		data, err := c.Sender.ReadTextCache()
		if err != nil {
			fmt.Printf("failed read text data : %s\n", err.Error())
			return
		}
		if len(data) == 0 {
			fmt.Println("no data text")
		}
		for _, v := range data {
			fmt.Printf("id: %d, text: %s, info: %s\n", v.ID, v.Text, v.MetaInfo)
		}
	case binaryDataType:
		c.Log.Debug().Msg("read binary")

		data, err := c.Sender.ReadBinaryCache()
		if err != nil {
			fmt.Printf("failed read binary data : %s\n", err.Error())
			return
		}
		if len(data) == 0 {
			fmt.Println("no data binary")
		}
		for _, v := range data {
			fmt.Printf("id: %d, binary: %s, info: %s\n", v.ID, v.Binary, v.MetaInfo)
		}
	case bankCardDataType:
		c.Log.Debug().Msg("read bank card")

		data, err := c.Sender.ReadBankCardCache()
		if err != nil {
			fmt.Printf("failed read bank card data : %s\n", err.Error())
			return
		}
		if len(data) == 0 {
			fmt.Println("no data bank card")
		}
		for _, v := range data {
			fmt.Printf("id: %d, number: %s, info: %s\n", v.ID, v.Number, v.MetaInfo)
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
