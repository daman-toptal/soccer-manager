package service

import "sync"

type asyncWaitGroup struct {
	wg *sync.WaitGroup
}

type AsyncWaitGroup interface {
	Wait()
	Add(int)
	Done()
}

func NewAsyncWaitGroupService() AsyncWaitGroup {
	return &asyncWaitGroup{
		wg: &sync.WaitGroup{},
	}
}

func (a asyncWaitGroup) Wait() {
	a.wg.Wait()
}

func (a asyncWaitGroup) Done() {
	a.wg.Done()
}

func (a asyncWaitGroup) Add(i int) {
	a.wg.Add(i)
}
