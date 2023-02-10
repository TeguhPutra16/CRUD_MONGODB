package repository

import (
	"teguh/features/people"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

func FromCoreToModel(dataCore people.CorePerson) Person { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	Person := Person{
		Firstname: dataCore.Firstname,
		Lastname:  dataCore.Lastname,
	} ///formating data berdasarkan data gorm dan kita mapping data yang kita butuhkan untuk inputan  klien
	return Person //insert user
}

func (dataModel *Person) ModelToCore() people.CorePerson {
	var Core = people.CorePerson{
		ID:        dataModel.ID,
		Firstname: dataModel.Firstname,
		Lastname:  dataModel.Lastname,
	}
	return Core
}

func ListModelToCore(dataModel []Person) []people.CorePerson {
	var dataCore []people.CorePerson

	for _, v := range dataModel {
		dataCore = append(dataCore, v.ModelToCore())

	}
	return dataCore

}
