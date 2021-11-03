package kvs

// Exists checks if the memory contains the key.
func (k *Kvs) Exists(key string) bool {
	_, ok := k.kv[key]
	return ok
}
