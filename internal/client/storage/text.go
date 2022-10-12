package storage

import "github.com/kotche/gophKeeper/internal/client/domain"

func (c *Cache) AddText(dt *domain.Text) error {
	c.Log.Debug().Msgf("cache add text data '%+v'", dt)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.textData[dt.ID] = dt

	return nil
}

func (c *Cache) UpdateText(dt *domain.Text) error {
	c.Log.Debug().Msgf("cache update text data '%+v'", dt)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.textData[dt.ID] = dt

	return nil
}

func (c *Cache) DeleteText(id int) error {
	c.Log.Debug().Msgf("cache delete text data '%d'", id)
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.textData, id)

	return nil
}

func (c *Cache) ReadAllText() ([]*domain.Text, error) {
	c.Log.Debug().Msgf("cache read all text data")
	c.mu.RLock()
	defer c.mu.RUnlock()

	data := make([]*domain.Text, 0, len(c.lpData))
	for _, v := range c.textData {
		data = append(data, v)
	}

	return data, nil
}

func (c *Cache) UpdateAllText(textData []*domain.Text) error {
	l := len(textData)
	//c.Log.Debug().Msgf("cache UpdateAllText, len: %d", l)
	mTemp := make(map[int]*domain.Text, l)
	for _, data := range textData {
		mTemp[data.ID] = data
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	c.textData = mTemp

	return nil
}
