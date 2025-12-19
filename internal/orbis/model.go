package orbis

type CreateVersionReq struct {
	Version          string   `json:"version"`
	Changelog        string   `json:"changelog"`
	HytaleVersionIDs []string `json:"hytaleVersionIds"`
	IsPreRelease     bool     `json:"isPreRelease"`
}

type SetPrimaryFileReq struct {
	FileID string `json:"fileId"`
}

type Version struct {
	ID string `json:"id"`
}

type VersionFile struct {
	ID string `json:"id"`
}
