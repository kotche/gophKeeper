package storage

import "github.com/kotche/gophKeeper/internal/client/domain"

func (c *Cache) AddBinary(dt *domain.Binary) error {
	c.Log.Debug().Msgf("cache add binary data '%+v'", dt)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.binaryData[dt.ID] = dt

	return nil
}

func (c *Cache) UpdateBinary(dt *domain.Binary) error {
	c.Log.Debug().Msgf("cache update binary data '%+v'", dt)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.binaryData[dt.ID] = dt

	return nil
}

func (c *Cache) DeleteBinary(id int) error {
	c.Log.Debug().Msgf("cache delete binary data '%d'", id)
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.binaryData, id)

	return nil
}

func (c *Cache) ReadAllBinary() ([]*domain.Binary, error) {
	c.Log.Debug().Msgf("cache read all binary data")
	c.mu.RLock()
	defer c.mu.RUnlock()

	data := make([]*domain.Binary, 0, len(c.lpData))
	for _, v := range c.binaryData {
		data = append(data, v)
	}

	return data, nil
}

func (c *Cache) UpdateAllBinary(binaryData []*domain.Binary) error {
	l := len(binaryData)
	//c.Log.Debug().Msgf("cache UpdateAllBinary, len: %d", l)
	mTemp := make(map[int]*domain.Binary, l)
	for _, data := range binaryData {
		mTemp[data.ID] = data
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	c.binaryData = mTemp

	return nil
}
