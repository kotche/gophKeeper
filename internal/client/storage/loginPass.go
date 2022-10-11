package storage

import "github.com/kotche/gophKeeper/internal/client/domain"

func (c *Cache) AddLoginPassword(dt *domain.LoginPass) error {
	c.Log.Debug().Msgf("cache add lp '%+v'", dt)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.lpData[dt.ID] = dt

	return nil
}

func (c *Cache) UpdateLoginPassword(dt *domain.LoginPass) error {
	c.Log.Debug().Msgf("cache update lp '%+v'", dt)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.lpData[dt.ID] = dt

	return nil
}

func (c *Cache) DeleteLoginPassword(id int) error {
	c.Log.Debug().Msgf("cache delete lp '%d'", id)
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.lpData, id)

	return nil
}

func (c *Cache) ReadAllLoginPassword() ([]*domain.LoginPass, error) {
	c.Log.Debug().Msgf("cache read all lp")
	c.mu.RLock()
	defer c.mu.RUnlock()

	lpPairs := make([]*domain.LoginPass, 0, len(c.lpData))
	for _, v := range c.lpData {
		lpPairs = append(lpPairs, v)
	}

	return lpPairs, nil
}

func (c *Cache) UpdateAllLoginPass(lpPairs []*domain.LoginPass) error {
	l := len(lpPairs)
	//c.Log.Debug().Msgf("cache updateAllLoginPass, len: %d", l)
	mTemp := make(map[int]*domain.LoginPass, l)
	for _, lp := range lpPairs {
		mTemp[lp.ID] = lp
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	c.lpData = mTemp

	return nil
}
