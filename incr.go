package kvs

import "strconv"

// Incr increments value of the key by 1. It is only applicable for integer
// convertable values. If the value is not able to convert integer, it returns
// error with empty string.
func (k *Kvs) Incr(key string) (string, error) {
	k.mu.Lock()
	defer k.mu.Unlock()

	val := k.kv[key]
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return "", err
	}
	k.kv[key] = convStr(valInt + 1)
	return k.kv[key], nil
}

// IncrBy increments the value by given value. If non-integer value is being
// tried to increment, it returns error.
func (k *Kvs) IncrBy(key string, val int) (string, error) {
	k.mu.Lock()
	defer k.mu.Unlock()

	v := k.kv[key]
	valInt, err := strconv.Atoi(v)
	if err != nil {
		return "", err
	}
	k.kv[key] = convStr(valInt + val)
	return k.kv[key], err
}

// convStr converts integer to string.
func convStr(i int) string {
	return strconv.Itoa(i)
}
