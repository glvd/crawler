package schema

// Video ...
type Video struct {
	No          string   `bson:"no"`
	Thumb       string   `bson:"thumb"`
	Cover       string   `bson:"cover"`
	Date        string   `bson:"date"`
	Length      string   `bson:"length"`
	Producer    string   `bson:"producer"`
	Publisher   string   `bson:"publisher"`
	Director    string   `bson:"director"`
	Series      string   `bson:"series"`
	Tag         []string `bson:"tags"`
	Actress     []string `bson:"actress"`
	MagnetLinks []string `bson:"magnetLinks"`
}

// Actress ...
type Actress struct {
	Name   string
	Age    int
	Avatar string
	Videos []string
}
