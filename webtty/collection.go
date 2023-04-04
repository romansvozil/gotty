package webtty

import "sync"

type CollectionWebTTY struct {
	sync.Mutex
	webTTYs []*WebTTY
}

func (collection *CollectionWebTTY) Push(item *WebTTY) {
	collection.Lock()
	defer collection.Unlock()

	collection.webTTYs = append(collection.webTTYs, item)
}

/*
	Expects the item is in the list only once
 */
func (collection *CollectionWebTTY) Remove(item *WebTTY) {
	collection.Lock()
	defer collection.Unlock()

	for index, val := range collection.webTTYs {
		if val == item {
			collection.webTTYs = append(collection.webTTYs[:index], collection.webTTYs[index+1:]...)
			return
		}
	}

}

func (collection *CollectionWebTTY) ForEach(callback func(tty *WebTTY)) {
	collection.Lock()
	defer collection.Unlock()

	for _, wtty := range collection.webTTYs {
		callback(wtty)
	}
}
