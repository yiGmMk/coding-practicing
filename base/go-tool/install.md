# install

## doc

usage: go install [build flags] [packages]

Install compiles and installs the packages named by the import paths.
编译并安装包

Executables are installed in the directory named by the GOBIN environment
variable, which defaults to $GOPATH/bin or $HOME/go/bin if the GOPATH
environment variable is not set. Executables in $GOROOT
are installed in $GOROOT/bin or $GOTOOLDIR instead of $GOBIN.
默认安装到GoBIn下(GOPATH未设置时默认值是GOPATH/bin或$HOME/go/bin)
在$GOROOT的包安装到$GOROOT/bin或$GOTOOLDIR下

If the arguments have version suffixes (like @latest or @v1.0.0), "go install"
builds packages in module-aware mode, ignoring the go.mod file in the current
directory or any parent directory, if there is one. This is useful for
installing executables without affecting the dependencies of the main module.
To eliminate(消除) ambiguity(歧义) about which module versions are used in the build, the
arguments must satisfy the following constraints:
带版本号时会使用module-aware模式,忽略当前路径的go.mod文件(可以在不影响当前main module的前提下安装包)
为消除歧义(在构建时应当使用哪个module版本),参数需要遵循以下规则

- Arguments must be package paths or package patterns (with "..." wildcards).
They must not be standard packages (like fmt), meta-patterns (std, cmd,
all), or relative or absolute file paths.

- All arguments must have the same version suffix. Different queries are not
allowed, even if they refer to the same version.

- All arguments must refer to packages in the same module at the same version.

- Package path arguments must refer to main packages. Pattern arguments
will only match main packages.

- No module is considered the "main" module. If the module containing
packages named on the command line has a go.mod file, it must not contain
directives (replace and exclude) that would cause it to be interpreted
differently than if it were the main module. The module must not require
a higher version of itself.

- Vendor directories are not used in any module. (Vendor directories are not
included in the module zip files downloaded by 'go install'.)

If the arguments don't have version suffixes, "go install" may run in
module-aware mode or GOPATH mode, depending on the GO111MODULE environment
variable and the presence of a go.mod file. See 'go help modules' for details.
If module-aware mode is enabled, "go install" runs in the context of the main
module.

When module-aware mode is disabled, other packages are installed in the
directory $GOPATH/pkg/$GOOS_$GOARCH. When module-aware mode is enabled,
other packages are built and cached but not installed.

The -i flag installs the dependencies of the named packages as well.
The -i flag is deprecated. Compiled packages are cached automatically.
-i参数已过时(被废弃),编译的包将自动缓存

For more about the build flags, see 'go help build'.
For more about specifying packages, see 'go help packages'.

See also: go build, go get, go clean.
