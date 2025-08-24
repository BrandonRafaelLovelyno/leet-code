import "math"

type Rectangle struct {
    top, bottom, left, right int
}

func minimumSum(grid [][]int) int {
    mainRec := Rectangle{
        top:    0,
        bottom: len(grid) - 1,
        left:   0,
        right:  len(grid[0]) - 1,
    }

    return int(math.Min(
        float64(calHybrid(grid, mainRec)),
        math.Min(
            float64(cal3Horizontal(grid, mainRec)),
            float64(cal3Vertical(grid, mainRec)),
        ),
    ))
}

func findMinRectangle(grid [][]int, rect Rectangle) Rectangle {
    t, l := math.MaxInt32, math.MaxInt32
    b, r := -1, -1

    if rect.top > rect.bottom || rect.left > rect.right {
        return Rectangle{0, -1, 0, -1}
    }

    for i := rect.top; i <= rect.bottom; i++ {
        for j := rect.left; j <= rect.right; j++ {
            if grid[i][j] == 1 {
                if i < t { t = i }
                if i > b { b = i }
                if j < l { l = j }
                if j > r { r = j }
            }
        }
    }

    if b == -1 {
        return Rectangle{0, -1, 0, -1}
    }

    return Rectangle{top: t, bottom: b, left: l, right: r}
}

func calculateArea(rect Rectangle) int {
    if rect.bottom < rect.top || rect.right < rect.left {
        return 0
    }
    return (rect.bottom - rect.top + 1) * (rect.right - rect.left + 1)
}

func cal3Vertical(grid [][]int, rect Rectangle) int {
    ans := math.MaxInt32
    for i := rect.left; i < rect.right-1; i++ {
        for j := i + 1; j < rect.right; j++ {
            r1 := findMinRectangle(grid, Rectangle{rect.top, rect.bottom, rect.left, i})
            r2 := findMinRectangle(grid, Rectangle{rect.top, rect.bottom, i + 1, j})
            r3 := findMinRectangle(grid, Rectangle{rect.top, rect.bottom, j + 1, rect.right})
            area := calculateArea(r1) + calculateArea(r2) + calculateArea(r3)
            if area < ans {
                ans = area
            }
        }
    }
    return ans
}

func cal3Horizontal(grid [][]int, rect Rectangle) int {
    ans := math.MaxInt32
    for i := rect.top; i < rect.bottom-1; i++ {
        for j := i + 1; j < rect.bottom; j++ {
            r1 := findMinRectangle(grid, Rectangle{rect.top, i, rect.left, rect.right})
            r2 := findMinRectangle(grid, Rectangle{i + 1, j, rect.left, rect.right})
            r3 := findMinRectangle(grid, Rectangle{j + 1, rect.bottom, rect.left, rect.right})
            area := calculateArea(r1) + calculateArea(r2) + calculateArea(r3)
            if area < ans {
                ans = area
            }
        }
    }
    return ans
}

func cal2Vertical(grid [][]int, rect Rectangle) int {
    ans := math.MaxInt32
    for i := rect.left; i < rect.right; i++ {
        r1 := findMinRectangle(grid, Rectangle{rect.top, rect.bottom, rect.left, i})
        r2 := findMinRectangle(grid, Rectangle{rect.top, rect.bottom, i + 1, rect.right})
        area := calculateArea(r1) + calculateArea(r2)
        if area < ans {
            ans = area
        }
    }
    return ans
}

func cal2Horizontal(grid [][]int, rect Rectangle) int {
    ans := math.MaxInt32
    for i := rect.top; i < rect.bottom; i++ {
        r1 := findMinRectangle(grid, Rectangle{rect.top, i, rect.left, rect.right})
        r2 := findMinRectangle(grid, Rectangle{i + 1, rect.bottom, rect.left, rect.right})
        area := calculateArea(r1) + calculateArea(r2)
        if area < ans {
            ans = area
        }
    }
    return ans
}

func calHybrid(grid [][]int, rect Rectangle) int {
    ans := math.MaxInt32

    for i := rect.left; i < rect.right; i++ {
        v1 := Rectangle{rect.top, rect.bottom, rect.left, i}
        v2 := Rectangle{rect.top, rect.bottom, i + 1, rect.right}

        areaA := cal2Horizontal(grid, v1) + calculateArea(findMinRectangle(grid, v2))
        if areaA < ans {
            ans = areaA
        }

        areaB := calculateArea(findMinRectangle(grid, v1)) + cal2Horizontal(grid, v2)
        if areaB < ans {
            ans = areaB
        }
    }

    for i := rect.top; i < rect.bottom; i++ {
        h1 := Rectangle{rect.top, i, rect.left, rect.right}
        h2 := Rectangle{i + 1, rect.bottom, rect.left, rect.right}

        areaA := cal2Vertical(grid, h1) + calculateArea(findMinRectangle(grid, h2))
        if areaA < ans {
            ans = areaA
        }

        areaB := calculateArea(findMinRectangle(grid, h1)) + cal2Vertical(grid, h2)
        if areaB < ans {
            ans = areaB
        }
    }

    return ans
}