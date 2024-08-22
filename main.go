package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    var numbers []int
    for {
        fmt.Print("Enter a number (or type 'exit' to quit): ")
        input, _ := reader.ReadString('\n')
        input = input[:len(input)-1] // Remove newline

        if input == "exit" {
            break
        }

        var number int
        _, err := fmt.Sscanf(input, "%d", &number)
        if err != nil {
            fmt.Println("Invalid input, please enter a number.")
            continue
        }

        numbers = append(numbers, number)

        if len(numbers) > 1 {
            lower, upper := calculateRange(numbers)
            fmt.Printf("%d --> range for the next input: %d %d\n", number, lower, upper)
        }
    }
}

func calculateRange(numbers []int) (int, int) {
    n := len(numbers)
    if n < 2 {
        return 0, 0
    }

    mean := calculateMean(numbers)
    stdDev := calculateStdDev(numbers, mean)

    lower := int(math.Max(float64(mean-2*stdDev), 0))
    upper := int(mean + 2*stdDev)

    return lower, upper
}

func calculateMean(numbers []int) float64 {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return float64(sum) / float64(len(numbers))
}

func calculateStdDev(numbers []int, mean float64) float64 {
    sum := 0.0
    for _, num := range numbers {
        sum += math.Pow(float64(num)-mean, 2)
    }
    return math.Sqrt(sum / float64(len(numbers)))
}


/*
package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    var numbers []int

    for {
        fmt.Print("Enter a number (or 'exit' to quit): ")
        input, _ := reader.ReadString('\n')
        if input == "exit\n" {
            break
        }

        var number int
        if _, err := fmt.Sscanf(input, "%d", &number); err != nil {
            fmt.Println("Invalid input. Please enter a number.")
            continue
        }
        numbers = append(numbers, number)

        if len(numbers) > 1 {
            lower, upper := calculateRange(numbers)
            fmt.Printf("%d --> range for the next input: %d %d\n", number, lower, upper)
        }
    }
}

func calculateRange(nums []int) (int, int) {
    mean := calculateMean(nums)
    stdDev := calculateStdDev(nums, mean)
    return int(math.Max(mean-2*stdDev, 0)), int(mean + 2*stdDev)
}

func calculateMean(nums []int) float64 {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return float64(sum) / float64(len(nums))
}

func calculateStdDev(nums []int, mean float64) float64 {
    var sum float64
    for _, num := range nums {
        sum += math.Pow(float64(num)-mean, 2)
    }
    return math.Sqrt(sum / float64(len(nums)))
}
	*/

/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func linearRegressionFormula(xValues, yValues []float64, curX float64) (int, int) {
	n := len(xValues)
	var sumX, sumY, sumXY, sumX2 float64

	for i, x := range xValues {
		y := yValues[i]
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
	}

	slope := (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumX2 - sumX*sumX)
	intercept := (sumY - slope*sumX) / float64(n)

	lowY := (slope*curX + intercept) - 6
	upY := (slope*curX + intercept) + 6

	return int(lowY), int(upY)
}

func main() {
	yVals := []float64{}
	xVals := []float64{}
	var low, up int
	start := 0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		xVals = append(xVals, float64(start))
		yVals = append(yVals, float64(num))
		start++

		if start != 1 {
			low, up = linearRegressionFormula(xVals, yVals, float64(start))
		} else {
			low = num
			up = num + 10
		}

		fmt.Println(low, up)
	}
}

*/

 