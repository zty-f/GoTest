package samber_lo

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
	"strconv"
	"strings"
	"testing"
)

type BooklistLecture struct {
	ID        int
	LectureID int
}

// 过滤两个数组不重复的元素集合
func TestFilterAndContainsBy(t *testing.T) {
	booklistLectures := []*BooklistLecture{
		{ID: 1, LectureID: 101},
	}
	lectureIDs := []int{101, 102, 103, 104}
	var filteredLectureIDs []int
	if len(booklistLectures) > 0 {
		// 不能重复添加，找出lectureIDs中不存在于booklistLectures的lectureID
		filteredLectureIDs = lo.Filter(lectureIDs, func(lectureID int, _ int) bool {
			return !lo.ContainsBy(booklistLectures, func(bll *BooklistLecture) bool {
				return bll.LectureID == lectureID
			})
		})
	} else {
		// 直接插入
		filteredLectureIDs = lectureIDs
	}
	fmt.Println(lectureIDs)
	fmt.Println(filteredLectureIDs)
}

// 遍历集合并返回符合条件的所有元素数组。
func TestFilter(t *testing.T) {
	even := lo.Filter([]int{1, 2, 3, 4}, func(x int, index int) bool {
		return x%2 == 0 // 为true的元素会被返回
	})
	fmt.Println(even) // []int{2, 4}
}

// 操作一种类型的切片并将其转换为另一种类型的切片
func TestMap(t *testing.T) {
	// 单协程处理
	res := lo.Map([]int64{1, 2, 3, 4}, func(x int64, index int) string {
		return strconv.FormatInt(x, 10)
	})
	fmt.Printf("%#v\n", res) // []string{"1", "2", "3", "4"}
	// 并发处理
	r2 := parallel.Map([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	fmt.Printf("%#v\n", r2) // []string{"1", "2", "3", "4"}
}

// 操作切片并将其转换为具有唯一值的另一种类型的切片。 去除重复值并且映射其他类型
func TestUniqMap(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	users := []User{{Name: "Alex", Age: 10}, {Name: "Alex", Age: 12}, {Name: "Bob", Age: 11}, {Name: "Alice", Age: 20}}

	names := lo.UniqMap(users, func(u User, index int) string {
		return u.Name
	})
	fmt.Printf("%#v\n", names) // []string{"Alex", "Bob", "Alice"}
}

// 过滤映射
// 返回使用给定的回调函数进行过滤和映射后获得的切片。
// 回调函数应该返回两个值：映射操作的结果和是否应该包含结果元素。 满足回调方法条件的元素会被返回
func TestFilterMap(t *testing.T) {
	matching := lo.FilterMap([]string{"cpu", "gpu", "mouse", "keyboard"}, func(x string, _ int) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return "xpu", true
		}
		return "", false
	})
	fmt.Printf("%#v\n", matching) // []string{"xpu", "xpu"}
}

// 操作切片并将其变换并展平为其他类型的切片。变换函数可以返回一个切片或一个nil，如果返回nil则不会将任何值添加到最终的切片中。
// 可以对于切片元素进行扩展 一个元素变多个相邻的元素
func TestFlatMap(t *testing.T) {
	res := lo.FlatMap([]int64{0, 1, 2}, func(x int64, _ int) []string {
		return []string{
			strconv.FormatInt(x, 10),
			strconv.FormatInt(x, 10),
		}
	})
	fmt.Printf("%#v\n", res) // []string{"0", "0", "1", "1", "2", "2"}
}

// 将集合简化为单个值。该值是通过累加器函数对集合中每个元素的运行结果进行累加计算得出的。每次后续调用都会获得前一次调用的返回值。
func TestReduce(t *testing.T) {
	// 计算集合的和 agg是累加器的值，item是当前元素的值，index是当前元素的索引
	sum := lo.Reduce([]int{1, 2, 3, 4}, func(agg int, item int, index int) int {
		return agg + item
	}, 0)
	fmt.Println(sum) // 10

	sum = lo.Reduce([]int{1, 2, 3, 4}, func(agg int, item int, index int) int {
		return agg + item
	}, 5) // 定义累加器的初始值为5
	fmt.Println(sum) // 15
}

