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

> 共 **7 张表**。命名规范参照 repository 包现有表：时间字段统一 `time.Time`（`ct`=创建时间、`ut`=更新时间），状态字段统一 `state`，JSON 配置字段统一 `conf`，无 `is_deleted`（通过 `state` 软删除）。

---

### 2.1 主任务表 `readcamp_task`

**DDL：**

```sql
CREATE TABLE `readcamp_task` (
  `id`           BIGINT       NOT NULL AUTO_INCREMENT         COMMENT '任务ID 自增主键',
  `biz_type`     INT          NOT NULL DEFAULT 1              COMMENT '业务线 1=阅读营 2=小班课 ...',
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
  KEY `idx_biz_ct`    (`biz_type`, `ct`),
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
  `sort`       INT         NOT NULL DEFAULT 0          COMMENT '排序 从大到小',
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
  `sort`      INT          NOT NULL DEFAULT 0        COMMENT '排序 从大到小',
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

### 2.6 奖励发放领取记录表 `readcamp_task_gift_record`

```sql
CREATE TABLE `readcamp_task_gift_record` (
  `id`        BIGINT   NOT NULL AUTO_INCREMENT,
  `uid`       BIGINT   NOT NULL             COMMENT '学员UID',
  `task_id`   BIGINT   NOT NULL             COMMENT '任务ID',
  `node_id`   BIGINT   NOT NULL DEFAULT 0   COMMENT '节点ID 0=任务级奖励',
  `gift_id`   BIGINT   NOT NULL             COMMENT '奖励权益ID',
  `state`     INT      NOT NULL DEFAULT 1   COMMENT '1=待领取 2=已领取',
  `ct`        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `ut`        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_uid_task_node` (`uid`, `task_id`, `node_id`),
  KEY `idx_uid_state` (`uid`, `state`)
) ENGINE=InnoDB COMMENT='任务奖励发放领取记录';
```

**字段说明：**
- `node_id=0`：任务级奖励，由 `readcamp_task.gift_id` 触发
- `node_id>0`：节点级奖励，由 `readcamp_task_node.gift_id` 触发
- `uk_uid_task_node`：唯一键防重复发放，同一用户同一任务/节点奖励只发一次
- 状态流转：任务/节点完成 → 写入 `state=1`（已发放）→ 用户在 APP 领取 → 更新 `state=2`（已领取）

---

### 2.7 用户节点行为事件明细表 `readcamp_user_node_event_log`

> **用途一（核心）**：存储 `ReportNodeEventReq.event_id` 实现幂等去重——同一 `event_id` 重复上报时 `uk_event_id` 报唯一键冲突，直接 ignore，`cur_value` 不会被重复累加。
>
> **用途二**：保留每次行为的原始记录，支持进度回溯、异常排查、数据分析（如"某节点每日完成人数"）。

**DDL：**

```sql
CREATE TABLE `readcamp_user_node_event_log` (
  `id`        BIGINT      NOT NULL AUTO_INCREMENT,
  `event_id`  VARCHAR(64) NOT NULL DEFAULT '' COMMENT '幂等键，由上报方生成，全局唯一',
  `uid`       BIGINT      NOT NULL            COMMENT '学员UID',
  `task_id`   BIGINT      NOT NULL            COMMENT '任务ID',
  `node_id`   BIGINT      NOT NULL            COMMENT '节点ID',
  `node_type` INT         NOT NULL DEFAULT 0  COMMENT '节点类型，冗余方便按类型分析',
  `value`     BIGINT      NOT NULL DEFAULT 0  COMMENT '本次上报的事件量（次数/秒）',
  `ct`        DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '事件发生时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_event_id`  (`event_id`),
  KEY        `idx_uid_node` (`uid`, `node_id`, `ct`)
) ENGINE=InnoDB COMMENT='用户节点行为事件明细（幂等日志）';
```

**与进度表的关系：**

```
ReportNodeEvent 调用链路（加入幂等）：

