package main

import (
	"go-mini-kv-store/internal/api"
	inmemoryrepo "go-mini-kv-store/internal/in-memory-repo"
	"go-mini-kv-store/internal/store"
)

func main() {
	// su := os.Getenv("STORAGE_URL")
	var imRepo = inmemoryrepo.NewIMRepo()
	var kvStore = store.NewKVStore(imRepo)
	var kvStoreApi = api.NewApi(kvStore)

	kvStoreApi.R.Run()
}
