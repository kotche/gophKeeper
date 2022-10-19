package storage

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"

	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/kotche/gophKeeper/internal/client/domain/dataType"
	"github.com/rs/zerolog"
)

// Cache client local repository
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

// SetUserParams records local repository data for an authorized user
func (c *Cache) SetUserParams(userID int, token string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.userID = userID
	c.token = token
	c.lpData = make(map[int]*domain.LoginPass)
	c.version.Swap(0)
}

// GetToken gets a token for authorization
func (c *Cache) GetToken() string {
	return c.token
}

// GetCurrentUserID gets the current user id
func (c *Cache) GetCurrentUserID() int {
	return c.userID
}

// GetVersion gets the current version data
func (c *Cache) GetVersion() int {
	return int(c.version.Load())
}

// IncVersion writes a new version of the data
func (c *Cache) IncVersion() {
	c.version.Add(1)
}

// SetVersion records the version of the server database data
func (c *Cache) SetVersion(version int) {
	c.version.Swap(uint64(version))
}

// Save writes data
func (c *Cache) Save(data any) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	switch d := data.(type) {
	case *domain.LoginPass:
		c.lpData[d.ID] = d
	case *domain.Text:
		c.textData[d.ID] = d
	case *domain.Binary:
		c.binaryData[d.ID] = d
	case *domain.BankCard:
		c.bankData[d.ID] = d
	default:
		err := fmt.Errorf("unsupported type '%v'", reflect.TypeOf(data))
		c.Log.Err(err).Msg("cache save error")
		return err
	}

	return nil
}

// Update updates data
func (c *Cache) Update(data any) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	switch d := data.(type) {
	case *domain.LoginPass:
		c.lpData[d.ID] = d
	case *domain.Text:
		c.textData[d.ID] = d
	case *domain.Binary:
		c.binaryData[d.ID] = d
	case *domain.BankCard:
		c.bankData[d.ID] = d
	default:
		err := fmt.Errorf("unsupported type '%v'", reflect.TypeOf(data))
		c.Log.Err(err).Msg("cache update error")
		return err
	}
	return nil
}

// Delete deletes data
func (c *Cache) Delete(data any) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	switch d := data.(type) {
	case *domain.LoginPass:
		delete(c.lpData, d.ID)
	case *domain.Text:
		delete(c.textData, d.ID)
	case *domain.Binary:
		delete(c.binaryData, d.ID)
	case *domain.BankCard:
		delete(c.bankData, d.ID)
	default:
		err := fmt.Errorf("unsupported type '%v'", reflect.TypeOf(data))
		c.Log.Err(err).Msg("cache delete error")
		return err
	}
	return nil
}

// GetAll gets data by data type
func (c *Cache) GetAll(dt dataType.DataType) (any, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	switch dt {
	case dataType.LP:
		data := make([]*domain.LoginPass, 0, len(c.lpData))
		for _, v := range c.lpData {
			data = append(data, v)
		}
		return data, nil
	case dataType.TEXT:
		data := make([]*domain.Text, 0, len(c.textData))
		for _, v := range c.textData {
			data = append(data, v)
		}
		return data, nil
	case dataType.BINARY:
		data := make([]*domain.Binary, 0, len(c.binaryData))
		for _, v := range c.binaryData {
			data = append(data, v)
		}
		return data, nil
	case dataType.BANKCARD:
		data := make([]*domain.BankCard, 0, len(c.bankData))
		for _, v := range c.bankData {
			data = append(data, v)
		}
		return data, nil
	default:
		err := fmt.Errorf("unsupported type '%v'", reflect.TypeOf(dt))
		c.Log.Err(err).Msg("cache getAll error")
		return nil, err
	}

}

// UpdateAll updates data
func (c *Cache) UpdateAll(data any) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	switch d := data.(type) {
	case []*domain.LoginPass:
		mTemp := make(map[int]*domain.LoginPass, len(d))
		for _, v := range d {
			mTemp[v.ID] = v
		}
		c.lpData = mTemp
	case []*domain.Text:
		mTemp := make(map[int]*domain.Text, len(d))
		for _, v := range d {
			mTemp[v.ID] = v
		}
		c.textData = mTemp
	case []*domain.Binary:
		mTemp := make(map[int]*domain.Binary, len(d))
		for _, v := range d {
			mTemp[v.ID] = v
		}
		c.binaryData = mTemp
	case []*domain.BankCard:
		mTemp := make(map[int]*domain.BankCard, len(d))
		for _, v := range d {
			mTemp[v.ID] = v
		}
		c.bankData = mTemp
	default:
		err := fmt.Errorf("unsupported type '%v'", reflect.TypeOf(data))
		c.Log.Err(err).Msg("cache updateAll error")
		return err
	}

	return nil
}
