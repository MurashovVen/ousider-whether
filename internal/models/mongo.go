package models

type WhetherConfigurationMDBv1 struct {
	ChatID      int64 `bson:"_id"`
	Temperature int   `bson:"temperature"`
}
