package solr

type SolrGroupResponse struct {
	Status int `json:"status"`
	QTime  int `json:"qtime"`
	Params struct {
		Query  string `json:"q"`
		Indent string `json:"indent"`
		Wt     string `json:"wt"`
	} `json:"params"`
	Grouped        map[string]Grouped `json:"grouped"`
	NextCursorMark string             `json:"nextCursorMark"`
	Adds           Adds               `json:"adds"`
}

type Grouped struct {
	Matches int           `json:"matches"`
	Groups  []GroupValues `json:"groups"`
}

type GroupValues struct {
	GroupValue string  `json:"groupValue"`
	DocList    DocList `json:"doclist"`
}

type DocList struct {
	NumFound uint32                   `json:"numFound"`
	Start    int                      `json:"start"`
	Docs     []map[string]interface{} `json:"docs"`
}
