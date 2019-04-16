package main

import "fmt"

func main() {
    nums := []int{2, 7, 11, 15}
    fmt.Println(twoNum(nums, 9))
}

func twoNum(nums []int, target int) []int {
    //创建 map,用来保存查询的数据, 按照 map[value]key的顺序保存, 例如 map[2] = 0
    m := make(map[int]int, len(nums))
    for key, value := range nums {
        //遍历切片, 如果找到 target - value, 例如9-2=7. 如果存在 map[7], 则说明查找成功, 返回索引.
        if j, ok := m[target-value]; ok {
            return []int{j, key}
        }
        m[value] = key
    }
    return nil
}
