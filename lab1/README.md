# Lab 1: 使用 Go 语言完成第一个数据结构链表（Linked List）

## 0x00 前置知识

- Go 语言基础
  - struct 定义
  - func 的定义及使用
  - func 的 Receiver 相关
- Go 语言指针基础
  - nil 指针
  - 指针的使用

## 0x01 实验介绍

链表是一个非常基础的数据结构，它的实现也是很简单的，但是它的应用却非常广泛，比如在 Linux 内核中，链表的使用非常广泛，比如进程链表、文件链表等等，所以我们有必要学习链表的实现。

其通过一组不连续的内存空间来存储数据，通过指针将这些内存空间串联起来，形成一个链表，链表的每个节点都包含了数据和指向下一个节点的指针。

看上去它很复杂，但是却十分容易。你需要实现一个链表以实现

## 0x02 实验内容

### 1. 节点（Node）

> task_id: node  
> file: [src/node/node.g0](src/node/node.g0)
> 
> 请完成 `src/node/node.g0` 中的代码，使得 `Node` 能够通过测试。


链表的节点使用 `Node` 结构体来表示，它包含了两个字段：
- val 表示值
- next 表示下一个节点

其用 Graph 图来表示为

```
+------+           +------+
| Node |    +----> | Node |
+------+    |      +------+
| val  |    |      | val  |
| next +----+      | next +----> Nil
+------+           +------+
```

当 Next 为空指针时，我们会发现其为结尾，否之则不是。

### 2. 链表（LinkedList）

> task_id: linkedlist  
> file: [src/linkedlist/linkedlist.g0](src/linkedlist/linkedlist.g0)

链表是对节点（Node）的高级封装，其支持了增删查改等操作。我们的练习会关注于增删查操作。

#### 增（Add）

链表的增有很多种，我们只关注一种：

考虑一个链表：

```
{a} -> [b] -> [c] -> nil
{} 表示 LinkedList 的 Head Node
[] 表示 Node
```

如果我们需要增加节点 d，我们可以：

```
CREATE [d]
MAKE [d] -> {a}
// 那么现在变成了 [d] -> {a} -> [b] -> [c] -> nil
MAKE {d} -> [a] -> [b] -> [c] -> nil
```

### 查（Has）

链表的查询依赖于迭代，尝试使用上面的例子来思考。

### 删（Remove）

去删除 节点 b
```
{d} -> [a] -> [b] -> [c] -> nil

FIND [b] and its prev.
prev := [a]
current := [b]
prev.next = current.next
          +-----------+
          |           v
{d} -> [a]+   [b] -> [c] -> nil
        ^      ^
      prev     current
```