# Lab 1：源代码目录

## 目录格式

所有的工作文件都是以如下格式组织的

```
- src
  +- {task_id}
  |  +- {code}.g0
  |
  +- {task_id}_check
     +- main.go
     +- ...
```

目录分为两种：

- `task_id`
- `task_id_check`

`task_id`表示任务id，你需要完成`task_id`文件夹内的代码。

`task_id_check`是测试用例，你不应该修改其中的代码。当你完成`task_id` 的代码后，请运行 `task_id_check` 的 `main.go` 中的 `main` 函数，以查看你是否完成了该任务。

举个例子，我有个任务叫 `node`，那么我需要完成 `node` 文件夹中所有的代码任务。在完成任务后，我会去 `node_check` 执行代码测试。

## 完成练习

所有Go的练习源码会以 `.g0` （dot G Zero）为拓展名保存。这样做是为了防止当完成练习时，错误的修改了原文件而忘记应该做什么。

正常的完成练习的流程是：
1. 复制 `{code}.g0` 到 `{code.go}`
2. 完成 `{code}.go`
3. 如果代码混乱了，那么可以重复1 恢复

例如在 `node` 任务中，我有如下文件

- node
  - node.g0 <-- 需要完成
  - node2.g0 <-- 需要完成

在完成步骤1后，你会变成

- node
  - node.go <-- 需要完成
  - node.g0 <-- 备份（原始题目）
  - node2.g0 <-- 备份（原始题目） 
  - node2.go <-- 需要完成