1. INSERT IGNORE INTO readcamp_user_node_event_log (event_id, ...)
   → 影响行数 = 0：说明重复上报，直接返回（幂等）
   → 影响行数 = 1：首次上报，继续

2. 更新 readcamp_user_node_progress.cur_value += value
3. 若节点完成 → 写奖励记录、累计 node_done、检查任务完成
```

---

> 严格对应 DDL 字段，标签格式 `json:"xxx" bdb:"xxx"`，参照 repository 包现有规范。

### 3.1 ReadcampTask

```go
const (
    tableReadcampTask = "readcamp_task"

    BizTypeReadCamp = 1 // 阅读营
    BizTypeClassRoom = 2 // 小班课（示例，按需扩展）

    TaskTypeCombo  = 1 // 组合任务
    TaskTypeSingle = 2 // 单项任务

    TaskStateOpen  = 1 // 开启（APP 可见）
    TaskStateClose = 2 // 关闭（APP 不可见）
)

type ReadcampTask struct {
    Id        int64     `json:"id"         bdb:"id"`         // 任务ID
    BizType   int       `json:"biz_type"   bdb:"biz_type"`   // 业务线
    Name      string    `json:"name"       bdb:"name"`       // 任务名称
    Desc      string    `json:"desc"       bdb:"desc"`       // 任务描述
    TaskType  int       `json:"task_type"  bdb:"task_type"`  // 1=组合任务 2=单项任务
    StartTime time.Time `json:"start_time" bdb:"start_time"` // 开始时间
    EndTime   time.Time `json:"end_time"   bdb:"end_time"`   // 结束时间
    State     int       `json:"state"      bdb:"state"`      // 1=开启 2=关闭
    UserConf  string    `json:"user_conf"  bdb:"user_conf"`  // 参与人群ID "1,2,3"
    GiftId    int64     `json:"gift_id"    bdb:"gift_id"`    // 奖励权益ID 0=无
    Ct        time.Time `json:"ct"         bdb:"ct"`
    Ut        time.Time `json:"ut"         bdb:"ut"`
}

// TaskStatus 运行时计算，不存库
type TaskStatus int

const (
    TaskStatusNotStarted TaskStatus = 1 // 未开始
    TaskStatusInProgress TaskStatus = 2 // 进行中
    TaskStatusEnded      TaskStatus = 3 // 已结束
)

func (t *ReadcampTask) ComputeStatus() TaskStatus {
    now := time.Now()
    if now.Before(t.StartTime) {
        return TaskStatusNotStarted
    }
    if t.EndTime.IsZero() || !now.After(t.EndTime) {
        return TaskStatusInProgress
    }
    return TaskStatusEnded
}
```

### 3.2 ReadcampTaskModule

```go
const tableReadcampTaskModule = "readcamp_task_module"

type ReadcampTaskModule struct {
    Id     int64     `json:"id"      bdb:"id"`
    TaskId int64     `json:"task_id" bdb:"task_id"` // 所属任务ID
    Title  string    `json:"title"   bdb:"title"`   // 模块标题 限10汉字
    Sort   int       `json:"sort"    bdb:"sort"`
    Ct     time.Time `json:"ct"      bdb:"ct"`
    Ut     time.Time `json:"ut"      bdb:"ut"`
}
```

### 3.3 ReadcampTaskNode

```go
const (
    tableReadcampTaskNode = "readcamp_task_node"

    NodeTypeWatchVideo = 1 // 看视频
    NodeTypeRead       = 2 // 阅读
    NodeTypeCheckIn    = 3 // 打卡
    NodeTypeHomework   = 4 // 完成作业
)

