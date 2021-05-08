package main

import (
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name: "",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. swapPairs() = %v, want %v", tt.name, got, tt.want)
			var index int
			for {
				if got == nil && tt.want == nil {
					break
				} else if got == nil || tt.want == nil {
					t.Errorf("%v got:%v, want:%v", index, got, tt.want)
					break
				}
				t.Logf("got: %v", got.Val)
				t.Logf("want: %v", tt.want.Val)
				if got.Val != tt.want.Val {
					t.Errorf("%v got:%v, want:%v", index, got.Val, tt.want.Val)
				}
				got = got.Next
				tt.want = tt.want.Next
				index++
			}
		}
	}
}
