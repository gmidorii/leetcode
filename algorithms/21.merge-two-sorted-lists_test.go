package main

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				l1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "",
			args: args{
				l1: nil,
				l2: nil,
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				l1: nil,
				l2: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			want: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
		{
			name: "",
			args: args{
				l1: &ListNode{
					Val: -10,
					Next: &ListNode{
						Val: -6,
						Next: &ListNode{
							Val: -6,
							Next: &ListNode{
								Val: -6,
								Next: &ListNode{
									Val: -3,
									Next: &ListNode{
										Val: 5,
									},
								},
							},
						},
					},
				},
				l2: nil,
			},
			want: &ListNode{
				Val: -10,
				Next: &ListNode{
					Val: -6,
					Next: &ListNode{
						Val: -6,
						Next: &ListNode{
							Val: -6,
							Next: &ListNode{
								Val: -3,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		if got := mergeTwoLists2(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
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
