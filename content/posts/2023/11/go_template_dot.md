dot: 在模板里表示为`.`，表示当前作用域。

`{{range}}`, `{{if}}`, `{{with}}`均有自己的作用域。

`{{if pipeline}}`和`{{with pipeline}}`的区别是，前者不会修改`.`的值，而后者会。

## with

with设置`.`的值：

```go
{{with pipeline}} T1 {{end}}
{{with pipeline}} T1 {{else}} T0 {{end}}
```

当pipeline不为0值时，点`.`**设置为pipeline运算的值**，否则跳过。

例如：

```go
{{with "hello"}} {{println .}} {{end}}
```

将输入`hello`，因为`.`被设置为了`hello`.
