package main

import (
	"go-mini-kv-store/internal/api"
	"go-mini-kv-store/internal/store"
)

func main() {
	var kvStore = store.NewKVStore()
	var kvStoreApi = api.NewApi(kvStore)

	kvStoreApi.R.Run()
}
