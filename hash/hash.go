package hash

type Hash interface {
	Generate(string) string
	Verify(string, string) bool
}
