package crawler_test

type listItem struct {
	thumbnail string
	no        string
}

type video struct {
	no       string
	thumb    string
	cover    string
	date     string
	length   string
	director string
	actress  []string
}

const (
	url = "https://www.seedmm.life"
)
