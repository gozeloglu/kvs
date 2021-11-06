package kvs

// Get returns the value of the given key.
func (k *Kvs) Get(key string) string {
	k.mu.Lock()
	val := k.kv[key]
	k.mu.Unlock()
	return val
}
