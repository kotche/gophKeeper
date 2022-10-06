package storage

import "sync"

type Cache struct {
	UserID int
	Token  string
	Data   map[string]string
	mu     sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		UserID: -1,
		Data:   make(map[string]string),
	}
}

func (c *Cache) SetUserParams(userID int, token string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.UserID = userID
	c.Token = token
	return nil
}

func (c *Cache) GetCurrentUserID() (int, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.UserID, nil
}
