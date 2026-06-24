package sdk

import "errors"

// WorkspaceModule describes the SDK module a workspace should install for a
// child module runtime.
type WorkspaceModule struct {
	Name   string
	Source string
}

// WorkspaceModuleForRuntime returns the workspace module that exposes a
// built-in runtime SDK. Unknown external SDK refs are intentionally left for
// the normal SDK loader and do not have a static workspace module mapping here.
func WorkspaceModuleForRuntime(runtime string) (WorkspaceModule, bool, error) {
	sdkName, suffix, err := parseSDKName(runtime)
	if errors.Is(err, errUnknownBuiltinSDK) {
		return WorkspaceModule{}, false, nil
	}
	if err != nil {
		return WorkspaceModule{}, false, err
	}

	mod, ok := workspaceModuleForBuiltinSDK(sdkName, suffix)
	return mod, ok, nil
}

func workspaceModuleForBuiltinSDK(sdkName sdk, suffix string) (WorkspaceModule, bool) {
	switch sdkName {
	case sdkGo:
		return WorkspaceModule{Name: "go-sdk", Source: "github.com/dagger/go-sdk@adf60252445825c0d1fa2849f69eefe230d7f154"}, true
	case sdkDang:
		return WorkspaceModule{Name: "dang-sdk", Source: "github.com/dagger/dang-sdk@a21728f312f3595e067b715d2a93ec4dbb16fb88"}, true
	case sdkPython:
		return WorkspaceModule{Name: "python-sdk", Source: "github.com/dagger/python-sdk@da2ed282fd1cffda8a5ef7242b2aa562182b5e4f"}, true
	case sdkTypescript:
		return WorkspaceModule{Name: "typescript-sdk", Source: "github.com/dagger/typescript-sdk@19dd24aabefaf25f49e21a232eb5b72103627cea"}, true
	case sdkJava:
		return WorkspaceModule{Name: "java-sdk", Source: "github.com/dagger/dagger/sdk/java" + suffix}, true
	case sdkPHP:
		return WorkspaceModule{Name: "php-sdk", Source: "github.com/dagger/dagger/sdk/php" + suffix}, true
	case sdkElixir:
		return WorkspaceModule{Name: "elixir-sdk", Source: "github.com/dagger/dagger/sdk/elixir" + suffix}, true
	default:
		return WorkspaceModule{}, false
	}
}
