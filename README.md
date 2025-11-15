# Formula-Go

一个用 Go 语言实现的通达信公式解析器和执行引擎，为开发者和量化交易者提供解析和执行通达信技术指标公式的能力。

## 项目状态

✅ **Phase 1-3 完成** - 核心功能已全部实现并测试通过

本项目参考 [formula-ts](https://github.com/DTrader-store/formula-ts) TypeScript 实现，使用 Go 语言重新实现。

## 特性

- ✅ **类型安全**: 使用 Go 的强类型系统，确保代码安全性
- ✅ **高性能**: Go 语言的高性能特性，适合大规模数据处理
- ✅ **完整实现**: 词法分析、语法分析、解释执行全流程
- ✅ **丰富的内置函数**: 23 个内置函数，覆盖常用技术指标
- ✅ **易于集成**: 简洁的 API 设计，易于集成到现有项目
- ✅ **测试完善**: 单元测试和集成测试覆盖率超过 80%

## 安装

```bash
go get github.com/DTrader-store/formula-go
```

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/DTrader-store/formula-go"
)

func main() {
    // 创建市场数据
    data := []*formula.MarketData{
        formula.NewMarketData(100, 105, 107, 99, 1000, 100000),
        formula.NewMarketData(105, 103, 108, 102, 1100, 110000),
        formula.NewMarketData(103, 107, 109, 101, 1200, 120000),
        formula.NewMarketData(107, 110, 112, 106, 1300, 130000),
        formula.NewMarketData(110, 108, 113, 107, 1400, 140000),
    }

    // 创建公式引擎
    engine := formula.NewFormulaEngine()

    // 执行公式
    result, err := engine.Run("MA5 := MA(CLOSE, 5)", data)
    if err != nil {
        fmt.Printf("错误: %v\n", err)
        return
    }

    // 输出结果
    for _, output := range result.Outputs {
        fmt.Printf("%s: %v\n", output.Name, output.Data)
    }
}
```

## 支持的功能

### 1. 语法特性

- ✅ 变量声明: `MA5 := MA(CLOSE, 5)`
- ✅ 算术运算: `+`, `-`, `*`, `/`
- ✅ 比较运算: `>`, `<`, `>=`, `<=`, `=`, `<>`
- ✅ 逻辑运算: `AND`, `OR`
- ✅ 函数调用: `MA(CLOSE, 5)`
- ✅ 括号表达式: `(a + b) * c`
- ✅ 一元运算: `-x`

### 2. 内置函数

现已支持 **23 个内置函数**！

**数学统计函数**
- `MA(data, period)` - 简单移动平均
- `SMA(data, period)` - 简单移动平均（MA 的别名）
- `EMA(data, period)` - 指数移动平均
- `WMA(data, period)` - 加权移动平均
- `SUM(data, period)` - 求和
- `STD(data, period)` - 标准差
- `VAR(data, period)` - 方差
- `AVEDEV(data, period)` - 平均绝对偏差
- `MAX(a, b)` - 最大值
- `MIN(a, b)` - 最小值
- `ABS(value)` - 绝对值
- `SQRT(value)` - 平方根

**引用函数**
- `REF(data, n)` - 引用 n 期前的数据
- `HHV(data, period)` - 周期内最高值
- `LLV(data, period)` - 周期内最低值

**条件和逻辑函数**
- `IF(condition, trueValue, falseValue)` - 条件判断
- `COUNT(condition, period)` - 统计满足条件的周期数
- `EVERY(condition, period)` - 检查是否所有周期都满足条件
- `EXIST(condition, period)` - 检查是否存在满足条件的周期
- `BETWEEN(value, lower, upper)` - 检查值是否在范围内

**技术分析函数**
- `CROSS(a, b)` - 交叉检测（a 上穿 b）
- `BARSLAST(condition)` - 距离最后一次满足条件的周期数
- `FILTER(condition, period)` - 过滤信号，防止频繁触发

### 3. 内置变量

- `OPEN` - 开盘价
- `CLOSE` - 收盘价
- `HIGH` - 最高价
- `LOW` - 最低价
- `VOLUME` - 成交量
- `AMOUNT` - 成交额

## 使用示例

### 简单移动平均

```go
formula := "MA5 := MA(CLOSE, 5)"
result, _ := engine.Run(formula, marketData)
```

### MACD 指标

```go
formula := `
    EMA12 := EMA(CLOSE, 12)
    EMA26 := EMA(CLOSE, 26)
    DIF := EMA12 - EMA26
    DEA := EMA(DIF, 9)
    MACD := (DIF - DEA) * 2
`
result, _ := engine.Run(formula, marketData)
```

### 金叉检测

```go
formula := `
    MA5 := MA(CLOSE, 5)
    MA10 := MA(CLOSE, 10)
    SIGNAL := CROSS(MA5, MA10)
`
result, _ := engine.Run(formula, marketData)
```

### 条件选股

```go
formula := `
    MA5 := MA(CLOSE, 5)
    MA10 := MA(CLOSE, 10)
    GOLDEN := CROSS(MA5, MA10)
    STRONG := CLOSE > MA5 AND EVERY(CLOSE > OPEN, 3)
    SELECT := GOLDEN AND STRONG
`
result, _ := engine.Run(formula, marketData)
// SELECT 中为 1 的位置表示满足选股条件
```

### 布林带指标

```go
formula := `
    MID := MA(CLOSE, 20)
    STDEV := STD(CLOSE, 20)
    UPPER := MID + 2 * STDEV
    LOWER := MID - 2 * STDEV
    BREAK_UP := CROSS(CLOSE, UPPER)
    BREAK_DOWN := CROSS(LOWER, CLOSE)
`
result, _ := engine.Run(formula, marketData)
```

### KDJ 指标

```go
formula := `
    LOW9 := LLV(LOW, 9)
    HIGH9 := HHV(HIGH, 9)
    RSV := (CLOSE - LOW9) / (HIGH9 - LOW9) * 100
    K := SMA(RSV, 3)
    D := SMA(K, 3)
    J := 3 * K - 2 * D
`
result, _ := engine.Run(formula, marketData)
```

### 信号过滤

```go
formula := `
    MA5 := MA(CLOSE, 5)
    MA10 := MA(CLOSE, 10)
    GOLDEN := CROSS(MA5, MA10)
    FILTERED := FILTER(GOLDEN, 10)
`
result, _ := engine.Run(formula, marketData)
// FILTERED 会过滤掉 10 个周期内的重复信号
```

## 项目结构

```
formula-go/
├── engine/              # 公式引擎
│   ├── engine.go       # FormulaEngine 主类
│   └── engine_test.go  # 集成测试
├── errors/              # 错误类型定义
│   ├── errors.go       # 各类错误
│   └── errors_test.go
├── interpreter/         # 解释器
│   ├── interpreter.go  # 解释执行
│   ├── functions.go    # 内置函数
│   └── registry.go     # 函数注册
├── lexer/              # 词法分析器
│   ├── lexer.go        # 词法分析主逻辑
│   ├── token.go        # Token 定义
│   ├── token_type.go   # Token 类型
│   └── *_test.go
├── parser/             # 语法分析器
│   ├── parser.go       # 语法分析主逻辑
│   ├── parser_test.go
│   └── ast/            # 抽象语法树
│       └── nodes.go
├── types/              # 类型定义
│   ├── market_data.go  # 市场数据
│   ├── formula_result.go # 公式结果
│   └── *_test.go
├── formula.go          # 主入口，导出 API
├── go.mod
└── README.md
```

## 测试

```bash
# 运行所有测试
go test ./...

