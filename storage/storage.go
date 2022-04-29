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
	return s.Items[key].value
}