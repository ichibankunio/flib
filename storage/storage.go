package storage

type Storage struct {
	Items map[string]*item
}

type item struct {
	key string
	value interface{}
}

func (s *Storage) Init() {
	s.Items = map[string]*item{}
}

func (s *Storage) SetItem(key string, value interface{}) {
	s.Items[key] = &item{
		key: key,
		value: value,
	}
}

func (s *Storage) GetItem(key string) interface{} {
	item, isThere := s.Items[key]
	if isThere {
		return item.value
	}

	return nil
}

func (s *Storage) IsThere(key string) bool {
	_, isThere := s.Items[key]
	return isThere
}