package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/moira-alert/moira/database"
)

// AddTriggersToCheck gets trigger IDs and save it to Redis Set
func (connector *DbConnector) AddTriggersToCheck(triggerIDs []string) error {
	return connector.addTriggersToCheck(localTriggersToCheckKey, triggerIDs)
}

// AddRemoteTriggersToCheck gets remote trigger IDs and save it to Redis Set
func (connector *DbConnector) AddRemoteTriggersToCheck(triggerIDs []string) error {
	return connector.addTriggersToCheck(remoteTriggersToCheckKey, triggerIDs)
}

// GetTriggerToCheck return random trigger ID from Redis Set
func (connector *DbConnector) GetTriggerToCheck() (string, error) {
	return connector.getTriggerToCheck(localTriggersToCheckKey)

}

// GetRemoteTriggerToCheck return random remote trigger ID from Redis Set
func (connector *DbConnector) GetRemoteTriggerToCheck() (string, error) {
	return connector.getTriggerToCheck(remoteTriggersToCheckKey)
}

// GetTriggersToCheckCount return number of triggers ID to check from Redis Set
func (connector *DbConnector) GetTriggersToCheckCount() (int64, error) {
	return connector.getTriggersToCheckCount(localTriggersToCheckKey)
}

// GetRemoteTriggersToCheckCount return number of remote triggers ID to check from Redis Set
func (connector *DbConnector) GetRemoteTriggersToCheckCount() (int64, error) {
	return connector.getTriggersToCheckCount(remoteTriggersToCheckKey)
}

func (connector *DbConnector) addTriggersToCheck(key string, triggerIDs []string) error {
	c := connector.pool.Get()
	defer c.Close()

	c.Send("MULTI")
	for _, triggerID := range triggerIDs {
		c.Send("SADD", key, triggerID)
	}
	_, err := redis.Values(c.Do("EXEC"))
	if err != nil {
		return fmt.Errorf("failed to add triggers to check: %s", err.Error())
	}
	return nil
}

func (connector *DbConnector) getTriggerToCheck(key string) (string, error) {
	c := connector.pool.Get()
	defer c.Close()
	triggerID, err := redis.String(c.Do("SPOP", key))
	if err != nil {
		if err == redis.ErrNil {
			return "", database.ErrNil
		}
		return "", fmt.Errorf("failed to pop trigger to check: %s", err.Error())
	}
	return triggerID, err
}

func (connector *DbConnector) getTriggersToCheckCount(key string) (int64, error) {
	c := connector.pool.Get()
	defer c.Close()
	triggersToCheckCount, err := redis.Int64(c.Do("SCARD", key))
	if err != nil {
		if err == redis.ErrNil {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to get trigger to check count: %s", err.Error())
	}
	return triggersToCheckCount, nil
}

var remoteTriggersToCheckKey = "moira-remote-triggers-to-check"
var localTriggersToCheckKey = "moira-triggers-to-check"
