package config

import (
    "os"
    "path/filepath"
    "sync"

    "gopkg.in/yaml.v3"
)


type Config struct {
    values map[interface{}]interface{}
    path string
    mutex sync.Mutex
}

func (config *Config) Get(key string, defvalue interface{}) interface{} {
    config.mutex.Lock()
    defer config.mutex.Unlock()
    value, ok := config.values[key]
    if !ok {
        return defvalue
    }
    return value
}

func (config *Config) Set(key string, value interface{}) {
    config.mutex.Lock()
    defer config.mutex.Unlock()
    config.values[key] = value
}

func (config *Config) Save() error {
    config.mutex.Lock()
    defer config.mutex.Unlock()
    data, err := yaml.Marshal(&config.values)
    if err != nil {
        return err
    }
    err = os.WriteFile(config.path, data, 0644)
    if err != nil {
        return err
    }
    return nil
}

func NewConfig(path string) (*Config, error){
    config := &Config{
        values: make(map[interface{}]interface{}),
    }

    config.path = filepath.Join(path, "config")
    if file, err := os.ReadFile(config.path); err == nil {
        err = yaml.Unmarshal(file, &config.values)
        if err != nil {
            return nil, err
        }
    }

    return config, nil
}
