package mgo

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

type Client struct {
	database   string
	collection string
	session    *mgo.Session
}

func New(database, collection string) *Client {
	di := &mgo.DialInfo{
		Addrs:    []string{"ds019028.mlab.com:19028"},
		Timeout:  60 * time.Second,
		Database: database,
		Username: "local_dev",
		Password: "local_dev",
	}
	session, err := mgo.DialWithInfo(di)
	if err != nil {
		log.Fatalf("Failed to create mongoDB session. Err: %v", err)
	}

	// Consider making this Monotonic. Don't think this matters as much for now
	// since there's only 1 backend anyway.
	session.SetMode(mgo.Strong, true)

	return &Client{database, collection, session}
}

func (c *Client) C() *mgo.Collection {
	return c.Collection()
}

func (c *Client) Collection() *mgo.Collection {
	return c.session.DB(c.database).C(c.collection)
}
