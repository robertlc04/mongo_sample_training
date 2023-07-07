package structs

type Zip struct {
	_id  string `bson:"_id"`
	city string `bson:"city"`
	zip  string `bson:"zip"`
	loc  struct {
		x float64 `bson:"x"`
		y float64 `bson:"y"`
	} `bson:"loc"`
	pop int `bson:"pop"`
	state string `bson:"state"`
}
