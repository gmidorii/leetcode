package main

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
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
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
				n: 2,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name: "2",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
				n: 1,
			},
			want: nil,
		},
		{
			name: "3",
			args: args{
				head: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
				n: 1,
			},
			want: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	for _, tt := range tests {
		// if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
		// 	t.Errorf("%q. removeNthFromEnd() = %v, want %v", tt.name, got, tt.want)
		// }
		if got := removeNthFromEnd2(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. removeNthFromEnd2() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
