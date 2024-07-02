## `vscode`里为不同项目配置不同的`Go`版本

在项目根目录里添加`.vscode/settings.json`，并添加如下内容：

```json
{
    "go.goroot": "/usr/local/go1.22.2",
    "go.alternateTools": {
        "go": "/usr/local/go1.22.2/bin/go"
    },
    "gopls": {
        "build.env": {
            "GOROOT": "/usr/local/go1.22.2"
        }
    }
}
```

即可为该项目配置使用`go1.22.2`版本。
