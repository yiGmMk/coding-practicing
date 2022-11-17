package testsuite

import (
	"fmt"
	"testing"
)

// 包级测试固件
func pkgSetUp(name string) func() {
	fmt.Println("pkg set up:", name)
	return func() {
		fmt.Println("pkg set down:", name)
	}
}

func TestMain(m *testing.M) {
	defer pkgSetUp("testsuite")()
	m.Run()
}

func testBuilder(t *testing.T) {
	fmt.Println("Builder")
}

func testBuilderString(t *testing.T) {
	fmt.Println("BuilderString")
}

func testBuilderReset(t *testing.T) {
	fmt.Println("BuilderReset")
}

func testBuilderGrow(t *testing.T) {
	fmt.Println("BuilderGrow")
}

// 测试用例 测试固件
func builderSetUp(name string) func() {
	fmt.Println("set up:", name)
	return func() {
		fmt.Println("set down:", name)
	}
}

func TestBuilder(t *testing.T) {
	defer builderSetUp(t.Name())()

	t.Run("TestBuilder", testBuilder)
	t.Run("TestBuilderString", testBuilderString)
	t.Run("TestBuilderReset", testBuilderReset)
	t.Run("TestBuilderGrow", testBuilderGrow)
}
