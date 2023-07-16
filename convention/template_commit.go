package convention

type CommitRenderTemplate struct {
	Host       string `handlebars:"host"`
	Owner      string `handlebars:"owner"`
	Repository string `handlebars:"repository"`
	Hash       string `handlebars:"hash"`
}
