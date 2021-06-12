package util

type StringSet struct {
	m map[string]struct{}
}

func NewStringSet() *StringSet {
	return &StringSet{
		m: make(map[string]struct{}),
	}
}

func (s *StringSet) Add(elem string) {
	s.m[elem] = struct{}{}
}

func (s *StringSet) Delete(elem string) {
	delete(s.m, elem)
}

func (s *StringSet) Has(elem string) bool {
	_, has := s.m[elem]
	return has
}

func StringSetFromSlice(elems []string) *StringSet {
	s := NewStringSet()
	for _, elem := range elems {
		s.Add(elem)
	}
	return s
}
