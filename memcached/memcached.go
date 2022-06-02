package memcached

import (
	"bytes"
	"encoding/gob"
	"os"
	"strconv"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/j4real2208/golang-db/directory"
	newError "github.com/j4real2208/golang-db/error"
)



var logger = newError.Getlogger()

type Client struct {
	client *memcache.Client
}

func IntializeMEM() (*Client , error){
	client := memcache.New(os.Getenv("MEMCACHED"))

	if err := client.Ping(); err != nil {
		return nil, err
	}
	
	client.Timeout = 100 * time.Millisecond
	client.MaxIdleConns = 100
	logger.Info("Memcache connection established .. ")
	return &Client{
		client: client,
	},nil
}


func (c *Client) GetName(id string) (directory.Directory , error ) {
	item , err := c.client.Get(id)
	if err != nil {		
		return directory.Directory{}, err
	}
	b := bytes.NewReader(item.Value)

	var res directory.Directory

	if err := gob.NewDecoder(b).Decode(&res); err != nil {		
		return directory.Directory{}, err
	}	

	return res, nil
}

func (c *Client) SetName(d directory.Directory) error {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(d); err != nil {
		return err
	}

	return c.client.Set(&memcache.Item{
		Key:        strconv.Itoa(int(d.Customer_id)),
		Value:      b.Bytes(),
		Expiration: int32(time.Now().Add(25 * time.Second).Unix()),
	})
}