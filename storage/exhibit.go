package storage

import (
	"museum/entity"
	"sync"
)

type ExhibitMx struct {
	mtx      sync.RWMutex
	iter     uint32
	exhibits map[uint32]entity.Exhibit
}

var exhibitMx ExhibitMx

func init() {
	exhibitMx = ExhibitMx{
		exhibits: make(map[uint32]entity.Exhibit),
	}
}

func ExhibitCreate(exhibit entity.Exhibit) *entity.Exhibit {
	exhibitMx.mtx.Lock()
	defer exhibitMx.mtx.Unlock()

	exhibitMx.iter++
	exhibit.Uuid = exhibitMx.iter
	exhibitMx.exhibits[exhibitMx.iter] = exhibit

	return &exhibit
}

func ExhibitRead(id uint32) *entity.Exhibit {
	exhibitMx.mtx.RLock()
	defer exhibitMx.mtx.RUnlock()

	if el, ok := exhibitMx.exhibits[id]; ok {
		return &el
	}

	return nil
}

func ExhibitsRead() []entity.Exhibit {
	exhibitMx.mtx.RLock()
	defer exhibitMx.mtx.RUnlock()

	exhibitList := make([]entity.Exhibit, len(exhibitMx.exhibits))
	iter := 0
	for key := range exhibitMx.exhibits {
		exhibitList[iter] = exhibitMx.exhibits[key]
		iter++
	}

	return exhibitList
}

func ExhibitUpdate(new entity.Exhibit, id uint32) *entity.Exhibit {
	exhibitMx.mtx.Lock()
	defer exhibitMx.mtx.Unlock()

	current := exhibitMx.exhibits[id]

	if new.Name != "" {
		current.Name = new.Name
	} else if new.Desc != "" {
		current.Desc = new.Desc
	}

	exhibitMx.exhibits[id] = current
	return &current
}

func ExhibitDelete(id uint32) string {
	exhibitMx.mtx.Lock()
	defer exhibitMx.mtx.Unlock()

	delete(exhibitMx.exhibits, id)

	return "successfully deleted"
}
