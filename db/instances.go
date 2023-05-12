package db

import (
	"fmt"
	"main/config"
)

// code到map的映射连接实例
var CodeMap = new(RedisInstance)

// token到id的映射连接实例
var TokenMap = new(RedisInstance)

// 登录管理的数据库连接实例
var LoginDatabase = new(MongoInstance)

// 注册管理的数据库连接实例
var RegisterDatabase = new(MongoInstance)

var UserDatabase = new(MongoInstance)

func Init() {

	var mongodburi = config.Global.Database.Mongodb.URI

	var mongodbInstances = config.Global.Database.Mongodb.Instances

	var redisuri = config.Global.Database.Redis.URI

	var redisInstances = config.Global.Database.Redis.Instances

	fmt.Println(mongodburi, mongodbInstances, redisuri, redisInstances)

	LoginDatabase.Connect(mongodburi).Use(mongodbInstances["login"][0], mongodbInstances["login"][1])

	RegisterDatabase.Connect(mongodburi).Use(mongodbInstances["register"][0], mongodbInstances["register"][1])

	UserDatabase.Connect(mongodburi).Use(mongodbInstances["user"][0], mongodbInstances["user"][1])

	CodeMap.Connect(redisuri, redisInstances["codemap"])

	TokenMap.Connect(redisuri, redisInstances["tokenmap"])
	
}
