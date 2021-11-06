package kvs

// Set assigns the value to key in map.
func (k *Kvs) Set(key, value string) {
	k.mu.Lock()
	k.kv[key] = value
	k.mu.Unlock()
}
