package people

type CorePerson struct {
	ID        byte   `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

type ServiceEntities interface {
	Create(input CorePerson) (err error)
}

type RepositoryEntities interface {
	Create(input CorePerson) (err error)
}
