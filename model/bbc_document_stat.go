package model

type DocumentStat struct {
	//ProjectId					uint64					`json:"project_id"`
	Done						int						`json:"done"`
	Inprogress					int						`json:"inprogress"`
	Rejected					int						`json:"rejected"`

	Sum							int						`json:"sum"`
}

