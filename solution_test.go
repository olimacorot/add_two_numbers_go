package add_two_numbers

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestAddTwoNumbers_ReturnSumExplode(t *testing.T) {

    tests := map[string]struct {
        l1 *ListNode
        l2 *ListNode
        want *ListNode
    }{
        "it_should_return_7_0_8": {
            l1: &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
            l2: &ListNode{5, &ListNode{6, &ListNode{4, nil}}},
            want: &ListNode{7, &ListNode{0, &ListNode{8, nil}}},
        },
        "it_should_return_0": {
            l1: &ListNode{0, nil},
            l2: &ListNode{0, nil},
            want: &ListNode{0, nil},
        },
        "it_should_return_8_9_9_9_0_0_0_1": {
            l1: &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}},
            l2: &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
            want: &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}},
        },
    }
    for name, tt := range tests {
        t.Run(name, func(t *testing.T) {
            assert.Equal(t, tt.want, addTwoNumbers(tt.l1, tt.l2))
        })
    }
}