# 运行测试并显示覆盖率
go test ./... -cover

# 运行详细测试
go test ./... -v

# 生成覆盖率报告
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

**当前测试覆盖率**:
- engine: 81.2%
- errors: 100%
- lexer: 90.5%
- parser: 77.8%
- types: 100%
- **总体**: >80%

## API 文档

### FormulaEngine

```go
type FormulaEngine struct {}

// 创建新引擎
func NewFormulaEngine() *FormulaEngine

// 编译公式为 AST
func (e *FormulaEngine) Compile(formula string) (*Program, error)

// 执行已编译的程序
func (e *FormulaEngine) Execute(program *Program, marketData []*MarketData) (*FormulaResult, error)

// 一步编译并执行
func (e *FormulaEngine) Run(formula string, marketData []*MarketData) (*FormulaResult, error)
```

### MarketData

```go
type MarketData struct {
    Open   float64
    Close  float64
    High   float64
    Low    float64
    Volume float64
    Amount float64
}

func NewMarketData(open, close, high, low, volume, amount float64) *MarketData
func (m *MarketData) Validate() error
```

### FormulaResult

```go
type FormulaResult struct {
    Outputs   []*OutputLine
    Variables map[string]float64
}

type OutputLine struct {
    Name  string
    Data  []float64
    Style *LineStyle
}
```

## 开发路线图

### ✅ Phase 1: 基础类型系统

- [x] 错误处理系统
- [x] Token 系统
- [x] AST 节点定义
- [x] 市场数据类型
- [x] 公式结果类型

### ✅ Phase 2: 词法分析器和语法分析器

- [x] 实现 Lexer（词法分析器）
- [x] 实现 Parser（语法分析器）
- [x] 支持基础语法规则
- [x] 完整的错误报告

### ✅ Phase 3: 解释器和内置函数

- [x] 实现 Interpreter（解释器）
- [x] 实现 12 个核心内置函数
- [x] 变量管理和求值
- [x] 函数注册机制
- [x] 数组和标量运算

### 🚧 Phase 4: 完善功能（进行中）

- [ ] 增量计算支持
- [ ] 性能优化
- [ ] 更多内置函数（30+）
- [ ] 格式化器
- [ ] 完整示例和文档

## 性能

当前性能指标（在 10 个数据点上）:
- 词法分析: < 1ms
- 语法分析: < 1ms
- 执行计算: < 5ms
- 总耗时: < 10ms

## 参考项目

- [formula-ts](https://github.com/DTrader-store/formula-ts) - TypeScript 实现版本

## 技术要求

- Go >= 1.18
- 无外部运行时依赖

## 开发

```bash
# 克隆仓库
git clone https://github.com/DTrader-store/formula-go.git
cd formula-go

# 运行测试
go test ./...

# 构建
go build

# 格式化代码
go fmt ./...

# 静态检查
go vet ./...
```

## 贡献

欢迎贡献代码！请遵循以下步骤：

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 提交 Pull Request

## 许可证

ISC License

## 联系方式

- GitHub: https://github.com/DTrader-store/formula-go
- Issues: https://github.com/DTrader-store/formula-go/issues

---

**最后更新**: 2024-11-14
**版本**: v1.0.0
**状态**: Phase 1-3 完成，Phase 4 进行中
