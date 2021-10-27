package kvs

// Del deletes the key from the memory.
func (k *Kvs) Del(key string) {
	delete(k.kv, key)
}
