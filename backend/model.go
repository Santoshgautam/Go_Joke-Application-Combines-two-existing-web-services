package backend

type Person struct {
	Gender  string `json:"gender" bson:"gender"`
	Name    string `json:"name" bson:"name"`
	Surname string `json:"surname" bson:"surname"`
	Region  string `json:"region" bson:"region"`
}

type JokeObj struct {
	Type  string `json:"type"`
	Value struct {
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}
