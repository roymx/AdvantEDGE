/*
 * Copyright (c) 2019  InterDigital Communications, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package redisdb

import (
	"errors"
	"strings"
	"time"

	log "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger"

	"github.com/KromDaniel/rejonson"
	"github.com/go-redis/redis"
)

const dbMaxRetryCount = 2

// Connector - Implements a Redis connector
type Connector struct {
	addr          string
	table         int
	connected     bool
	client        *rejonson.Client
	pubsub        *redis.PubSub
	isListening   bool
	doneListening chan bool
}

// NewConnector - Creates and initialize a Redis connector
func NewConnector(addr string, table int) (rc *Connector, err error) {
	rc = new(Connector)

	// Connect to Redis DB
	for retry := 0; !rc.connected && retry <= dbMaxRetryCount; retry++ {
		err = rc.connectDB(addr, table)
		if err != nil {
			log.Warn("Failed to connect to DB. Retrying... Error: ", err)
		}
	}
	if err != nil {
		return nil, err
	}

	log.Info("Successfully connected to DB")
	return rc, nil
}

func (rc *Connector) connectDB(addr string, table int) error {
	if addr == "" {
		rc.addr = "meep-redis-master:6379"
	} else {
		rc.addr = addr
	}
	rc.table = table
	log.Debug("Redis Connector connecting to ", rc.addr)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     rc.addr,
		Password: "",    // no password set
		DB:       table, // 0 is default DB
	})
	rc.client = rejonson.ExtendClient(redisClient)

	pong, err := rc.client.Ping().Result()

	if pong == "" || err != nil {
		log.Error("Redis Connector unable to connect ", rc.addr)
		return err
	}

	rc.connected = true
	log.Info("Redis Connector connected to ", rc.addr)
	return nil
}

// DBFlush - Empty DB
func (rc *Connector) DBFlush(module string) error {
	var cursor uint64
	var err error
	log.Debug("DBFlush module: ", module)

	// Find all module keys
	// Process in chunks of 50 matching entries to optimize processing speed & memory
	keyMatchStr := module + ":*"
	for {
		var keys []string
		keys, cursor, err = rc.client.Scan(cursor, keyMatchStr, 50).Result()
		if err != nil {
			log.Debug("ERROR: ", err)
			break
		}

		// Delete all matching entries
		if len(keys) > 0 {
			_, err = rc.client.Del(keys...).Result()
			if err != nil {
				log.Debug("Failed to retrieve entry fields")
				break
			}
		}

		// Stop searching if cursor is back at beginning
		if cursor == 0 {
			break
		}
	}

	return nil
}

// EntryExists - true if entry exists; false otherwise
func (rc *Connector) EntryExists(key string) bool {
	value := rc.client.Exists(key).Val()
	return value != 0
}

// // DBForEachEntry - Search for matching keys and run handler for each entry
func (rc *Connector) ForEachEntry(keyMatchStr string, entryHandler func(string, map[string]string, interface{}) error, userData interface{}) error {
	var cursor uint64
	var err error

	// Process in chunks of 50 matching entries to optimize processing speed & memory
	for {
		var keys []string
		keys, cursor, err = rc.client.Scan(cursor, keyMatchStr, 50).Result()
		if err != nil {
			log.Debug("ERROR: ", err)
			break
		}

		if len(keys) > 0 {
			for i := 0; i < len(keys); i++ {
				fields, err := rc.client.HGetAll(keys[i]).Result()
				if err != nil || fields == nil {
					log.Debug("Failed to retrieve entry fields")
					break
				}

				// Invoke handler to process entry
				err = entryHandler(keys[i], fields, userData)
				if err != nil {
					return err
				}
			}
		}

		// Stop searching if cursor is back at beginning
		if cursor == 0 {
			break
		}
	}
	return nil
}

func (rc *Connector) ForEachJSONEntry(keyMatchStr string, param1 string, param2 string, entryHandler func(string, string, string, string, interface{}) error, userData interface{}) error {
	var cursor uint64
	var err error

	// Process in chunks of 50 matching entries to optimize processing speed & memory
	for {
		var keys []string
		keys, cursor, err = rc.client.Scan(cursor, keyMatchStr, 50).Result()
		if err != nil {
			log.Debug("ERROR: ", err)
			break
		}
		if len(keys) > 0 {
			for i := 0; i < len(keys); i++ {
				jsonInfo, err := rc.client.JsonGet(keys[i], ".").Result()
				if err != nil || jsonInfo == "" {
					log.Debug("Failed to retrieve entry fields")
					break
				}

				// Invoke handler to process entry
				err = entryHandler(keys[i], jsonInfo, param1, param2, userData)
				if err != nil {
					return err
				}
			}
		}

		// Stop searching if cursor is back at beginning
		if cursor == 0 {
			break
		}
	}
	return nil
}

// SetEntry - Update existing entry or create new entry if it does not exist
func (rc *Connector) SetEntry(key string, fields map[string]interface{}) error {
	if !rc.connected {
		return errors.New("Redis Connector is disconnected (SetEntry)")
	}
	// Update existing entry or create new entry if it does not exist
	_, err := rc.client.HMSet(key, fields).Result()
	if err != nil {
		return err
	}
	return nil
}

// DelEntry - delete an existing entry from DB
func (rc *Connector) DelEntry(key string) error {
	if !rc.connected {
		return errors.New("Redis Connector is disconnected (DelEntry)")
	}
	// Delete entry if it exists
	_, err := rc.client.Del(key).Result()
	if err != nil {
		return err
	}
	return nil
}

// JSONGetEntry - Retrieve entry from DB
func (rc *Connector) JSONGetEntry(key string, path string) (string, error) {
	if !rc.connected {
		return "", errors.New("Redis Connector is disconnected (JSONGetEntry)")
	}
	// Retreive JSON entry if it exists
	json, err := rc.client.JsonGet(key, path).Result()
	if err != nil {
		return "", err
	}
	return json, nil
}

// JSONGetList -
func (rc *Connector) JSONGetList(elem1 string, elem2 string, elementPath string, entryHandler func(string, string, string, string, interface{}) error, dataList interface{}) error {
	if !rc.connected {
		return errors.New("Redis Connector is disconnected (JSONGetList)")
	}

	keyName := elementPath + "*"
	err := rc.ForEachJSONEntry(keyName, elem1, elem2, entryHandler, dataList)
	if err != nil {
		log.Error("keyName: ", keyName, ": ", err.Error())
		return err
	}
	return nil
}

// JSONSetEntry - update existing entry from DB or create a new one if it doesnt't exist
func (rc *Connector) JSONSetEntry(key string, path string, json string) error {
	if !rc.connected {
		return errors.New("Redis Connector is disconnected (JSONSetEntry)")
	}
	// Update existing entry or create new entry if it does not exist
	_, err := rc.client.JsonSet(key, path, json).Result()
	if err != nil {
		log.Error("key: ", key, ": ", err.Error())
		return err
	}
	return nil
}

// JSONDelEntry - delete an existing entry from DB
func (rc *Connector) JSONDelEntry(key string, path string) error {
	if !rc.connected {
		return errors.New("Redis Connector is disconnected (JSONDelEntry)")
	}
	// Update existing entry or create new entry if it does not exist
	_, err := rc.client.JsonDel(key, path).Result()
	if err != nil {
		return err
	}
	return nil
}

// Subscribe - Register as a listener for provided channels
func (rc *Connector) Subscribe(channels ...string) error {
	if !rc.connected {
		return errors.New("Redis Connector is disconnected (Subscribe)")
	}

	rc.pubsub = rc.client.Subscribe(channels...)
	return nil
}

// Unsubscribe - Unregister as a listener for provided channels
func (rc *Connector) Unsubscribe(channels ...string) error {
	if !rc.connected {
		return errors.New("Redis Connector is disconnected (Unsubscribe)")
	}
	if rc.pubsub != nil {
		_ = rc.pubsub.Unsubscribe(channels...)
	}
	return nil
}

// Listen - Wait for subscribed events
func (rc *Connector) Listen(handler func(string, string)) error {
	if !rc.connected {
		return errors.New("Redis Connector is disconnected (Listen)")
	}
	if rc.pubsub == nil {
		return errors.New("Not subscribed to pubsub (Listen)")
	}

	rc.isListening = true
	rc.doneListening = make(chan bool, 1)
	// Main listening loop
	for {
		// Wait for subscribed channel events, or timeout
		msg, err := rc.pubsub.ReceiveTimeout(time.Second)
		if err != nil {
			if !strings.Contains(err.Error(), "timeout") {
				log.Debug("Listen Error: ", err)
			}
		} else {
			channel := ""
			payload := ""

			// Process published event
			switch m := msg.(type) {
			// Process Subscription
			case *redis.Subscription:
				log.Info("Subscription Message: ", m.Kind, " to channel ", m.Channel, ". Total subscriptions: ", m.Count)
				continue
			// Process received Message
			case *redis.Message:
				channel = m.Channel
				payload = m.Payload
				log.Trace("RX-MSG [", channel, "] ", payload)
				handler(channel, payload)
			}
		}

		if !rc.isListening {
			log.Debug("Redis Connector exiting listen routine")
			rc.doneListening <- true
			return nil
		}
	}
}

// StopListen - Stop the listening goroutine
func (rc *Connector) StopListen() {
	if rc.isListening {
		// stop the listen goroutine
		rc.isListening = false
		// synchronize on completion
		<-rc.doneListening
	}
}

// Publish - Publish message to channel
func (rc *Connector) Publish(channel string, message string) error {
	if !rc.connected {
		return errors.New("Redis Connector is disconnected (Publish)")
	}

	log.Trace("TX-MSG [", channel, "] ", message)
	_, err := rc.client.Publish(channel, message).Result()
	return err
}
