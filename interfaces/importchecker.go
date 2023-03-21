package interfaces

type ImportChecker interface {
	CheckImports(*CheckImportRequest) error
}

type CheckImportRequest struct {
	PackagePath string
	ImportPaths []string
}
