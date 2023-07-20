package query

import (
	"main/db"
	"main/interface"
)

type A = map[string]any

type S interface {
	QueryUser(intf.SignInput)
}

func QueryUser(filter A) map[string]any {
	return db.SignDB.Find(filter)
}

func UserExist(filter A) bool {
	return len(db.SignDB.Find(filter)) != 0
}

func AddUser(data A) any {
	return db.SignDB.Add(data)
}
