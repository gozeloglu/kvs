package kvs

// Get returns the value of the given key.
func (k *Kvs) Get(key string) string {
	return k.kv[key]
}
