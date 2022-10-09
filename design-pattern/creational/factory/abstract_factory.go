package designpattern

// 抽象工厂
// IConfigParserFactory 工厂方法接口
type IConfigParserFactory interface {
	CreateParser() IConfigParser
}

// yamlParserFactory yamlRuleConfigParser 的工厂类
type yamlParserFactory struct {
}

// CreateParser CreateParser
func (y yamlParserFactory) CreateParser() IConfigParser {
	return yamlParser{}
}

// jsonParserFactory jsonRuleConfigParser 的工厂类
type jsonParserFactory struct {
}

// CreateParser CreateParser
func (j jsonParserFactory) CreateParser() IConfigParser {
	return jsonParser{}
}

// NewIRuleConfigParserFactory 用一个简单工厂封装工厂方法
func NewIRuleConfigParserFactory(t string) IConfigParserFactory {
	switch t {
	case "json":
		return jsonParserFactory{}
	case "yaml":
		return yamlParserFactory{}
	}
	return nil
}
