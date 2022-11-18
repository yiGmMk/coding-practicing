package tests

// gen mock代码
// mockgen -source=base/tests/mock_interface.go -destination=base/tests/mock_gen.go -package=tests
type Driver interface {
	Open(name string) (string, error)
}

var _ Driver = (*defaultDriveer)(nil)

type defaultDriveer struct {
}

func (d defaultDriveer) Open(name string) (string, error) {
	return "success", nil
}
