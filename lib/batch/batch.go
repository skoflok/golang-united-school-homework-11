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
	var mu sync.Mutex
	var wg sync.WaitGroup
	var i int64
	sem := make(chan struct{}, pool)

	work := func(u int64) {
		sem <- struct{}{}
		defer wg.Done()
		user := getOne(u)
		mu.Lock()
		res = append(res, user)
		mu.Unlock()
		<-sem
	}

	for i = 0; i < n; i++ {
		ii := i
		wg.Add(1)
		go work(ii)
	}
	wg.Wait()
	return res
}
