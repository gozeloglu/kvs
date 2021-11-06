package kvs

// Del deletes the key from the memory.
func (k *Kvs) Del(key string) {
	k.mu.Lock()
	delete(k.kv, key)
	k.mu.Unlock()
}
