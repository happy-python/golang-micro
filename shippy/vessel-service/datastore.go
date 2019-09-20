package main

import "gopkg.in/mgo.v2"

// 创建与 MongoDB 交互的主回话
func CreateSession(url string) (session *mgo.Session, err error) {
	session, err = mgo.Dial(url)
	if err != nil {
		return
	}

	session.SetMode(mgo.Monotonic, true)
	return
}

