# 阅读营CRM任务中心 — 技术方案 V1.0

> 对应产品文档：阅读营CRM任务中心V1.0
> 编写时间：2026-04-07
> 技术栈：Go / Gin / MySQL / Redis

---

## 一、PRD 要点梳理

### 1.1 任务类型差异（核心）

PRD 技术评审明确区分两种任务类型，数据结构有本质区别：

```
组合任务 (task_type=1)           单项任务 (task_type=2)
──────────────────────           ──────────────────────
Task                             Task
 ├── Module A                     ├── Node 1
 │    ├── Node 1                  ├── Node 2
 │    └── Node 2                  └── Node 3
 └── Module B
      └── Node 3
```

- **组合任务**：有模块（分组）层级，子任务数 = 所有模块子任务之和
- **单项任务**：无模块层级，子任务直挂任务下（`module_id = 0`）

### 1.2 任务状态（运行时计算，不存库）

| 状态 | 判断条件 |
|------|----------|
| 未开始 | `now < start_time` |
| 进行中 | `start_time ≤ now`，且 `end_time` 为零值或 `now ≤ end_time` |
| 已结束 | `end_time` 非零值 且 `now > end_time` |

> `state=0`（关闭开关）= APP 不展示，CRM 列表状态照常显示。

### 1.3 已确认的关键业务规则

- 任务创建时**不绑定课程**，课程关联由「课节管理」侧主动发起
- 任务列表**不展示**关联课程和参与人次
- 参与人群**多选**，存为 JSON 数组引用外部人群 ID
- 任务奖励**非必填**，引用外部权益系统 ID
- 列表默认按**创建时间倒序**，每页 **50 条**

---

## 二、数据模型设计

> 共 **4 张表**。命名规范参照 repository 包现有表：时间字段统一 `time.Time`（`ct`=创建时间、`ut`=更新时间），状态字段统一 `state`，JSON 配置字段统一 `conf`，无 `is_deleted`（通过 `state` 软删除）。

---

### 2.1 主任务表 `readcamp_task`

**DDL：**

```sql
CREATE TABLE `readcamp_task` (
  `id`           BIGINT       NOT NULL AUTO_INCREMENT         COMMENT '任务ID 自增主键',
  `name`         VARCHAR(100) NOT NULL DEFAULT ''             COMMENT '任务名称',
  `desc`         VARCHAR(256) NOT NULL DEFAULT ''             COMMENT '任务描述',
  `task_type`    INT          NOT NULL DEFAULT 1              COMMENT '任务类型 1=组合任务 2=单项任务',
  `start_time`   DATETIME     NOT NULL                        COMMENT '开始时间',
  `end_time`     DATETIME     NOT NULL                        COMMENT '结束时间',
  `state`        INT          NOT NULL DEFAULT 2              COMMENT '状态 1=开启 2=关闭',
  `user_conf`    TEXT         NOT NULL                        COMMENT '参与人群配置 1,2,3,...',
  `gift_id`      BIGINT       NOT NULL DEFAULT 0              COMMENT '奖励权益ID 0=无奖励',
  `ct`           DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut`           DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_ct`        (`ct`),
  KEY `idx_task_type` (`task_type`, `state`)
) ENGINE=InnoDB COMMENT='任务主表';
```

**状态筛选的 SQL 条件（列表搜索用）：**

```sql
-- 未开始: start_time > now()
-- 进行中: start_time <= now() AND (end_time = '0000-00-00 00:00:00' OR end_time >= now())
-- 已结束: end_time != '0000-00-00 00:00:00' AND end_time < now()
```

---

### 2.2 任务模块表 `readcamp_task_module`

> **仅组合任务使用**，单项任务无此层级。

**DDL：**

