package inmemoryrepo

import "errors"

type IMRepo struct {
	store map[string]string
}

func NewIMRepo() *IMRepo {
	imRepo := new(IMRepo)
	imRepo.store = make(map[string]string)

	return imRepo
}

func (imRepo *IMRepo) Put(key string, value string) (string, error) {
	if _, ok := imRepo.store[key]; ok {
		return "", errors.New("key already exists")
	}
	imRepo.store[key] = value

	return value, nil
}

func (imRepo *IMRepo) Get(key string) (string, error) {
	if _, ok := imRepo.store[key]; !ok {
		return "", errors.New("key not found")
	}

	return imRepo.store[key], nil
}

func (imRepo *IMRepo) Delete(key string) (string, error) {
	delete(imRepo.store, key)
	return key, nil
}