type ReadcampTaskNode struct {
    Id       int64     `json:"id"        bdb:"id"`
    TaskId   int64     `json:"task_id"   bdb:"task_id"`   // 所属任务ID（冗余）
    ModuleId int64     `json:"module_id" bdb:"module_id"` // 所属模块ID，单项任务=0
    Title    string    `json:"title"     bdb:"title"`     // 节点标题
    Desc     string    `json:"desc"      bdb:"desc"`      // 节点描述
    NodeType int       `json:"node_type" bdb:"node_type"` // 节点类型
    Conf     string    `json:"conf"      bdb:"conf"`      // 类型专属配置 JSON string
    Target   int       `json:"target"    bdb:"target"`    // 目标完成量（次数/秒）
    GiftId   int64     `json:"gift_id"   bdb:"gift_id"`   // 节点奖励ID 0=无
    Sort     int       `json:"sort"      bdb:"sort"`
    Ct       time.Time `json:"ct"        bdb:"ct"`
    Ut       time.Time `json:"ut"        bdb:"ut"`
}
```

### 3.4 ReadcampUserTaskProgress

```go
const (
    tableReadcampUserTaskProgress = "readcamp_user_task_progress"

    UserTaskStateInProgress = 1 // 进行中
    UserTaskStateFinished   = 2 // 已完成
)

type ReadcampUserTaskProgress struct {
    Id         int64      `json:"id"          bdb:"id"`
    Uid        int64      `json:"uid"         bdb:"uid"`
    TaskId     int64      `json:"task_id"     bdb:"task_id"`
    State      int        `json:"state"       bdb:"state"`       // 1=进行中 2=已完成
    NodeDone   int        `json:"node_done"   bdb:"node_done"`   // 已完成节点数
    FinishTime *time.Time `json:"finish_time" bdb:"finish_time"` // 完成时间
    Ct         time.Time  `json:"ct"          bdb:"ct"`
    Ut         time.Time  `json:"ut"          bdb:"ut"`
}
```

### 3.5 ReadcampUserNodeProgress

```go
const (
    tableReadcampUserNodeProgress = "readcamp_user_node_progress"

    UserNodeStateInProgress = 1 // 进行中
    UserNodeStateFinished   = 2 // 已完成
)

type ReadcampUserNodeProgress struct {
    Id         int64      `json:"id"          bdb:"id"`
    Uid        int64      `json:"uid"         bdb:"uid"`
    TaskId     int64      `json:"task_id"     bdb:"task_id"`   // 冗余
    NodeId     int64      `json:"node_id"     bdb:"node_id"`
    CurValue   int        `json:"cur_value"   bdb:"cur_value"` // 当前完成量
    State      int        `json:"state"       bdb:"state"`     // 1=进行中 2=已完成
    FinishTime *time.Time `json:"finish_time" bdb:"finish_time"`
    Ct         time.Time  `json:"ct"          bdb:"ct"`
    Ut         time.Time  `json:"ut"          bdb:"ut"`
}
```

### 3.6 ReadcampTaskGiftRecord

```go
const (
    tableReadcampTaskGiftRecord = "readcamp_task_gift_record"

    GiftStateIssued  = 1 // 已发放（待用户领取）
    GiftStateClaimed = 2 // 已领取
)

type ReadcampTaskGiftRecord struct {
    Id       int64     `json:"id"        bdb:"id"`
    Uid      int64     `json:"uid"       bdb:"uid"`
    TaskId   int64     `json:"task_id"   bdb:"task_id"`
    NodeId   int64     `json:"node_id"   bdb:"node_id"`   // 0=任务级奖励
    GiftId   int64     `json:"gift_id"   bdb:"gift_id"`
    State    int       `json:"state"     bdb:"state"`     // 1=已发放 2=已领取
    Ct       time.Time `json:"ct"        bdb:"ct"`
    Ut       time.Time `json:"ut"        bdb:"ut"`
}
```

### 3.7 ReadcampUserNodeEventLog

```go
const tableReadcampUserNodeEventLog = "readcamp_user_node_event_log"

