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
	Tags        []string `bson:"tags"`
	Stars       []string `bson:"stars"`
	MagnetLinks []string `bson:"magnetLinks"`
}

// Stars ...
type Stars struct {
	Name   string
	Age    int
	Avatar string
	Videos []string
}

// Failed ...
type Failed struct {
	Reason string `bson:"reason"`
	Part   string `bson:"part"`
	No     string `bson:"no"`
}
