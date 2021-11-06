package kvs

import "strconv"

// Decr decreases value of the given key by one. ıt is only applicable for
// integer convertable values. If the value is not able to convert integer, it
// error with empty string.
func (k *Kvs) Decr(key string) (string, error) {
	k.mu.Lock()
	defer k.mu.Unlock()

	val := k.kv[key]
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return "", err
	}
	k.kv[key] = strconv.Itoa(valInt - 1)
	return k.kv[key], nil
}

// DecrBy decreases the value by given value. If non-integer value is being
// tried to decrement, it returns error.
func (k *Kvs) DecrBy(key string, val int) (string, error) {
	k.mu.Lock()
	defer k.mu.Unlock()

	v := k.kv[key]
	valInt, err := strconv.Atoi(v)
	if err != nil {
		return "", err
	}
	k.kv[key] = strconv.Itoa(valInt - val)
	return k.kv[key], nil
}
