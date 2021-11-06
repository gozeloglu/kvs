package kvs

// Exists checks if the memory contains the key.
func (k *Kvs) Exists(key string) bool {
	k.mu.Lock()
	_, ok := k.kv[key]
	k.mu.Unlock()
	return ok
}
