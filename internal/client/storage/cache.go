package storage

import (
	"sync"
	"sync/atomic"

	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/rs/zerolog"
)

type Cache struct {
	userID     int
	token      string
	version    atomic.Uint64
	lpData     map[int]*domain.LoginPass
	textData   map[int]*domain.Text
	binaryData map[int]*domain.Binary
	bankData   map[int]*domain.BankCard

	mu  sync.RWMutex
	Log *zerolog.Logger
}

func NewCache(log *zerolog.Logger) *Cache {
	return &Cache{
		userID:     -1,
		Log:        log,
		lpData:     make(map[int]*domain.LoginPass),
		textData:   make(map[int]*domain.Text),
		binaryData: make(map[int]*domain.Binary),
		bankData:   make(map[int]*domain.BankCard),
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

func (c *Cache) IncVersion() error {
	c.version.Add(1)
	return nil
}

func (c *Cache) AddLoginPassword(dt *domain.LoginPass) error {
	c.Log.Debug().Msgf("add lp '%+v' from cache", dt)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.lpData[dt.ID] = dt

	return nil
}

func (c *Cache) ReadAllLoginPassword() ([]*domain.LoginPass, error) {
	c.Log.Debug().Msgf("read all lp from cache")
	c.mu.RLock()
	defer c.mu.RUnlock()

	lpPairs := make([]*domain.LoginPass, 0, len(c.lpData))
	for _, v := range c.lpData {
		lpPairs = append(lpPairs, v)
	}

	return lpPairs, nil
}
