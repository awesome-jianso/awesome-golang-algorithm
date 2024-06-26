package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs *TreeNode
		expect *TreeNode
	}{
		{"TestCase1", &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 10},
			},
			Right: &TreeNode{
				Val:   9,
				Right: &TreeNode{Val: 7},
			},
		}, &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 7},
			},
			Right: &TreeNode{
				Val:   0,
				Right: &TreeNode{Val: 11},
			},
		}},
		{"TestCase1", &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		}, &TreeNode{
			Val:   0,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 0},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
