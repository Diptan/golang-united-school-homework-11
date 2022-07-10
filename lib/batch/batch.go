package batch

import (
	"context"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var mu sync.Mutex
	errG, _ := errgroup.WithContext(context.Background())
	errG.SetLimit(int(pool))
	for i := int64(0); i < n; i++ {
		func(i int64) {
			errG.Go(func() error {
				u := getOne(i)
				mu.Lock()
				res = append(res, u)
				mu.Unlock()
				return nil
			})
		}(i)
	}
	errG.Wait()
	return
}