type ReadcampUserNodeEventLog struct {
    Id       int64     `json:"id"        bdb:"id"`
    EventId  string    `json:"event_id"  bdb:"event_id"`  // 幂等键
    Uid      int64     `json:"uid"       bdb:"uid"`
    TaskId   int64     `json:"task_id"   bdb:"task_id"`
    NodeId   int64     `json:"node_id"   bdb:"node_id"`
    NodeType int64     `json:"node_type" bdb:"node_type"` // 冗余
    Value    int64     `json:"value"     bdb:"value"`     // 本次事件量
    Ct       time.Time `json:"ct"        bdb:"ct"`
}
```

---

## 四、核心逻辑设计

### 4.1 单项 vs 组合任务的保存逻辑

```
保存任务（Service 层）：

task_type = 1（组合任务）：
  1. 写入/更新 readcamp_task
  2. diff modules：
       id=0 → InsertOne readcamp_task_module
       id>0 → UpdateById readcamp_task_module
       前端未传的旧 module.id → DeleteById（级联删其下所有 node）
  3. 每个 module diff nodes：
       id=0 → InsertOne readcamp_task_node (module_id=module.id)
       id>0 → UpdateById readcamp_task_node
       前端未传的旧 node.id → DeleteById

task_type = 2（单项任务）：
  1. 写入/更新 readcamp_task
  2. diff nodes（module_id 固定=0）：
       id=0 → InsertOne readcamp_task_node
       id>0 → UpdateById readcamp_task_node
       前端未传的旧 node.id → DeleteById

diff 原则：不全删重建，保留用户进度数据；
已有进度的 node 被删除时，对应进度记录保留（历史数据），不影响完成判断（按当前存活节点 COUNT）。
```

**查询任务详情时的分支处理：**

```go
func (s *TaskService) Detail(ctx context.Context, taskID int64) (*TaskDetailResp, error) {
    task, _ := s.taskRepo.GetOne(ctx, QueryReadcampTask{Id: taskID})
    resp := buildBaseResp(task) // 填充 name/desc/state/gift_id/user_conf 等

    switch task.TaskType {
    case TaskTypeCombo:
        modules, _ := s.moduleRepo.BatchGet(ctx, QueryReadcampTaskModule{TaskId: taskID})
        allNodes, _ := s.nodeRepo.BatchGet(ctx, QueryReadcampTaskNode{TaskId: taskID})
        nodesByModule := groupByModuleId(allNodes)
        for _, m := range modules {
            resp.Modules = append(resp.Modules, buildModuleResp(m, nodesByModule[m.Id]))
        }
    case TaskTypeSingle:
        // module_id=0 的节点直接挂任务
        nodes, _ := s.nodeRepo.BatchGet(ctx, QueryReadcampTaskNode{TaskId: taskID, ModuleId: 0})
        resp.Nodes = buildNodeRespList(nodes)
    }
    return resp, nil
}
```

---

### 4.2 任务完成判断

`readcamp_user_task_progress` 不存 `node_total`，完成判断时实时查询当前存活节点数：

```go
func (s *ProgressService) checkTaskDone(ctx context.Context, p *ReadcampUserTaskProgress) (bool, error) {
    // 查该任务当前存活的节点总数
    total, err := s.nodeRepo.Count(ctx, QueryReadcampTaskNode{TaskId: p.TaskId})
    if err != nil {
        return false, err
    }
    return int64(p.NodeDone) >= total, nil
}
```

---

### 4.3 子任务策略模式（NodeHandler）& 多业务线路由

新增子任务类型 = 新增一个文件 + 注册，**核心调度逻辑零改动**。
接入新业务线 = 为该业务线构建独立 `NodeRegistry`，挂载到 `TaskService`，**不改已有业务逻辑**。

```go
// NodeHandler 每种节点类型实现此接口
type NodeHandler interface {
    Type() int
    // ValidateConf 保存时校验 conf JSON 合法性
    ValidateConf(conf string) error
    // CheckCompletion 判断用户是否满足完成条件，返回是否完成及当前完成量
    CheckCompletion(ctx context.Context, node *ReadcampTaskNode, uid int64) (done bool, curVal int, err error)
}

