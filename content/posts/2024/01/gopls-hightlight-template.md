# Gopls高亮template

## 配置

```json
{
    "gopls": {
        "templateExtensions": [
            "tpl",
            "tmpl",
        ],
    },
    "files.associations": {
        "*.tpl": "gotmpl",
        "*.tmpl": "gotmpl"
    }

    // ...
}
```

当文件扩展名为`tpl`, `tmpl`时，均会视为是符合`Go`的`template`文件。

在编辑器里会有变量的高亮和智能提示。
