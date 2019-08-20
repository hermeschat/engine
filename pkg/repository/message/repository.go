package message

import (
	"context"
	"git.raad.cloud/cloud/hermes/pkg/drivers/mongo"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	MessageID   primitive.ObjectID `bson:"_id" json:"_id"`
	From        string
	To          string
	Time        string
	ChannelID   string
	MessageType string
	Body        string
	Read        bool
}

/*
	case normal_text = 0
    case location = 1
    case contact = 2
    case payment_request = 3
    case payment_response = 4
    case url = 5
    case image = 6
    case video = 7
    case voice = 8
*/
//ConstructFromMap ...
func ConstructFromMap(m map[string]interface{}) (*Message, error) {
	message := &Message{}
	err := mapstructure.Decode(m, message)
	if err != nil {
		return nil, errors.Wrap(err, "can't construct message from given map")
	}
	return message, nil
}

func Get(id string) (*Message, error) {
	s, err := mongo.FindOneById("messages", id)
	if err != nil {
		return nil, errors.Wrap(err, "can't find message with given id")
	}
	message := &Message{}
	err = mapstructure.Decode(s, message)
	if err != nil {
		return nil, errors.Wrap(err, "can't construct message from given map from mongo")
	}
	return message, nil
}

func (s *Message) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{}
	err := mapstructure.Decode(s, &m)
	if err != nil {
		return nil, errors.Wrap(err, "can't create map from this message")
	}
	return m, nil
}
func GetAll(query map[string]interface{}) ([]map[string]interface{}, error) {
	s, err := mongo.FindAll("messages", query)
	if err != nil {
		return nil, errors.Wrap(err, "can't find message with given query")
	}
	//messages := []*Message{}
	messages := []map[string]interface{}{}
	for s.Next(context.Background()) {
		msg := &map[string]interface{}{}
		err = s.Decode(msg)
		if err != nil {
			return nil, errors.Wrap(err, "error in decoding")
		}
		messages = append(messages, *msg)
	}

	return messages, nil
}

//Add adds message
func Add(message *Message) error {

	err := mongo.InsertOne("messages", message)
	if err != nil {
		return errors.Wrap(err, "can't add this message to mongodb")
	}
	return nil
}

func AddAll(messages []Message) error {
	messagesMap := []interface{}{}
	for _, sess := range messages {
		m := map[string]interface{}{}
		err := mapstructure.Decode(sess, m)
		if err != nil {
			return errors.Wrap(err, "can't convert messages to map")
		}
		messagesMap = append(messagesMap, m)
	}
	err := mongo.InsertAll("messages", messagesMap)

	if err != nil {
		return errors.Wrap(err, "can't add this message to mongodb")
	}
	return nil

}
func Delete(id string) error {
	err := mongo.DeleteById("messages", id)
	if err != nil {
		return errors.Wrap(err, "can't delete this message from mongodb")
	}
	return nil
}

func DeleteAll(query map[string]interface{}) error {
	err := mongo.DeleteAllMatched("messages", query)
	if err != nil {
		return errors.Wrap(err, "can't delete with given query from mongo")
	}
	return nil
}

//TODO add count of updated docs
func UpdateOne(id string, query map[string]interface{}) error {
	err := mongo.UpdateOne("messages", id, query)
	if err != nil {
		return errors.Wrap(err, "can't update this id with given query")
	}
	return nil
}

func UpdateAll(selector map[string]interface{}, updator map[string]interface{}) error {
	err := mongo.UpdateAllMatched("messages", selector, updator)
	if err != nil {
		return errors.Wrap(err, "can't update message with given query")
	}
	return nil
}