// NodeRegistry 单业务线的节点策略注册表
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

func (r *NodeRegistry) Get(nodeType int) (NodeHandler, bool) {
    h, ok := r.m[nodeType]
    return h, ok
}
```

**TaskService 持有多业务线 Registry Map，按 `biz_type` 路由：**

```go
type TaskService struct {
    registries map[int]*NodeRegistry // key: biz_type
    taskRepo   ReadcampTaskRepo
    moduleRepo ReadcampTaskModuleRepo
    nodeRepo   ReadcampTaskNodeRepo
}

// getRegistry 按 biz_type 取对应注册表，不存在则返回 error
func (s *TaskService) getRegistry(bizType int) (*NodeRegistry, error) {
    r, ok := s.registries[bizType]
    if !ok {
        return nil, fmt.Errorf("unsupported biz_type: %d", bizType)
    }
    return r, nil
}
```

**初始化时为每个业务线构建独立 Registry：**

```go
func NewTaskService(taskRepo ReadcampTaskRepo, ...) *TaskService {
    return &TaskService{
        registries: map[int]*NodeRegistry{
            BizTypeReadCamp: NewNodeRegistry([]NodeHandler{
                &WatchVideoHandler{videoRepo: videoRepo},
                &ReadHandler{readRepo: readRepo},
                &CheckInHandler{checkInRepo: checkInRepo},
            }),
            // 接入新业务线：在此追加，完全不动已有逻辑
            // BizTypeClassRoom: NewNodeRegistry([]NodeHandler{
            //     &MathExerciseHandler{exerciseRepo: exerciseRepo},
            //     &CheckInHandler{checkInRepo: checkInRepo}, // 公共类型可复用
            // }),
        },
        taskRepo:   taskRepo,
        moduleRepo: moduleRepo,
        nodeRepo:   nodeRepo,
    }
}
```

**保存任务时按 `biz_type` 校验节点类型合法性：**

```go
func (s *TaskService) Save(ctx context.Context, req *SaveTaskReq) error {
    registry, err := s.getRegistry(req.BizType)
    if err != nil {
        return err
    }
    for _, node := range req.AllNodes() { // 组合/单项通用
        handler, ok := registry.Get(node.NodeType)
        if !ok {
            return fmt.Errorf("biz_type=%d 不支持节点类型 node_type=%d", req.BizType, node.NodeType)
        }
        if err := handler.ValidateConf(node.Conf); err != nil {
            return err
        }
    }
    // ... 写库逻辑
}
```

**示例：看视频 Handler（仅阅读营使用）**

```go
type WatchVideoConf struct {
    VideoId     int64 `json:"video_id"`
    MinDuration int   `json:"min_duration"` // 单位秒
}

type WatchVideoHandler struct {
    videoRepo VideoRepo
}

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

### 4.4 进度更新链路（事件驱动）

```
学员行为（看完视频/完成阅读/打卡）
        │
        ▼
  业务系统发 Event（MQ），携带唯一 event_id
        │
        ▼
  ProgressConsumer.Handle(event)
        │
        ├─ INSERT IGNORE INTO readcamp_user_node_event_log (event_id, uid, task_id, node_id, ...)
        │    ├─ 影响行数 = 0 → 重复事件，直接 return（幂等）
        │    └─ 影响行数 = 1 → 首次，继续处理
        │
        ├─ 查询用户进行中的任务列表（readcamp_user_task_progress，state=1）
        ├─ 匹配对应 node_id（readcamp_task_node，by task_id + node_type）
        ├─ NodeRegistry.Get(node_type).CheckCompletion()
        ├─ 更新 readcamp_user_node_progress（cur_value / state / finish_time）
        │
        ├─ 若节点完成 且 node.gift_id > 0：
        │    └─ InsertOne readcamp_task_gift_record（state=1）
        │         ↑ uk_uid_task_node 防重，ignore duplicate key error
        │
        ├─ node_done++ → UpdateById readcamp_user_task_progress
        ├─ checkTaskDone() → COUNT 存活节点 vs node_done
        │
        └─ 任务完成 且 task.gift_id > 0：
             └─ InsertOne readcamp_task_gift_record（node_id=0, state=1）
                  + 更新 readcamp_user_task_progress.state=2 + finish_time
```

