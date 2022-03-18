package syncutil

import "sync"

// 構造体の型をエクスポートして、フィールドをエクスポートしない用途の一例
// カウンターは、不用意なカウント数の変更を防ぎ、排他的にカウントアップが行う必要がある
type Counter struct {
	Name string

	m     sync.RWMutex
	count int
}

func (c *Counter) Increment() int {
	c.m.Lock()
	defer c.m.Unlock()
	c.count++
	return c.count
}

func (c *Counter) View() int {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.count
}
