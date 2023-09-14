package store

type Repo interface {
	Put(key string, value string) (string, error)
	Get(key string) (string, error)
	Delete(key string) (string, error)
}

type KVStore struct {
	repo Repo
}

func NewKVStore(repo Repo) *KVStore {
	kvStore := new(KVStore)
	kvStore.repo = repo

	return kvStore
}

func (kvStore *KVStore) Put(key string, value string) (string, error) {
	return kvStore.repo.Put(key, value)
}

func (kvStore *KVStore) Get(key string) (string, error) {
	return kvStore.repo.Get(key)
}

func (kvStore *KVStore) Delete(key string) (string, error) {
	return kvStore.repo.Delete(key)
}
