package main

import (
	"fmt"
	"math"

	"github.com/DTrader-store/formula-go"
)

func main() {
	// 创建测试数据
	data := []*formula.MarketData{
		formula.NewMarketData(100, 105, 107, 99, 1000, 100000),
		formula.NewMarketData(105, 103, 108, 102, 1100, 110000),
		formula.NewMarketData(103, 107, 109, 101, 1200, 120000),
		formula.NewMarketData(107, 110, 112, 106, 1300, 130000),
		formula.NewMarketData(110, 108, 113, 107, 1400, 140000),
		formula.NewMarketData(108, 111, 114, 107, 1500, 150000),
		formula.NewMarketData(111, 109, 115, 108, 1600, 160000),
		formula.NewMarketData(109, 112, 116, 108, 1700, 170000),
		formula.NewMarketData(112, 115, 117, 110, 1800, 180000),
		formula.NewMarketData(115, 113, 118, 112, 1900, 190000),
	}

	engine := formula.NewFormulaEngine()

	fmt.Println("=== Formula-Go 示例 ===")
	fmt.Println()

	// 示例 1: 简单移动平均
	fmt.Println("1. 简单移动平均 (MA)")
	result, _ := engine.Run("MA5 := MA(CLOSE, 5)", data)
	printResult(result)

	// 示例 2: 标准差
	fmt.Println("\n2. 标准差 (STD)")
	result, _ = engine.Run("STD5 := STD(CLOSE, 5)", data)
	printResult(result)

	// 示例 3: 加权移动平均
	fmt.Println("\n3. 加权移动平均 (WMA)")
	result, _ = engine.Run("WMA5 := WMA(CLOSE, 5)", data)
	printResult(result)

	// 示例 4: 条件计数
	fmt.Println("\n4. 条件计数 (COUNT)")
	result, _ = engine.Run(`
		UP := CLOSE > OPEN
		UP_COUNT := COUNT(UP, 5)
	`, data)
	printResult(result)

	// 示例 5: 布林带 (Bollinger Bands) - 使用较短周期
	fmt.Println("\n5. 布林带 (Bollinger Bands)")
	result, _ = engine.Run(`
		MID := MA(CLOSE, 5)
		STDEV := STD(CLOSE, 5)
		UPPER := MID + 2 * STDEV
		LOWER := MID - 2 * STDEV
	`, data)
	printResult(result)

	// 示例 6: 突破检测
	fmt.Println("\n6. 突破检测 (CROSS + FILTER)")
	result, _ = engine.Run(`
		MA5 := MA(CLOSE, 5)
		MA10 := MA(CLOSE, 10)
		GOLDEN := CROSS(MA5, MA10)
		SIGNAL := FILTER(GOLDEN, 5)
	`, data)
	printResult(result)

	// 示例 7: EVERY 和 EXIST
	fmt.Println("\n7. 连续条件检测 (EVERY/EXIST)")
	result, _ = engine.Run(`
		UP := CLOSE > OPEN
		ALL_UP := EVERY(UP, 3)
		ANY_UP := EXIST(UP, 3)
	`, data)
	printResult(result)

	// 示例 8: 范围检测
	fmt.Println("\n8. 范围检测 (BETWEEN)")
	result, _ = engine.Run(`
		MA20 := MA(CLOSE, 5)
		IN_RANGE := BETWEEN(CLOSE, MA20 * 0.95, MA20 * 1.05)
	`, data)
	printResult(result)

	// 示例 9: 距离最后信号
	fmt.Println("\n9. 距离最后信号 (BARSLAST)")
	result, _ = engine.Run(`
		CROSS_UP := CLOSE > HIGH
		BARS := BARSLAST(CROSS_UP)
	`, data)
	printResult(result)

	// 示例 10: 复杂策略 - KDJ 指标
	fmt.Println("\n10. KDJ 指标")
	result, _ = engine.Run(`
		LOW9 := LLV(LOW, 9)
		HIGH9 := HHV(HIGH, 9)
		RSV := (CLOSE - LOW9) / (HIGH9 - LOW9) * 100
		K := SMA(RSV, 3)
		D := SMA(K, 3)
		J := 3 * K - 2 * D
	`, data)
	printResult(result)
}

func printResult(result *formula.FormulaResult) {
	for _, output := range result.Outputs {
		fmt.Printf("  %s: ", output.Name)
		for i, v := range output.Data {
			if i >= 5 && i < len(output.Data)-2 {
				if i == 5 {
					fmt.Print("... ")
				}
				continue
			}
			if math.IsNaN(v) {
				fmt.Print("NaN ")
			} else {
				fmt.Printf("%.2f ", v)
			}
		}
		fmt.Println()
	}
}
