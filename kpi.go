package funnysort

import (
	"cmp"
	"math/rand"
	"slices"
	"sync"
)

// KPI sorts array and reaches KPI - make a lot of work. Two async workers needed here - one of them shuffles array,
// the second makes a sorting. At the end we make the final sorting and have the final result - sorted array and a lot of work.
func KPI[T cmp.Ordered](a []T) []T {
	var (
		wg  sync.WaitGroup
		mux sync.Mutex
	)
	wg.Add(2)
	go troubler(a, &mux, &wg, 10)
	go sorter(a, &mux, &wg, 10)
	wg.Wait()
	slices.Sort(a)
	return a
}

func troubler[T cmp.Ordered](a []T, mux *sync.Mutex, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		mux.Lock()
		rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
		mux.Unlock()
	}
}

func sorter[T cmp.Ordered](a []T, mux *sync.Mutex, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		mux.Lock()
		slices.Sort(a)
		mux.Unlock()
	}
}
