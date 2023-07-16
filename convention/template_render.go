package convention

type CommitRenderTemplate struct {
	GitUrlRenderTemplate
	Hash string `handlebars:"hash"`
}

type CompareRenderTemplate struct {
	GitUrlRenderTemplate
	PreviousTag string `handlebars:"previousTag"`
	CurrentTag  string `handlebars:"currentTag"`
}

type IssueRenderTemplate struct {
	GitUrlRenderTemplate
	Id string `handlebars:"id"`
}

type GitUrlRenderTemplate struct {
	Host       string `handlebars:"host"`
	Owner      string `handlebars:"owner"`
	Repository string `handlebars:"repository"`
}
