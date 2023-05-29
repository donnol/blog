---
author: "jdlau"
date: 2023-05-29
linktitle: flutter Widget Element
menu:
next:
prev:
title: flutter Widget Element
weight: 10
categories: ['flutter', 'Widget', 'Element']
tags: ['flutter']
---

## Element

```dart
abstract class Element extends DiagnosticableTree implements BuildContext
```

> package:flutter/src/widgets/framework.dart
>
> An instantiation of a [Widget] at a particular location in the tree.
> -- 在树里的特定位置上的一个`Widget`的实例。
>
> Widgets describe how to configure a subtree but the same widget can be used to configure multiple subtrees simultaneously because widgets are immutable. An [Element] represents the use of a widget to configure a specific location in the tree. Over time, the widget associated with a given element can change, for example, if the parent widget rebuilds and creates a new widget for this location.
> -- `Widget`描述了如何配置一棵子树，但同一个`Widget`可以被用来配置多棵相似的子树，因为`Widget`是不可变的。一个`Element`代表了一个`Widget`配置在树里的特定位置的使用。随着时间变化，每个`Widget`与一个可以改变的`Element`关联。
>
> Elements form a tree. Most elements have a unique child, but some widgets (e.g., subclasses of [RenderObjectElement]) can have multiple children.

## Widget

```dart
abstract class Widget extends DiagnosticableTree
```

> package:flutter/src/widgets/framework.dart
> 
> Describes the configuration for an [Element].
> -- 对`Element`的配置。
> 
> Widgets are the central class hierarchy in the Flutter framework. A widget is an immutable description of part of a user interface. Widgets can be inflated into elements, which manage the underlying render tree.
> -- 一个`Widget`是一个不可变的描述。
> 
> Widgets themselves have no mutable state (all their fields must be final). If you wish to associate mutable state with a widget, consider using a [StatefulWidget], which creates a [State] object (via [StatefulWidget.createState]) whenever it is inflated into an element and incorporated into the tree.
> -- `Widget`自身没有可变状态，它们的所有字段都必须是final的。如果你想让`Widget`关联可变状态，考虑使用一个`StatefulWidget`，无论它什么时候被放入到一个`element`里，它都会创建一个`State`对象。

## RenderObject

```dart
abstract class RenderObject extends AbstractNode with DiagnosticableTreeMixin implements HitTestTarget
```

> package:flutter/src/rendering/object.dart
>
> An object in the render tree.
> 
> The [RenderObject] class hierarchy is the core of the rendering library's reason for being.
> -- 渲染库的核心
> 
> www.youtube.com/watch?v=zmbmrw07qBc
> 
> [RenderObject]s have a [parent], and have a slot called [parentData] in which the parent [RenderObject] can store child-specific data, for example, the child position. The [RenderObject] class also implements the basic layout and paint protocols.
> 
> The [RenderObject] class, however, does not define a child model (e.g. whether a node has zero, one, or more children). It also doesn't define a coordinate system (e.g. whether children are positioned in Cartesian coordinates, in polar coordinates, etc) or a specific layout protocol (e.g. whether the layout is width-in-height-out, or constraint-in-size-out, or whether the parent sets the size and position of the child before or after the child lays out, etc; or indeed whether the children are allowed to read their parent's [parentData] slot).
> -- 不会定义一个子模型。不会定义一个协作系统。不干实事，只作约束。
> 
> The [RenderBox] subclass introduces the opinion that the layout system uses Cartesian coordinates.
> 
