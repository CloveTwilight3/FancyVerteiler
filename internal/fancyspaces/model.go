package fancyspaces

type CreateVersionReq struct {
	Name                      string   `json:"name"`
	Platform                  string   `json:"platform"`
	Channel                   string   `json:"channel"`
	Changelog                 string   `json:"changelog"`
	SupportedPlatformVersions []string `json:"supported_platform_versions"`
}