```sql
CREATE TABLE `readcamp_task_module` (
  `id`         BIGINT      NOT NULL AUTO_INCREMENT     COMMENT '模块ID',
  `task_id`    BIGINT      NOT NULL                    COMMENT '所属任务ID',
  `title`      VARCHAR(48) NOT NULL DEFAULT ''         COMMENT '模块标题 限10个汉字',
  `sort`       INT         NOT NULL DEFAULT 0          COMMENT '排序 升序',
  `ct`         DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `ut`         DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_task_id` (`task_id`)
) ENGINE=InnoDB COMMENT='任务模块表（组合任务专用）';
```

---
### 2.3 任务节点表 `readcamp_task_node`

> 组合任务和单项任务**共用**此表：组合任务通过 `module_id` 挂模块；单项任务 `module_id = 0` 直挂任务。

**DDL：**

```sql
CREATE TABLE `readcamp_task_node` (
  `id`        BIGINT       NOT NULL AUTO_INCREMENT   COMMENT '任务节点ID',
  `task_id`   BIGINT       NOT NULL                  COMMENT '所属任务ID（冗余）',
  `module_id` BIGINT       NOT NULL DEFAULT 0        COMMENT '所属模块ID 单项任务固定=0',
  `title`     VARCHAR(100) NOT NULL DEFAULT ''       COMMENT '任务节点标题',
  `desc`      VARCHAR(256) NOT NULL DEFAULT ''       COMMENT '任务节点描述',
  `node_type` INT          NOT NULL DEFAULT 1        COMMENT '任务节点类型 见枚举',
  `conf`      LONGTEXT     NOT NULL                  COMMENT '任务节点类型专属配置 JSON',
  `target`    INT          NOT NULL DEFAULT 1        COMMENT '目标完成量（次数/秒）看情况使用',
  `gift_id`   BIGINT       NOT NULL DEFAULT 0        COMMENT '奖励权益ID 0=无奖励',
  `sort`      INT          NOT NULL DEFAULT 0        COMMENT '排序 升序',
  `ct`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `ut`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_task_id`   (`task_id`),
  KEY `idx_module_id` (`module_id`)
) ENGINE=InnoDB COMMENT='任务节点表（组合/单项通用）';
```

**`conf` 各类型示例：**

| node_type | 类型 | conf 示例 |
|-----------|------|-----------|
| 1 | 看视频 | `{"video_id":101,"min_duration":60}` 单位秒 |
| 2 | 阅读   | `{"article_id":201,"min_read_seconds":120}` |
| 3 | 打卡   | `{"check_in_type":"daily","day_index":1}` |
| 4 | 完成作业 | `{"homework_type":"write","lesson_id":301}` |
| … | 可扩展 | 新增类型只加枚举+注册 Handler，不改表结构 |

---

### 2.4 用户任务进度表 `readcamp_user_task_progress`

```sql
CREATE TABLE `readcamp_user_task_progress` (
  `id`          BIGINT   NOT NULL AUTO_INCREMENT,
  `uid`         BIGINT   NOT NULL             COMMENT '学员UID',
  `task_id`     BIGINT   NOT NULL             COMMENT '任务ID',
  `state`       INT      NOT NULL DEFAULT 1   COMMENT '1=进行中 2=已完成',
  `node_done`   INT      NOT NULL DEFAULT 0   COMMENT '已完成子任务数',
  `finish_time` DATETIME                      COMMENT '任务完成时间',
  `ct`          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `ut`          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_uid_task`  (`uid`, `task_id`),
  KEY `idx_task_state` (`task_id`, `state`)
) ENGINE=InnoDB COMMENT='用户主任务进度';
```

---

### 2.5 用户子任务进度表 `readcamp_user_node_progress`

```sql
CREATE TABLE `readcamp_user_node_progress` (
  `id`          BIGINT   NOT NULL AUTO_INCREMENT,
  `uid`         BIGINT   NOT NULL             COMMENT '学员UID',
  `task_id`     BIGINT   NOT NULL             COMMENT '任务ID（冗余，按任务查全量进度用）',
  `node_id`     BIGINT   NOT NULL             COMMENT '任务节点ID',
  `cur_value`   INT      NOT NULL DEFAULT 0   COMMENT '当前完成量（次数/秒）',
  `state`       INT      NOT NULL DEFAULT 1   COMMENT '1=进行中 2=已完成',
  `finish_time` DATETIME                      COMMENT '完成时间',
  `ct`          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `ut`          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_uid_node`  (`uid`, `node_id`),
  KEY `idx_uid_task`  (`uid`, `task_id`)
) ENGINE=InnoDB COMMENT='用户子任务进度';
```

---

## 三、核心设计

### 4.1 单项 vs 组合任务的保存逻辑

```
保存任务（Service 层）：

task_type = 1（组合任务）：
  1. 写入 readcamp_task
  2. diff modules → 新增/更新 readcamp_task_module，删除不在列表中的
  3. 每个 module diff nodes → 新增/更新 readcamp_task_node(module_id=module.id)
  4. 统计 node_count = sum(module.node_count)，回写 readcamp_task

task_type = 2（单项任务）：
  1. 写入 readcamp_task
  2. diff nodes → 新增/更新 readcamp_task_node(module_id=0)
  3. node_count = len(nodes)，回写 readcamp_task

diff 策略：传入 id=0 新建；id>0 更新；前端不传的旧 id = 已删除 → DeleteById
（不全删重建，保留用户进度数据）
```

---

### 4.2 子任务策略模式（NodeHandler）

新增子任务类型 = 新增一个文件 + 注册，**核心调度逻辑零改动**。

```go
// NodeHandler 每种子任务类型实现此接口
type NodeHandler interface {
    Type() int
    ValidateConf(conf string) error   // 保存时校验 conf JSON
    CheckCompletion(ctx context.Context, node *ReadcampTaskNode, uid int64) (done bool, curVal int, err error)
}

// NodeRegistry 策略注册表
type NodeRegistry struct {
    m map[int]NodeHandler
}

func NewNodeRegistry(handlers []NodeHandler) *NodeRegistry {
    r := &NodeRegistry{m: make(map[int]NodeHandler)}
    for _, h := range handlers {
        r.m[h.Type()] = h
    }
    return r
}
```

**示例：看视频 Handler**

```go
type WatchVideoConf struct {
    VideoId     int64 `json:"video_id"`
    MinDuration int   `json:"min_duration"` // 秒
}

type WatchVideoHandler struct{ videoRepo VideoRepo }

func (h *WatchVideoHandler) Type() int { return NodeTypeWatchVideo }

func (h *WatchVideoHandler) ValidateConf(conf string) error {
    var c WatchVideoConf
    if err := json.Unmarshal([]byte(conf), &c); err != nil {
        return err
    }
    if c.VideoId <= 0 {
        return errors.New("video_id 不能为空")
    }
    return nil
}

func (h *WatchVideoHandler) CheckCompletion(ctx context.Context, node *ReadcampTaskNode, uid int64) (bool, int, error) {
    var c WatchVideoConf
    _ = json.Unmarshal([]byte(node.Conf), &c)
    watched, err := h.videoRepo.GetWatchedSeconds(ctx, uid, c.VideoId)
    if err != nil {
        return false, 0, err
    }
    return watched >= c.MinDuration, watched, nil
}
```

---

### 4.3 进度更新链路（事件驱动，与核心业务解耦）

```
学员行为（看完视频/完成阅读）
        │
        ▼
  业务系统发 Event（MQ）
        │
        ▼
  ProgressConsumer.Handle(event)
        │
        ├─ 查询用户关联的进行中任务
        ├─ 找到对应 node_id
        ├─ NodeRegistry.Get(node_type).CheckCompletion()
        ├─ 更新 readcamp_user_node_progress
        ├─ 判断主任务是否全部完成（node_done == node_total）
        └─ 完成 → 更新 readcamp_user_task_progress.state=2 + 发放奖励
```

---

## 四、API 设计

> 统一返回：`{ "ret": 0, "msg": "ok", "data": {} }`

### 4.1 CRM 后台接口

#### 任务列表
```
POST /admin/readcamp/task/list
```
```json
// Request
{
  "keyword":           "关键词（任务名/ID）",
  "task_type":         0,
  "task_status":       0,
  "start_time_begin":  "2026-01-01 00:00:00",
  "start_time_end":    "2026-12-31 23:59:59",
  "page":              1,
  "size":              50
}

// Response data
{
  "total": 100,
  "list": [
    {
      "id":          1,
      "name":        "新手任务",
      "task_type":   1,
      "task_status": 2,
      "node_count":  5,
      "start_time":  "2026-03-01 00:00:00",
      "end_time":    "",
      "state":       1,
      "ct":          "2026-03-01 10:00:00"
    }
  ]
}
```

#### 保存任务（创建/编辑通用）
```
POST /admin/readcamp/task/save
```

**组合任务 Request：**
```json
{
  "id":           0,
  "name":         "新手成长之路",
  "task_type":    1,
  "start_time":   "2026-04-01 00:00:00",
  "end_time":     "",
  "audience_ids": [1, 2],
  "reward_id":    3,
  "modules": [
    {
      "id":    0,
      "title": "第一阶段",
      "sort":  1,
      "nodes": [
        {
          "id":        0,
          "node_type": 1,
          "title":     "观看入门视频",
          "conf":      { "video_id": 101, "min_duration": 60 },
          "target":    1,
          "jump_type": 1,
          "jump_url":  "/h5/video/101",
          "sort":      1
        }
      ]
    }
  ]
}
```

**单项任务 Request：**
```json
{
  "id":           0,
  "name":         "每日打卡",
  "task_type":    2,
  "start_time":   "2026-04-01 00:00:00",
  "end_time":     "2026-04-30 23:59:59",
  "audience_ids": [1],
  "reward_id":    0,
  "nodes": [
    {
      "id":        0,
      "node_type": 3,
      "title":     "完成今日打卡",
      "conf":      { "check_in_type": "daily", "day_index": 1 },
      "target":    1,
      "jump_type": 1,
      "jump_url":  "/h5/checkin",
      "sort":      1
    }
  ]
}
```

#### 其他后台接口

| Path | 功能 | 关键参数 |
|------|------|---------|
| `POST /admin/readcamp/task/detail` | 任务详情（含模块/子任务） | `{"id": 1}` |
| `POST /admin/readcamp/task/toggle` | 开关切换 | `{"id": 1, "state": 0}` |
| `POST /admin/readcamp/task/delete` | 逻辑删除（state=2） | `{"id": 1}` |
| `POST /admin/readcamp/task/preview` | 获取预览二维码 URL | `{"id": 1}` |
| `POST /admin/readcamp/reward/search` | 搜索奖励权益（联想下拉） | `{"keyword": ""}` |

### 4.2 学员端接口

| Path | 功能 |
|------|------|
| `POST /api/readcamp/task/list` | 当前用户可见任务列表（含进度） |
| `POST /api/readcamp/task/detail` | 任务详情（含子任务进度） |
| `POST /api/readcamp/task/node/report` | 上报子任务完成行为（触发进度检查） |

---

## 五、目录结构

```
readcamp/task/
├── handler/
│   ├── admin.go         # 后台 handler
│   └── app.go           # 学员端 handler
├── service/
│   ├── task.go          # 任务 CRUD，按 task_type 分支处理
│   └── progress.go      # 进度更新、完成判断、奖励发放
├── strategy/
│   ├── registry.go      # NodeRegistry
│   ├── watch_video.go
│   ├── read.go
│   └── check_in.go
└── types/
    ├── req.go
    └── resp.go
```

> Repo 层文件放在 `readcamp/service/repository/` 下，与现有文件同目录，命名规范保持一致：
> - `readcamp_task.go`
> - `readcamp_task_module.go`
> - `readcamp_task_node.go`
> - `readcamp_user_task_progress.go`
> - `readcamp_user_node_progress.go`

---

## 六、扩展性说明

| 扩展点 | 方式 |
|--------|------|
| 新增子任务类型 | 实现 `NodeHandler` 接口 + 注册，不改表 |
| 多业务线复用 | `readcamp_task` 加 `biz_type` 字段，按 biz_type 路由不同 Registry |
| 不同人群不同目标 | V1 建独立任务；V2 可在 `conf` 扩展 `target_by_audience` |

---