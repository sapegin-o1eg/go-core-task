package utils

type StringIntMap struct {
	M map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		M: make(map[string]int),
	}
}

func (sm *StringIntMap) Add(key string, value int) {
	sm.M[key] = value
}

func (sm *StringIntMap) Remove(key string) {
	delete(sm.M, key)
}

func (sm *StringIntMap) Copy() map[string]int {
	copy := make(map[string]int, len(sm.M))
	for key, value := range sm.M {
		copy[key] = value
	}
	return copy
}

func (sm *StringIntMap) Get(key string) (value int, ok bool) {
	value, ok = sm.M[key]
	return
}

func (sm *StringIntMap) Exists(key string) bool {
	_, ok := sm.M[key]
	return ok
}
