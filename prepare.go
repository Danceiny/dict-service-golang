package main

import (
    "fmt"
    . "github.com/Danceiny/dict-service/service"
    . "github.com/Danceiny/go.utils"
    "github.com/go-redis/redis"
    "github.com/jinzhu/gorm"
    "log"
    "os"
)

func Prepare() {
    InitEnv()
    InitDB()
    InitRedis()
    ScanComponent()
}

var (
    client *redis.Client
    db     *gorm.DB
)

var (
    repositoryServiceImplCpt *RepositoryServiceImpl
    redisImplCpt             *RedisImpl
    treeCacheServiceImplCpt  *TreeCacheServiceImpl
    treeServiceImplCpt       *TreeServiceImpl
    baseCrudServiceImplCpt   *BaseCrudServiceImpl
    baseCacheServiceImplCpt  *BaseCacheServiceImpl
)

func InitEnv() {
    _ = os.Setenv("MYSQL_PASSWORD", "1996")
    _ = os.Setenv("MYSQL_DATABASE", "dict")
    _ = os.Setenv("SHOW_SQL", "1")
}

func ScanComponent() {
    repositoryServiceImplCpt = &RepositoryServiceImpl{db}
    redisImplCpt = &RedisImpl{client}
    baseCacheServiceImplCpt = &BaseCacheServiceImpl{redisImplCpt}
    baseCrudServiceImplCpt = &BaseCrudServiceImpl{repositoryServiceImplCpt, baseCacheServiceImplCpt}
    treeCacheServiceImplCpt = &TreeCacheServiceImpl{redisImplCpt, baseCacheServiceImplCpt}
    treeServiceImplCpt = &TreeServiceImpl{repositoryServiceImplCpt, baseCrudServiceImplCpt, treeCacheServiceImplCpt}
    ScanService()
}

func ScanService() {
    AreaServiceImplCpt = &AreaServiceImpl{
        repositoryServiceImplCpt,
        treeServiceImplCpt,
        baseCacheServiceImplCpt}
}

func InitDB() {
    var err error
    db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
        GetEnvOrDefault("MYSQL_USERNAME", "root").(string),
        GetEnvOrDefault("MYSQL_PASSWORD", "root").(string),
        GetEnvOrDefault("MYSQL_ADDR", "127.0.0.1:3306").(string),
        GetEnvOrDefault("MYSQL_DATABASE", "").(string),
        GetEnvOrDefault("MYSQL_CHARSET", "utf8").(string)))
    if err != nil {
        log.Fatalf("connect to mysql failed: %v", err)
    }
    db.LogMode(GetEnvOrDefault("SHOW_SQL", false).(bool))
}

func InitRedis() {
    client = redis.NewClient(&redis.Options{
        Addr:     GetEnvOrDefault("REDIS_ADDR", "127.0.0.1:6379").(string),
        Password: GetEnvOrDefault("REDIS_PASSWORD", "").(string),
        DB:       GetEnvOrDefault("REDIS_DB", 0).(int),
    })
}
