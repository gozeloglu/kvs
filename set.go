package kvs

// Set assigns the value to key in map.
func (k *Kvs) Set(key, value string) {
	k.kv[key] = value
}
