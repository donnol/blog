# Go enum

Go是没有内置枚举类型的，那么，当需要使用枚举时，该怎么办呢？

枚举说白了，就是一连串`互斥的值`，每个值代表一样事物或一个类型。

比如，现在需要一个颜色枚举，可以这样定义：

```go
const (
    Red = "Red" // 红色
    Blue = "Blue" // 蓝色
    Green = "Green" // 绿色
)
```

也有这样定义的：

```go
type Color string // 定义一个特定类型

// 枚举常量均声明为该类型
const (
    Red     Color = "Red" // 红色
    Blue    Color = "Blue" // 蓝色
    Green   Color = "Green" // 绿色
)
```

这样做的好处是可以通过这个类型来更明显的标记出枚举字段来：

```go
type Car struct {
    Name string
    Color Color // 颜色字段声明为Color类型，在阅读代码的时候就能知道这个字段正常的可选值范围
}
```

但是，上面的做法都需要面临一个问题，就是我需要一个返回全部枚举值的集合时，需要这样做：

```go
func All() []Color {
    return []Color{
        Red,
        Blue,
        Green,
    }
}

func (color Color) Name() string {
    switch color {
    case Red:
        return "红色"
    case Blue:
        return "蓝色"
    case Green:
        return "绿色"
    }
    return ""
}
```

当在定义处新增值时，`All`和`Name`也要同步添加，对于开发人员来说，非常容易遗漏。

考虑到枚举值必然是不同的，那在Go里什么东西是必然不同的呢？首先想到的是结构体的字段。

```go
type Color int
var ColorEnumObj struct {
    Enum[int] // 初始化时将每个字段的内容收集起来存到这里

    Blue  EnumField[Color] `enum:"1,blue,蓝色"`
    Green EnumField[Color] `enum:"2,green,绿色"`
    Red   EnumField[Color] `enum:"3,red,红色"`
}
```

通过反射拿到`enum`标签里的内容，并相应的给字段赋值。

```go
func init() {
	panicIf(Init[int](&ColorEnumObj))
}
```

使用：

```go
var (
    _ = ColorEnumObj.Blue.Value()
    _ = ColorEnumObj.Green.Value()
    _ = ColorEnumObj.Red.Value()

    _ = ColorEnumObj.Blue.Name()
    _ = ColorEnumObj.Green.Name()
    _ = ColorEnumObj.Red.Name()

    _ = ColorEnumObj.Blue.ZhName()
    _ = ColorEnumObj.Green.ZhName()
    _ = ColorEnumObj.Red.ZhName()
)
```
