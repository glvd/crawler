package schema

// Video ...
type Video struct {
	No          string   `bson:"no"`
	Title       string   `bson:"title"`
	Thumb       string   `bson:"thumb"`
	Cover       string   `bson:"cover"`
	Date        string   `bson:"date"`
	Length      string   `bson:"length"`
	Producer    string   `bson:"producer"`
	Publisher   string   `bson:"publisher"`
	Director    string   `bson:"director"`
	Series      string   `bson:"series"`
	Tags        []string `bson:"tags"`
	Stars       []string `bson:"stars"`
	MagnetLinks []string `bson:"magnetLinks"`
	Uncensored  bool     `bson:"uncensored"`
}

// Stars ...
type Star struct {
	Name       string `bson:"name"`
	Birthday   string `bson:"birthday"`
	Age        string `bson:"age"`
	Avatar     string `bson:"avatar"`
	Height     string `bson:"height"`
	Cup        string `bson:"cup"`
	Chest      string `bson:"chest"`
	Waist      string `bson:"waist"`
	Hipline    string `bson:"hipline"`
	BirthPlace string `bson:"birthPlace"`
	Hobby      string `bson:"hobby"`
	Uncensored bool   `bson:"uncensored"`
}

// Failed ...
type Failed struct {
	Reason string `bson:"reason"`
	Part   string `bson:"part"`
	No     string `bson:"no"`
}
