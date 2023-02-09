package delivery

import "teguh/features/people"

type PersonRequest struct {
	Firstname string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

func PersonRequestToPersonCore(data PersonRequest) people.CorePerson {
	return people.CorePerson{
		Firstname: data.Firstname,
		Lastname:  data.Lastname,
	}
}
