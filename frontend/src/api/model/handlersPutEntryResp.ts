/**
 * Generated by orval v7.1.1 🍺
 * Do not edit manually.
 * AIKnow API
 * ## Sidebar 多级列表

列表的每一项是一个`entry`，主要字段是`id`，`type`和`title`。

- `id`: 全局唯一标识
- `type`: 合法取值为`note`，`group`或`link`
- `title`: 标题

一个`entry`下可以嵌套子`entry`，形成树形结构。

### 查询流程

1. 调用`Get Workspace`拿到 Private 和 Team Space 分别对应的最外层`entry id`。
2. 调用`Get Sub-Entries`，传参`entry id`，拿到该`entry`直接嵌套的子`entry`列表。
3. 如果子`entry`继续嵌套子`entry`（`has_sub_entries`为`true`），递归调用`Get Sub-Entries`。

### 拖拽/移动流程

调用`Update Entry`，更新`entry`的`parent`和`position`。
 * OpenAPI spec version: 1.0
 */
import type { HandlersEntry } from './handlersEntry'

export interface HandlersPutEntryResp {
  entry?: HandlersEntry
}