**学员端领取奖励流程（APP 侧）：**

```
用户点击"领取奖励"
        │
        ▼
  查询 readcamp_task_gift_record（uid + task_id，state=1）
        │
        ▼
  调用外部权益系统发放 gift_id
        │
        ▼
  UpdateById state=2（已领取）
```

---

## 五、API 设计

> 统一返回：`{"ret": 0, "msg": "ok", "data": {}}`
> 时间字段格式：`"2006-01-02 15:04:05"`，`end_time` 传空字符串表示永久有效。

### 5.1 CRM 后台接口

#### 任务列表
```
POST /admin/readcamp/task/list
```
```json
// Request
{
  "keyword":          "关键词（任务名/ID 模糊）",
  "biz_type":         1,
  "task_type":        0,
  "task_status":      0,
  "start_time_begin": "2026-01-01 00:00:00",
  "start_time_end":   "2026-12-31 23:59:59",
  "limit":            50,
  "offset":           0
}

// Response data
{
  "total":  100,
  "offset": 50,
  "more":   true,
  "list": [
    {
      "id":          1,
      "biz_type":    1,
      "name":        "新手任务",
      "desc":        "完成后可领取体验课礼包",
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

> `node_count` 为查询时 COUNT `readcamp_task_node` 所得，非存储字段。

---

#### 保存任务（创建/编辑通用）
```
POST /admin/readcamp/task/save
```

**组合任务 Request：**
```json
{
  "id":           0,
  "biz_type":     1,
  "name":         "新手成长之路",
  "desc":         "完成全部任务解锁专属奖励",
  "task_type":    1,
  "start_time":   "2026-04-01 00:00:00",
  "end_time":     "",
  "audience_ids": ["1", "2"],
  "gift_id":      3,
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
          "desc":      "观看完整视频即可完成",
          "conf":      { "video_id": 101, "min_duration": 60 },
          "target":    1,
          "gift_id":   0,
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
  "biz_type":     1,
  "name":         "每日打卡",
  "desc":         "坚持打卡赢奖励",
  "task_type":    2,
  "start_time":   "2026-04-01 00:00:00",
  "end_time":     "2026-04-30 23:59:59",
  "audience_ids": ["1", "2"],
  "gift_id":      0,
  "nodes": [
    {
      "id":        0,
      "node_type": 3,
      "title":     "完成今日打卡",
      "desc":      "每天打卡一次",
      "conf":      { "check_in_type": "daily", "day_index": 1 },
      "target":    1,
      "gift_id":   0,
      "sort":      1
    }
  ]
}
```

> `audience_ids` 在 Service 层序列化为 `user_conf`（逗号分隔字符串）后写库。

---

#### 其他后台接口

| Path | 功能 | 关键参数 |
|------|------|---------|
| `POST /admin/readcamp/task/detail` | 任务详情（含模块/节点） | `{"id": 1}` |
| `POST /admin/readcamp/task/toggle` | 开关切换（开启/关闭） | `{"id": 1, "state": 2}` |
| `POST /admin/readcamp/task/delete` | 删除任务 | `{"id": 1}` |

---

### 5.2 RPC 接口（基础服务对外暴露）

> 任务中心作为基础服务，通过 RPC 对上层服务（CRM 后台、APP 服务、行为上报服务、班主任服务）统一暴露能力。上层 HTTP 接口调用 RPC，任务服务本身不直接依赖外部权益系统。

```protobuf
service TaskService {

  // ─── 任务管理（CRM 后台调用）───────────────────────────────────

  // 保存任务：id=0 为创建，id>0 为编辑，组合/单项通用
  rpc SaveTask(SaveTaskReq) returns (SaveTaskResp);

  // 任务列表：分页 + 多条件筛选，task_status 由 start_time/end_time 计算
  rpc TaskList(TaskListReq) returns (TaskListResp);

  // 任务详情（管理端）：含全部模块/节点结构，不含用户进度
  rpc TaskDetail(TaskDetailReq) returns (TaskDetailResp);

  // 开关切换：state=1 开启 / state=2 关闭
  rpc ToggleTask(ToggleTaskReq) returns (CommonResp);

  // 删除任务（物理删除，级联删除模块和节点）
  rpc DeleteTask(DeleteTaskReq) returns (CommonResp);

  // ─── 学员端（APP 服务调用）───────────────────────────────────

  // 学员任务列表：返回用户可见的任务及进度摘要
  // - query_type=1：进行中的任务（readcamp_user_task_progress.state=1）
  // - query_type=2：已参与的全部任务（有进度记录的）
  // - query_type=3：可参与的任务（人群命中 + 任务开启 + 时间进行中，含未开始进度的）
  rpc GetUserTaskList(GetUserTaskListReq) returns (GetUserTaskListResp);

  // 学员任务详情：指定 task_id，返回任务结构 + 该用户各节点完整进度
  // APP 任务详情页的核心接口
  rpc GetUserTaskDetail(GetUserTaskDetailReq) returns (GetUserTaskDetailResp);

  // 领取奖励：更新 state=2，调用基础服务里面的方法发放奖励
  rpc ClaimGift(ClaimGiftReq) returns (ClaimGiftResp);

  // ─── 进度上报（行为系统 / MQ 消费者调用）────────────────────────

  // 上报节点行为事件：任务服务内部路由 NodeHandler → 更新进度 → 写奖励记录
  // 幂等：重复上报安全，唯一键防重
  rpc ReportNodeEvent(ReportNodeEventReq) returns (ReportNodeEventResp);
}
```

## 六、扩展性说明

| 扩展点 | 方案 |
|--------|------|
| **新增节点类型** | 实现 `NodeHandler` 接口，注册到对应业务线的 `NodeRegistry`；`conf` 字段存 JSON，不改表结构 |
| **接入新业务线** | `readcamp_task.biz_type` 区分数据；`NewTaskService` 中为新 `biz_type` 构建独立 `NodeRegistry`；已有业务线逻辑不受影响，详见 §4.3 |
| **公共节点类型跨业务线复用** | 同一个 `NodeHandler` 实例可注册到多个 `NodeRegistry`，如 `CheckInHandler` 阅读营和小班课均可使用 |
| **不同人群不同目标** | V1：为不同人群建独立任务；V2：`conf` 扩展 `"target_by_audience":{"1":1,"2":2}` |

---

## 八、一期开发范围

| 优先级 | 功能 | 涉及表 |
|--------|------|--------|
| P0 | 任务 CRUD（组合+单项） | `readcamp_task` + `readcamp_task_module` + `readcamp_task_node` |
| P0 | 任务列表搜索 / 状态计算 / 开关 | `readcamp_task` |
| P0 | 学员端任务详情（含节点进度） | `readcamp_user_task_progress` + `readcamp_user_node_progress` |
| P1 | 节点完成上报（看视频/打卡）+ 幂等去重 | `readcamp_user_node_event_log` + `strategy/*` + `progress.go` |
| P1 | 奖励发放（任务级 + 节点级） | `readcamp_task_gift_record` |
| P1 | 学员端奖励领取 | `readcamp_task_gift_record` |
| P2 | 预览二维码生成 | — |
| P2 | 班主任进度查看 | `readcamp_user_task_progress` |