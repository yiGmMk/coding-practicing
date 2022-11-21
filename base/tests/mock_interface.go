package tests

// gen mock代码
// mockgen -source=base/tests/mock_interface.go -destination=base/tests/mock_gen.go -package=tests

// 支持在vscode里直接点击生成,在下方增加go:generate

//go:generate mockgen -source=./mock_interface.go -destination=./mock_gen.go -package=tests Driver
type Driver interface {
	Open(name string) (string, error)
}

var _ Driver = (*defaultDriveer)(nil)

type defaultDriveer struct {
}

func (d defaultDriveer) Open(name string) (string, error) {
	return "success", nil
}
