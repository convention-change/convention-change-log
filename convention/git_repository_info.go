package convention

type GitRepositoryHttpInfo struct {
	// Scheme is the protocol scheme of the remote host. https or http.
	// do not use git+ssh, it will be some error
	Scheme string

	// Host is the hostname:port of the remote host.
	Host string

	// Owner is the owner of the repository.
	Owner string

	// Repository is the name of the repository.
	Repository string
}
