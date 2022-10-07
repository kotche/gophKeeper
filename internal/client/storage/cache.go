package storage

import "sync"

type Cache struct {
	userID int
	token  string
	data   map[string]string
	mu     sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		userID: -1,
		data:   make(map[string]string),
	}
}

func (c *Cache) SetUserParams(userID int, token string) error {
	c.userID = userID
	c.token = token
	return nil
}

func (c *Cache) GetToken() (string, error) {
	return c.token, nil
}

func (c *Cache) GetCurrentUserID() (int, error) {
	return c.userID, nil
}
