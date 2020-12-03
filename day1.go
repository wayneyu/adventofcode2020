package main

import (
    "fmt"
    "strconv"
    "bufio"
    "os"
    "sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
    input_path := "./input/day1"
    file, err := os.Open(input_path)
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)

    // Read numbers from input file
    nums := make([]int, 0)
    for scanner.Scan() {
        v, err := strconv.Atoi(scanner.Text())
        check(err)

        nums = append(nums, v)
    }

    // Part 1
    seen := make(map[int]struct{})
    target := 2020
    for i:= 0; i < len(nums); i++ {
        this := nums[i]
        that := target - this
        _, prs := seen[that]
        if prs {
            fmt.Printf("Part 1: %d\n", this * that)
        } else {
            seen[this] = struct{}{}
        }
    }

    // Part 2
    sort.Ints(nums)
    for i:= 0; i < len(nums); i++ {
        remain := target - nums[i]
        j := i + 1
        k := len(nums) - 1
        for j <= k {
            if remain == nums[j] + nums[k] {
                fmt.Printf("Part 2: %d\n", nums[i] * nums[j] * nums[k])
                break
            }
            if remain > nums[j] + nums[k] {
                j++
            } else {
                k--
            }
        }
    }

}
