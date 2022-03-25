package services

import (
	"context"
	"encoding/json"
	"imagetopdf/data"
	"imagetopdf/helpers"
	"imagetopdf/models"
	"log"
	"os"
	"time"
)

var ctx = context.Background()

func CreateSession() (string, error) {
	var sessionId = helpers.GetGuid()

	var newSession = models.SessionModel{
		SessionId: sessionId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	var sessionObj, err = json.Marshal(newSession)

	if err != nil {
		log.Printf("Cannot create session: Reason %s", err.Error())

		return "", err
	}

	status := data.RedisClient.Set(ctx, sessionId, string(sessionObj), 0)

	res, err := status.Result()

	if err != nil {
		log.Printf("Cannot create session. Reason: %s %s", err.Error(), res)
		return "", err
	}

	return sessionId, nil
}

func UpdateSessionTime(sessionId string) error {
	currentSession, err := GetSession(sessionId)

	if err != nil {
		log.Printf("Cannot update session time. Reason: %s", err.Error())
		return err
	}

	currentSession.UpdatedAt = time.Now()

	currentSessionJson, err := json.Marshal(currentSession)

	if err != nil {
		log.Printf("Cannot parsed session object. Reason: %s", err.Error())
		return err
	}

	data.RedisClient.Set(ctx, sessionId, currentSessionJson, 0)

	return nil
}

func DeleteSession(sessionId string) error {

	err := data.RedisClient.Del(ctx, sessionId).Err()

	if err != nil {
		log.Printf("Cannot delete session %s - Reason: %s", sessionId, err.Error())
		return err
	}

	err = os.RemoveAll(Config.StoragePath + sessionId)

	if err != nil {
		log.Printf("Cannot delete session %s - Reason: %s", sessionId, err.Error())
		return err
	}

	return nil
}

func DeleteAllSessions() error {
	iter := data.RedisClient.Scan(ctx, 0, "", 0).Iterator()

	for iter.Next(ctx) {
		key := iter.Val()

		duration, err := data.RedisClient.TTL(ctx, key).Result()

		if err != nil {
			log.Printf("Cannot delete %s session. Reason: %s", key, err.Error())
			return err
		}

		log.Printf("Key %s \n", key)

		//We only delete a session if it is not active
		if CheckIfSessionIsActive(key) {
			continue
		} else {
			DeleteSession(key)
		}

		if duration == -1 {
			if err := data.RedisClient.Del(ctx, key).Err(); err != nil {
				log.Printf("Cannot delete session. Reason: %s", err.Error())
				return err
			}
		}
	}

	err := iter.Err()

	if err != nil {
		log.Printf("Cannot delete all sessions. Reason: %s", err.Error())
		return err
	}

	return nil
}

func GetSession(sessionId string) (models.SessionModel, error) {
	sessionStatus := data.RedisClient.Get(ctx, sessionId)

	res, err := sessionStatus.Result()

	currentSession := &models.SessionModel{}

	if err != nil {
		log.Printf("Cannot get session. Reason: %s", err.Error())
		return *currentSession, err
	}

	sessionBytes := []byte(res)

	err = json.Unmarshal(sessionBytes, currentSession)

	if err != nil {
		log.Printf("Cannot get session. Reason: %s", err.Error())
		return *currentSession, err
	}

	return *currentSession, nil
}

func CheckIfSessionIsActive(sessionId string) bool {
	currentSession, err := GetSession(sessionId)

	if err != nil {
		return false
	}

	log.Printf("Fecha sesion %s Fecha actual menos 5 minutos %s\n", currentSession.CreatedAt.String(), time.Now().Add(-(time.Minute * 5)).String())

	halfHour := (time.Minute * 5)

	isActive := currentSession.UpdatedAt.After(time.Now().Add(-halfHour))

	return isActive
}

func SessionExists(sessionId string) bool {
	res := data.RedisClient.Exists(ctx, sessionId)
	val := res.Val()
	exists := val != 0
	return exists
}
