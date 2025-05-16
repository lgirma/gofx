package licensing

type Vault interface {
	Store(string) error
	Read() (string, error)
}
