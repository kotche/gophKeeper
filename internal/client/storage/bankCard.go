package storage

import "github.com/kotche/gophKeeper/internal/client/domain"

func (c *Cache) AddBankCard(dt *domain.BankCard) error {
	c.Log.Debug().Msgf("cache add bank card data '%+v'", dt)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.bankData[dt.ID] = dt

	return nil
}

func (c *Cache) UpdateBankCard(dt *domain.BankCard) error {
	c.Log.Debug().Msgf("cache update bank card data '%+v'", dt)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.bankData[dt.ID] = dt

	return nil
}

func (c *Cache) DeleteBankCard(id int) error {
	c.Log.Debug().Msgf("cache delete bank card data '%d'", id)
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.bankData, id)

	return nil
}

func (c *Cache) ReadAllBankCard() ([]*domain.BankCard, error) {
	c.Log.Debug().Msgf("cache read all bank card data")
	c.mu.RLock()
	defer c.mu.RUnlock()

	data := make([]*domain.BankCard, 0, len(c.lpData))
	for _, v := range c.bankData {
		data = append(data, v)
	}

	return data, nil
}

func (c *Cache) UpdateAllBankCard(bankCardData []*domain.BankCard) error {
	l := len(bankCardData)
	//c.Log.Debug().Msgf("cache UpdateAllBankCard, len: %d", l)
	mTemp := make(map[int]*domain.BankCard, l)
	for _, data := range bankCardData {
		mTemp[data.ID] = data
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	c.bankData = mTemp

	return nil
}