// 类似lo.Reduce，只是它从右到左迭代集合元素。 注意迭代的集合的类型可以多变，根据不同的场景进行更换
func TestReduceRight(t *testing.T) {
	result := lo.ReduceRight([][]int{{0, 1}, {2, 3}, {4, 5}}, func(agg []int, item []int, _ int) []int {
		return append(agg, item...)
	}, []int{})
	fmt.Println(result) // []int{4, 5, 2, 3, 0, 1}
}

// 遍历集合的元素并对每个元素调用函数。
func TestForEach(t *testing.T) {
	// 单协程处理
	lo.ForEach([]int{1, 2, 3}, func(x int, index int) {
		// 这里的x是当前元素的值，index是当前元素的索引
		// 你可以在这里执行任何操作，比如打印、修改等-调用其他函数
		fmt.Printf("index: %d, value: %d\n", index, x)
	})
	// 并发处理
	parallel.ForEach([]int{1, 2, 3}, func(x int, index int) {
		fmt.Printf("index: %d, value: %d\n", index, x)
	})
}

// 遍历集合元素并为每个元素集合调用 iteratee，返回值决定继续或中断，就像 do while() 一样。
func TestForEachWhile(t *testing.T) {
	list := []int64{1, 2, -42, 4}
	// 单协程处理 return false 终止遍历
	lo.ForEachWhile(list, func(x int64, index int) bool {
		// 当x小于0时，终止遍历
		if x < 0 {
			return false
		}
		// 调用其他函数
		fmt.Println(x)
		return true
	})
	// 1
	// 2
}

// Times 调用迭代器 n 次，返回每次调用结果的数组。迭代器以索引作为参数进行调用。
func TestTimes(t *testing.T) {
	// 单协程处理 表示调用10次迭代器，每次调用迭代器时传入的参数是索引
	// 顺序执行10次内部的方法
	res := lo.Times(10, func(index int) string {
		return strconv.FormatInt(int64(index), 10)
	})
	fmt.Printf("%#v\n", res) // []string{"0", "1", "2"}
	// 并发处理
	// 并发执行10次内部的方法，并且内部控制了响应的顺序，得到的结果和上面是一样的
	res = parallel.Times(10, func(index int) string {
		return strconv.FormatInt(int64(index), 10)
	})
	fmt.Printf("%#v\n", res) // []string{"0", "1", "2"}
}

// 返回一个数组的无重复版本，其中仅保留每个元素的第一次出现。结果值的顺序由它们在数组中出现的顺序决定。
func TestUniq(t *testing.T) {
	res := lo.Uniq([]int{1, 2, 3, 1, 4, 2, 3})
	fmt.Printf("%#v\n", res) // []int{1, 2, 3, 4}
}

// 返回一个数组的无重复版本，其中仅保留每个元素的第一次出现。结果值的顺序由它们在数组中出现的顺序决定。它接受iteratee对数组中每个元素调用的方法来生成计算唯一性的标准。
func TestUniqBy(t *testing.T) {
	// 通过对每个元素调用iteratee方法来计算唯一性---调用方法得到的结果来判断唯一性
	uniqValues := lo.UniqBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})
	fmt.Printf("%#v\n", uniqValues) // []int{0, 1, 2}

	uniqValues = lo.UniqBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 2 // 0,1,0,1,0,1 所以相当于对于结果数组取唯一
	})
	fmt.Printf("%#v\n", uniqValues) // []int{0, 1}
}

// 返回一个由通过 iteratee 运行 collection 中每个元素的结果生成的键组成的对象。
func TestGroupBy(t *testing.T) {
	// 单协程处理
	// 通过对每个元素调用iteratee方法来计算结果值，值相同的元素会被分到同一个数组中，顺序按照第一次出现固定，每次不会变。
	groups := lo.GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})
	fmt.Printf("%#v\n", groups) // map[int][]int{0: []int{0, 3}, 1: []int{1, 4}, 2: []int{2, 5}}
	// 并发处理
	groups = parallel.GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})
	fmt.Printf("%#v\n", groups) // map[int][]int{0: []int{0, 3}, 1: []int{1, 4}, 2: []int{2, 5}}
}

