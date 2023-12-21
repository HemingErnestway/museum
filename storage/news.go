package storage

import (
	"museum/entity"
	"sync"
	"time"
)

type NewsMx struct {
	mtx     sync.RWMutex
	iter    uint32
	newsMap map[uint32]entity.News
}

var newsMx NewsMx

func init() {
	newsMx = NewsMx{
		newsMap: make(map[uint32]entity.News),
	}
}

func NewsCreate(news entity.News) *entity.News {
	newsMx.mtx.Lock()
	defer newsMx.mtx.Unlock()

	newsMx.iter++
	news.Uuid = newsMx.iter
	news.DateTime = time.Now().Format("2006-01-02 15:04")
	newsMx.newsMap[newsMx.iter] = news

	return &news
}

func NewsReadSingle(id uint32) *entity.News {
	newsMx.mtx.RLock()
	defer newsMx.mtx.RUnlock()

	if el, ok := newsMx.newsMap[id]; ok {
		return &el
	}

	return nil
}

func NewsRead() []entity.News {
	newsMx.mtx.RLock()
	defer newsMx.mtx.RUnlock()

	newsList := make([]entity.News, len(newsMx.newsMap))
	iter := 0
	for key := range newsMx.newsMap {
		newsList[iter] = newsMx.newsMap[key]
		iter++
	}

	return newsList
}

func NewsUpdate(new entity.News, id uint32) *entity.News {
	newsMx.mtx.Lock()
	defer newsMx.mtx.Unlock()

	current := newsMx.newsMap[id]

	if new.Header != "" {
		current.Header = new.Header
	} else if new.Content != "" {
		current.Content = new.Content
	}

	newsMx.newsMap[id] = current
	return &current
}

func NewsDelete(id uint32) string {
	newsMx.mtx.Lock()
	defer newsMx.mtx.Unlock()

	delete(newsMx.newsMap, id)

	return "successfully deleted"
}

func NewsSlice() []entity.News {
	newsMx.mtx.RLock()
	defer newsMx.mtx.RUnlock()

	v := make([]entity.News, len(newsMx.newsMap))
	for _, val := range newsMx.newsMap {
		v = append(v, val)
	}

	return v
}
