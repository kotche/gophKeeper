package updater

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/kotche/gophKeeper/config/client"
	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/kotche/gophKeeper/internal/client/domain/dataType"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

type ISender interface {
	GetVersionServer(ctx context.Context) (int, error)
	GetAllLoginPass(ctx context.Context) ([]*domain.LoginPass, error)
	GetAllText(ctx context.Context) ([]*domain.Text, error)
	GetAllBinary(ctx context.Context) ([]*domain.Binary, error)
	GetAllBankCard(ctx context.Context) ([]*domain.BankCard, error)
}

type IService interface {
	GetVersionCache() (int, error)
	SetVersionCache(version int) error
	UpdateAll(data any) error
}

type Updater struct {
	Sender  ISender
	Service IService
	Conf    *client.Config
	Log     *zerolog.Logger
}

func NewUpdater(sender ISender, service IService, conf *client.Config, log *zerolog.Logger) *Updater {
	return &Updater{
		Sender:  sender,
		Service: service,
		Conf:    conf,
		Log:     log,
	}
}

func (w *Updater) Run(ctx context.Context) {
	ticker := time.NewTicker(w.Conf.Updater.Timeout)

	for {
		select {
		case <-ticker.C:
			verCache, err := w.Service.GetVersionCache()
			if err != nil {
				w.Log.Err(err).Msg("worker getVersionCache error")
				continue
			}
			verServer, err := w.Sender.GetVersionServer(ctx)
			if err != nil {
				w.Log.Err(err).Msg("worker getVersionServer error")
				continue
			}
			if verCache >= verServer {
				continue
			}
			w.Log.Debug().Msgf("worker run, ver cache: %d, ver server: %d", verCache, verServer)
			err = w.updateData(ctx)
			if err != nil {
				w.Log.Err(err).Msg("worker updateData error")
				continue
			}
			w.Service.SetVersionCache(verServer)
		case <-ctx.Done():
			return
		}
	}
}

func (w *Updater) updateData(ctx context.Context) error {
	dataTypes := []dataType.DataType{
		dataType.LP,
		dataType.TEXT,
		dataType.BINARY,
		dataType.BANKCARD,
	}

	grp, ctx := errgroup.WithContext(ctx)
	for _, dt := range dataTypes {
		dt := dt
		grp.Go(func() error {
			err := w.update(ctx, dt)
			if err != nil {
				return err
			}
			return nil
		})
	}

	if err := grp.Wait(); err != nil {
		return err
	}

	return nil
}

func (w *Updater) update(ctx context.Context, dt dataType.DataType) error {
	var (
		data any
		err  error
	)

	switch dt {
	case dataType.LP:
		data, err = w.Sender.GetAllLoginPass(ctx)
		if err != nil {
			w.Log.Err(err).Msg("worker GetAllLoginPass error")
			return err
		}
	case dataType.TEXT:
		data, err = w.Sender.GetAllText(ctx)
		if err != nil {
			w.Log.Err(err).Msg("worker GetAllText error")
			return err
		}
	case dataType.BINARY:
		data, err = w.Sender.GetAllBinary(ctx)
		if err != nil {
			w.Log.Err(err).Msg("worker GetAllBinary error")
			return err
		}
	case dataType.BANKCARD:
		data, err = w.Sender.GetAllBankCard(ctx)
		if err != nil {
			w.Log.Err(err).Msg("worker GetAllBankCard error")
			return err
		}
	default:
		err = fmt.Errorf("unsupported type '%v'", reflect.TypeOf(dt))
		w.Log.Err(err).Msg("updater update error")
		return err
	}

	err = w.Service.UpdateAll(data)
	if err != nil {
		w.Log.Err(err).Msg("updater update error")
		return err
	}
	return nil
}
