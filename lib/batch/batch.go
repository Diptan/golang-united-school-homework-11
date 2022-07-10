package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var ops int64
	var mu sync.Mutex
	var wg sync.WaitGroup
	for int64(len(res)) < n {
		for i := 0; int64(i) < pool; i++ {
			wg.Add(1)
			go func() {
				ops++
				u := getOne(ops)
				mu.Lock()
				res = append(res, u)
				mu.Unlock()
				wg.Done()
			}()
		}
		wg.Wait()
	}
	return
}
