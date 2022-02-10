---
author: "jdlau"
date: 2022-02-10
linktitle: go work
menu:
next:
prev:
title: go work
weight: 10
categories: ['go', 'work']
tags: ['workspace']
---

go1.18将要推出`workspace`模式，此举是为了方便在本地开发多个不同`module`时的依赖管理。

命令说明：

```sh
$ go help work
Go workspace provides access to operations on workspaces.

Note that support for workspaces is built into many other commands, not
just 'go work'.

See 'go help modules' for information about Go\'s module system of which
workspaces are a part.

A workspace is specified by a go.work file that specifies a set of
module directories with the "use" directive. These modules are used as
root modules by the go command for builds and related operations.  A
workspace that does not specify modules to be used cannot be used to do
builds from local modules.

go.work files are line-oriented. Each line holds a single directive,
made up of a keyword followed by arguments. For example:

        go 1.18

        use ../foo/bar
        use ./baz

        replace example.com/foo v1.2.3 => example.com/bar v1.4.5

The leading keyword can be factored out of adjacent lines to create a block,
like in Go imports.

        use (
          ../foo/bar
          ./baz
        )

The use directive specifies a module to be included in the workspace\'s
set of main modules. The argument to the use directive is the directory
containing the module\'s go.mod file.

The go directive specifies the version of Go the file was written at. It
is possible there may be future changes in the semantics of workspaces
that could be controlled by this version, but for now the version
specified has no effect.

The replace directive has the same syntax as the replace directive in a
go.mod file and takes precedence over replaces in go.mod files.  It is
primarily intended to override conflicting replaces in different workspace
modules.

To determine whether the go command is operating in workspace mode, use
the "go env GOWORK" command. This will specify the workspace file being
used.

Usage:

        go work <command> [arguments]

The commands are:

        edit        edit go.work from tools or scripts
        init        initialize workspace file
        sync        sync workspace build list to modules
        use         add modules to workspace file

Use "go help work <command>" for more information about a command.
```

使用`use`指令指定包含在`workspace`里的`module`集。`use`指令后紧接着的是包含了模块的`go.mod`文件的目录--相对`go.work`的目录。

`use`指定的模块被`go`命令用作根模块，执行构建等操作。

`replace`指令与`go.mod`里的用法一样。它会覆盖在`workspace`里不同的模块之间的冲突`replace`。

比如，分别有模块`bar`, `baz`:

它们的`go.mod`分别是：

```gomod
go 1.18

replace foo v0.0.1 => ./foo
```

```gomod
go 1.18

replace foo v0.0.1 => foo v0.0.2
```

可以看到，它们`replace`了不同的`foo`。

```gowork
go 1.18

use (
    ./bar
    ./baz
)

replace foo v0.0.1 => ../foo
```

在`go.work`里统一了对于`foo`的`replace`，均指向到了`../foo`。

不过，虽然多了这个`workspace`模式，依然无法解决`module`版本依赖问题--模块路径和模块版本问题。
