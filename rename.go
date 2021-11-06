package kvs

import "errors"

// Rename changes the name of the key. It deletes the old key-value pair. If the
// new key already exists, renaming cannot be done.
func (k *Kvs) Rename(key string, newKey string) (bool, error) {
	if k.Exists(newKey) {
		return false, errors.New("new key is already exists")
	}

	k.mu.Lock()
	k.kv[newKey] = k.kv[key]
	k.mu.Unlock()
	k.Del(key)
	return true, nil
}
