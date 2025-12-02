package modrinth

type CreateVersionReq struct {
	Name          string              `json:"name"`
	VersionNumber string              `json:"version_number"`
	Changelog     string              `json:"changelog"`
	Dependencies  []ProjectDependency `json:"dependencies"`
	GameVersions  []string            `json:"game_versions"`
	VersionType   string              `json:"version_type"`
	Loaders       []string            `json:"loaders"`
	Featured      bool                `json:"featured"`
	Status        string              `json:"status"`
	ProjectID     string              `json:"project_id"`
	FileParts     []string            `json:"file_parts"`
	PrimaryFile   string              `json:"primary_file"`
}

type ProjectDependency struct {
	VersionID      string `json:"version_id"`
	ProjectID      string `json:"project_id"`
	FileName       string `json:"file_name"`
	DependencyType string `json:"dependency_type"`
}
