package storage

import (
	"sync"
	"sync/atomic"

	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/kotche/gophKeeper/internal/client/domain/dataType"
	"github.com/rs/zerolog"
)

type Cache struct {
	userID  int
	token   string
	version atomic.Uint64
	data    map[dataType.DataType]interface{}

	//TODO SLICE temporarily
	lpData     []*domain.LoginPass
	textData   []*domain.Text
	binaryData []*domain.Binary
	bankData   []*domain.BankCard

	mu  sync.RWMutex
	Log *zerolog.Logger
}

func NewCache(log *zerolog.Logger) *Cache {
	data := make(map[dataType.DataType]interface{})
	data[dataType.LP] = make([]*domain.LoginPass, 0)
	data[dataType.TEXT] = make([]*domain.Text, 0)
	data[dataType.BINARY] = make([]*domain.Binary, 0)
	data[dataType.BANKCARD] = make([]*domain.BankCard, 0)

	return &Cache{
		userID: -1,
		data:   data,
		Log:    log,
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

//TODO MAP
func (c *Cache) ReadData(dt dataType.DataType) (interface{}, error) {

	c.Log.Debug().Msgf("read data '%s' from cache", dt)
	c.mu.RLock()
	defer c.mu.RUnlock()
	data := c.data[dt]

	c.Log.Debug().Msgf("len cache map lp example: %d\n", len(data.([]*domain.LoginPass)))

	return data, nil
}

func (c *Cache) AddLoginPassword(dt *domain.LoginPass) error {
	c.Log.Debug().Msgf("add data '%+v' from cache", dt)
	c.mu.Lock()
	defer c.mu.Unlock()
	lpPairs := c.data[dataType.LP].([]*domain.LoginPass)
	lpPairs = append(lpPairs, dt)

	c.Log.Debug().Msgf("len cache lp: %d", len(lpPairs))

	//TODO SLICE temporarily
	c.lpData = append(c.lpData, dt)

	return nil
}

//TODO SLICE
func (c *Cache) ReadLoginPassword() ([]*domain.LoginPass, error) {
	c.Log.Debug().Msgf("read data lp from cache")
	c.mu.RLock()
	defer c.mu.RUnlock()

	data := make([]*domain.LoginPass, len(c.lpData))
	copy(data, c.lpData)

	return data, nil
}
