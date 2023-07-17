package convention

type GitRepositoryInfo struct {
	// Scheme is the protocol scheme of the remote host. like https or http
	Scheme string

	// Host is the hostname:port of the remote host.
	Host string

	// Owner is the owner of the repository.
	Owner string

	// Repository is the name of the repository.
	Repository string
}
