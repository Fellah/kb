package cache

var cache map[string][]byte

func init() {
	cache = make(map[string][]byte)
}

func Set(k string, v []byte) {
	cache[k] = v
}

func Get(k string) []byte {
	if v, ok := cache[k]; ok {
		return v
	}

	return nil
}
