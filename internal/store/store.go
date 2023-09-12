package store

import "errors"

type KVStore struct {
	store map[string]string
}

func NewKVStore() *KVStore {
	kvStore := new(KVStore)
	kvStore.store = make(map[string]string)

	return kvStore
}

func (kvStore *KVStore) Put(key string, value string) (string, error) {
	kvStore.store[key] = value

	return value, nil
}

func (kvStore *KVStore) Get(key string) (string, error) {
	if _, ok := kvStore.store[key]; !ok {
		return "", errors.New("key not found")
	}

	return kvStore.store[key], nil
}

func (kvStore *KVStore) Delete(key string) (string, error) {
	delete(kvStore.store, key)
	return key, nil
}
