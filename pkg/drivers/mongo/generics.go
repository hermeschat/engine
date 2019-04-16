package mongo

import (
	"context"

	"github.com/pkg/errors"
)

//InsertOne insert a new document in db
func InsertOne(collName string, m map[string]interface{}) error {
	coll, err := GetCollection(collName)
	if err != nil {
		return errors.Wrap(err, "could not get collection ")
	}
	_, err = coll.InsertOne(context.Background(), m)
	if err != nil {
		return errors.Wrap(err, "could not insert a new document")
	}
	return nil
}

//InsertAll inserts given array of maps to mongoDB
func InsertAll(collName string, m []interface{}) error {
	coll, err := GetCollection(collName)
	if err != nil {
		return errors.Wrap(err, "could not get collection ")
	}
	_, err = coll.InsertMany(context.Background(), m)
	if err != nil {
		return errors.Wrap(err, "could not insert a new document")
	}
	return nil
}

//DeleteById removes a document with given Id
func DeleteById(collName string, id string) error {
	coll, err := GetCollection(collName)
	if err != nil {
		return errors.Wrap(err, "could not get collection ")
	}
	filter := map[string]string{
		"_id": id,
	}
	_, err = coll.DeleteOne(context.Background(), filter)
	if err != nil {
		return errors.Wrap(err, "could not deleteById ")
	}
	return nil
}

//DeleteAllMatched removed all matched documents
func DeleteAllMatched(collName string, filter map[string]interface{}) error {
	coll, err := GetCollection(collName)
	if err != nil {
		return errors.Wrap(err, "could not get collection ")
	}

	_, err = coll.DeleteMany(context.Background(), filter)
	if err != nil {
		return errors.Wrap(err, "could not deleteById ")
	}
	return nil
}

//UpdateAllMatched updates all matched documents
func UpdateAllMatched(collName string, query map[string]interface{}, data map[string]interface{}) error {
	coll, err := GetCollection(collName)
	if err != nil {
		return errors.Wrap(err, "could not get collection ")
	}
	_, err = coll.UpdateOne(context.Background(), query, data)
	if err != nil {
		return errors.Wrap(err, "could not updateAllMatched")
	}
	return nil
}

//UpdateOne updates document with given id and data
func UpdateOne(collName string, id string, data map[string]interface{}) error {
	coll, err := GetCollection(collName)
	if err != nil {
		return errors.Wrap(err, "could not get collection ")
	}
	filter := map[string]string{
		"_id": id,
	}
	_, err = coll.UpdateOne(context.Background(), filter, data)
	if err != nil {
		return errors.Wrap(err, "could not updateAllMatched")
	}
	return nil
}

func FindAll(collName string, query map[string]interface{}) ([]map[string]interface{}, error) {
	coll, err := GetCollection(collName)
	if err != nil {
		return nil, errors.Wrap(err, "could not get collection ")
	}

	res := coll.FindOne(context.Background(), query)
	if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "mongo find err")
	}
	output := []map[string]interface{}{}

	err = res.Decode(output)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse mongoSingleResult As map")
	}
	return output, nil
}
func FindOneById(collName string, id string) (map[string]interface{}, error) {
	coll, err := GetCollection(collName)
	if err != nil {
		return nil, errors.Wrap(err, "could not get collection ")
	}
	filter := map[string]string{
		"_id": id,
	}
	res := coll.FindOne(context.Background(), filter)
	if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "mongo find err")
	}
	output := map[string]interface{}{}

	err = res.Decode(output)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse mongoSingleResult As map")
	}
	return output, nil
}
