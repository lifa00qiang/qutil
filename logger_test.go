package QUtil

import (
	"golang.org/x/sync/errgroup"
	"testing"
)

func TestGetLogger(t *testing.T) {

	logger := GetLogger()
	var eg errgroup.Group

	eg.Go(func() error {

		for i := 0; i < 10000; i++ {
			logger.Info(i)
		}

		return nil
	})

	eg.Go(func() error {

		for i := 0; i < 10000; i++ {
			logger.Error(i)
		}

		return nil
	})

	eg.Go(func() error {
		for i := 0; i < 10000; i++ {
			logger.Debug(i)
		}

		return nil
	})

	eg.Wait()

}
