package db

import (
	"github.com/startracex/Auth-example/conf"
)

var CodeMap = new(RedisInstance)

var TokenMap = new(RedisInstance)

var SignDB = new(MongoInstance)

func Init() {

	mongodburi := conf.Global.Database.Mongodb.URI

	mongodbInstances := conf.Global.Database.Mongodb.Instances

	redisuri := conf.Global.Database.Redis.URI

	redisInstances := conf.Global.Database.Redis.Instances

	SignDB.Connect(mongodburi).Use(mongodbInstances["signdb"][0], mongodbInstances["signdb"][1])

	CodeMap.Connect(redisuri, redisInstances["codemap"])

	TokenMap.Connect(redisuri, redisInstances["tokenmap"])

}
