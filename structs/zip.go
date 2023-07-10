package structs

type Zip struct {
	Id   string `bson:"_id" json:"id"`
	City string `bson:"city" json:"city"`
	Zip  string `bson:"zip" json:"zip"`
	Loc  struct {
		X float64 `bson:"x" json:"x"`
		Y float64 `bson:"y" json:"y"`
	} `bson:"loc" json:"loc"`
	Pop 	int 	 `bson:"pop" json:"pop"`
	State string `bson:"state" json:"state"`
}
