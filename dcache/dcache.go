package dcache

import (
    "fmt"
    "hash/fnv"
    "os"
    "path/filepath"
    "sync"

    "github.com/hashicorp/golang-lru/v2"
)

const (
    KEY_EXT = ".key"
    DATA_EXT = ".dat"
)

type DiskCache struct {
    path string
    lruCache *lru.Cache[string, string]
    mutex sync.Mutex
}

func (dcache *DiskCache) remove(hash string) {
    filePath := filepath.Join(dcache.path, hash)
    os.Remove(filePath + KEY_EXT)
    os.Remove(filePath + DATA_EXT)
}

func (dcache *DiskCache) Get(key string) []byte {
    dcache.mutex.Lock()
    defer dcache.mutex.Unlock()
    hash, ok := dcache.lruCache.Get(key)
    if !ok {
        return nil
    }
    filePath := filepath.Join(dcache.path, hash)
    data, err := os.ReadFile(filePath + DATA_EXT)
    if err != nil {
        dcache.lruCache.Remove(hash)
        dcache.remove(hash)
        return nil
    }
    return data
}

func (dcache *DiskCache) Add(key string, data []byte) error {
    h := fnv.New64a()
    h.Write([]byte(key))
    hash := fmt.Sprintf("%016x", h.Sum64())
    filePath := filepath.Join(dcache.path, hash)

    dcache.mutex.Lock()
    defer dcache.mutex.Unlock()
    err := os.WriteFile(filePath + DATA_EXT, data, 0644)
    if err != nil {
        dcache.remove(hash)
        return err
    }
    err = os.WriteFile(filePath + KEY_EXT, []byte(key), 0644)
    if err != nil {
        dcache.remove(hash)
        return err
    }
    dcache.lruCache.Add(key, hash)

    return nil
}

func NewDiskCache(path string, size int) (*DiskCache, error) {
    dcache := &DiskCache{}
    dcache.path = path
    files, err := os.ReadDir(path)
    if err != nil {
        return nil, err
    }

    dcache.lruCache, err = lru.NewWithEvict[string, string](size,
        func(key string, value string) {
            dcache.remove(value)
        })

    for _, file := range files {
        fileName := file.Name()
        ext := filepath.Ext(fileName)
        if ext == KEY_EXT {
            filePath := filepath.Join(path, fileName)
            hash := fileName[:len(fileName) - len(ext)]
            if contents, err := os.ReadFile(filePath); err == nil {
                dcache.lruCache.Add(string(contents), hash)
            } else {
                dcache.remove(hash)
            }
        }
    }

    return dcache, nil
}