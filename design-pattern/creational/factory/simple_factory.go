package designpattern

// 简单工厂:Go 本身是没有构造函数的
// 采用 NewName 的方式创建对象/接口，当它返回的是接口的时候，其实就是简单工厂模式
// IConfigParser IConfigParser
type IConfigParser interface {
	Parse(data []byte)
}

// jsonParser jsonParser
type jsonParser struct {
}

// Parse Parse
func (J jsonParser) Parse(data []byte) {
	panic("implement me")
}

// yamlParser yamlParser
type yamlParser struct {
}

// Parse Parse
func (Y yamlParser) Parse(data []byte) {
	panic("implement me")
}

// NewIConfigParser NewIConfigParser
func NewIConfigParser(t string) IConfigParser {
	switch t {
	case "json":
		return jsonParser{}
	case "yaml":
		return yamlParser{}
	}
	return nil
}
