package convention

// CommitRenderTemplate
// default template is DefaultCommitUrlFormat
type CommitRenderTemplate struct {
	GitUrlRenderTemplate
	Hash string `handlebars:"hash"`
}

// CompareRenderTemplate
// default template is DefaultCompareUrlFormat
type CompareRenderTemplate struct {
	GitUrlRenderTemplate
	PreviousTag string `handlebars:"previousTag"`
	CurrentTag  string `handlebars:"currentTag"`
}

// IssueRenderTemplate
// default template is DefaultIssueUrlFormat
type IssueRenderTemplate struct {
	GitUrlRenderTemplate
	Id string `handlebars:"id"`
}

// GitUrlRenderTemplate
// default template is DefaultUserUrlFormat
type GitUrlRenderTemplate struct {
	Scheme     string `handlebars:"scheme"`
	Host       string `handlebars:"host"`
	Owner      string `handlebars:"owner"`
	Repository string `handlebars:"repository"`
}

// ReleaseCommitMessageRenderTemplate
// default template is DefaultReleaseCommitMessageFormat
type ReleaseCommitMessageRenderTemplate struct {
	CurrentTag string `handlebars:"currentTag"`
}
