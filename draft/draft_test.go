package draft

/* func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	list := &ListNode{
		Val: 1,
	}
	list1 := &ListNode{
		Val: 4,
	}
	list2 := &ListNode{
		Val: 3,
	}
	list3 := &ListNode{
		Val: 2,
	}
	list4 := &ListNode{
		Val: 5,
	}
	li := new(ListNode)
	l := li

	l.Next = list
	l = l.Next
	l.Next = list1
	l = l.Next
	l.Next = list2
	l = l.Next
	l.Next = list3
	l = l.Next
	l.Next = list4

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				head: li.Next,
				x:    3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(partition(tt.args.head, tt.args.x))
		})
	}
} */
