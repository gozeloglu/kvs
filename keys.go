package kvs

// Keys list the stored keys in a slice. If db is empty, empty slice is returned.
func (k *Kvs) Keys() []string {
	k.mu.Lock()
	defer k.mu.Unlock()
	var keys []string

	for key := range k.kv {
		keys = append(keys, key)
	}

	return keys
}