// 返回一个数组，该数组元素被分成长度为 size 的组。如果数组无法被均匀分割，则最终的块将是剩余元素。
func TestChunk(t *testing.T) {
	chunks := lo.Chunk([]int{1, 2, 3, 4, 5, 6}, 2)
	fmt.Printf("%#v\n", chunks) // [][]int{{1, 2}, {3, 4}, {5, 6}}
	chunks = lo.Chunk([]int{1, 2, 3, 4, 5, 6}, 4)
	fmt.Printf("%#v\n", chunks) // [][]int{{1, 2, 3, 4}, {5, 6}}
}

// 返回一个按组拆分的元素数组。分组值的顺序由它们在集合中出现的顺序决定。分组是通过迭代器运行集合中每个元素的结果生成的。
func TestPartitionBy(t *testing.T) {
	// 单协程处理
	// 每个元素调用内部方法，返回的值相同的元素会被分到同一个数组中，顺序按照第一次出现固定，每次不会变。
	partitions := lo.PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, func(x int) string {
		if x < 0 {
			return "negative"
		} else if x%2 == 0 {
			return "even"
		}
		return "odd"
	})
	fmt.Printf("%#v\n", partitions) // [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}}
	// 并发处理
	partitions = parallel.PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, func(x int) string {
		if x < 0 {
			return "negative"
		} else if x%2 == 0 {
			return "even"
		}
		return "odd"
	})
	fmt.Printf("%#v\n", partitions) // [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}}
}

// 返回单层深度的数组。 把二维数组变成一维数组
func TestFlatten(t *testing.T) {
	flat := lo.Flatten([][]int{{1, 2}, {3, 4, 5, 6}})
	fmt.Printf("%#v\n", flat) // []int{1, 2, 3, 4}
}

// 循环交替输入切片并按顺序将索引处的值附加到结果中。
func TestInterleave(t *testing.T) {
	interleaved := lo.Interleave([]int{1, 4, 7}, []int{2, 5, 8}, []int{3, 6, 9})
	fmt.Printf("%#v\n", interleaved) // []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	interleaved = lo.Interleave([]int{1}, []int{2, 5, 8}, []int{3, 6}, []int{4, 7, 9, 10})
	fmt.Printf("%#v\n", interleaved) // []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

// 返回已打乱顺序的值数组。使用 Fisher-Yates 打乱顺序算法。-- 洗牌算法
func TestShuffle(t *testing.T) {
	randomOrder := lo.Shuffle([]int{0, 1, 2, 3, 4, 5})
	fmt.Printf("%#v\n", randomOrder) // []int{4, 2, 0, 5, 3, 1}
}

// 反转数组，使第一个元素成为最后一个，第二个元素成为倒数第二个，依此类推。
func TestReverse(t *testing.T) {
	reversed := lo.Reverse([]int{1, 2, 3, 4, 5})
	fmt.Printf("%#v\n", reversed) // []int{5, 4, 3, 2, 1}
}

type foo struct {
	bar string
}

func (f foo) Clone() foo {
	return foo{f.bar}
}

// 用值填充数组元素----生成一个长度相同的数组，元素值都为传入的值
func TestFill(t *testing.T) {
	// 填充的值需要实现Clone方法
	initializedSlice := lo.Fill([]foo{foo{"a"}, foo{"a"}}, foo{"b"})
	fmt.Printf("%#v\n", initializedSlice) // []foo{foo{"b"}, foo{"b"}}
}

// 构建具有 N 个初始值副本的切片。
func TestRepeat(t *testing.T) {
	repeatedSlice := lo.Repeat(3, foo{"a"})
	fmt.Printf("%#v\n", repeatedSlice) // []samber_lo.foo{samber_lo.foo{bar:"a"}, samber_lo.foo{bar:"a"}, samber_lo.foo{bar:"a"}}
}
