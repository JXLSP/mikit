# Mikit - 在线数据库弱口令和未授权检测程序

## 项目架构与依赖注入

本项目采用清晰的分层架构，通过依赖注入模式实现各层之间的松耦合。以下是依赖注入的实现方式：

### 依赖注入链路

```
App初始化 -> 创建Store实例 -> 注入到Repository -> 注入到UseCase -> 注入到Controller
```

### 各层职责

1. **Store层**：负责数据库连接和基础数据访问
   - 单例模式确保全局只有一个数据库连接实例
   - 提供事务支持和数据库操作接口

2. **Repository层**：负责数据持久化逻辑
   - 依赖Store层进行数据库操作
   - 实现领域接口，隔离数据访问细节

3. **UseCase层**：负责业务逻辑
   - 依赖Repository层接口，不直接依赖实现
   - 通过接口实现业务逻辑与数据访问的解耦

4. **Controller层**：负责处理HTTP请求
   - 依赖UseCase层接口，处理请求并返回响应
   - 不包含业务逻辑，只负责请求参数解析和响应封装

### 依赖注入的优势

1. **松耦合**：各层通过接口而非具体实现进行交互，降低耦合度
2. **可测试性**：便于单元测试，可以轻松模拟依赖组件
3. **可维护性**：代码结构清晰，职责分明，便于维护和扩展
4. **灵活性**：可以轻松替换各层实现，如切换数据库或修改业务逻辑

### 示例：任务管理功能的依赖注入

```go
// 在app.go中初始化
dbInstance, _ := initDatabase()
storeInstance := store.NewStore(dbInstance)

// 在Controller初始化中注入依赖
func NewTasksController(ds *store.Store) *TasksController {
    // 创建Repository实例，注入Store依赖
    repo := repository.NewTaskRepository(ds)
    // 创建UseCase实例，注入Repository依赖
    return &TasksController{
        TUsecase: usecase.NewTasksUsecase(repo),
    }
}
```
