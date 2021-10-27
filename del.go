package kvs

func (k *Kvs) Del(key string) {
	delete(k.kv, key)
}