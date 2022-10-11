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
	c.mu.Lock()
	defer c.mu.Unlock()
	c.userID = userID
	c.token = token
	c.lpData = make(map[int]*domain.LoginPass)
	c.version.Swap(0)
	return nil
}

func (c *Cache) GetToken() (string, error) {
	return c.token, nil
}

func (c *Cache) GetCurrentUserID() (int, error) {
	return c.userID, nil
}

func (c *Cache) GetVersion() (int, error) {
	return int(c.version.Load()), nil
}

func (c *Cache) IncVersion() error {
	c.version.Add(1)
	return nil
}

func (c *Cache) SetVersion(version int) error {
	c.version.Swap(uint64(version))
	return nil
}
