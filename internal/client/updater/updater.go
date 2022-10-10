package updater

import (
	"context"
	"time"

	"github.com/kotche/gophKeeper/config/client"
	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

type ISender interface {
	GetVersionServer(ctx context.Context) (int, error)
	GetAllLoginPass(ctx context.Context) ([]*domain.LoginPass, error)
}

type IService interface {
	GetVersionCache() (int, error)
	SetVersionCache(version int) error
	UpdateAllLoginPassCache(lpPairs []*domain.LoginPass) error
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
	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		err := w.updateLoginPassCache(ctx)
		if err != nil {
			return err
		}
		return nil
	})

	if err := grp.Wait(); err != nil {
		return err
	}

	return nil
}

func (w *Updater) updateLoginPassCache(ctx context.Context) error {
	lpPairs, err := w.Sender.GetAllLoginPass(ctx)
	if err != nil {
		w.Log.Err(err).Msg("worker getAllLoginPass error")
		return err
	}
	err = w.Service.UpdateAllLoginPassCache(lpPairs)
	if err != nil {
		w.Log.Err(err).Msg("worker updateAllLoginPass error")
		return err
	}
	return nil
}
