package util

import (
	"log"
	"reflect"
	"testing"
)

func TestTreeNode(t *testing.T) {
	ts := [][]string{
		[]string{"1", "2", "3", "null", "null", "4", "5", "6", "7"},
		[]string{"1", "2", "3", "null", "null", "4", "5"},
		[]string{"1", "3", "2", "5", "3", "null", "9"},
		[]string{"1", "2", "3"},
		[]string{}}
	for _, v := range ts {
		tree := Strs2TreeNode(v)
		if tree == nil && len(v) != 0 {
			continue
		}

		source := TreeSource(tree)
		if !reflect.DeepEqual(source, v) {
			t.Errorf("gen tree not right,source_in:%s,source_gen:%s", v, source)
		}

		source = Tree2Strs(tree)
		if !reflect.DeepEqual(source, v) {
			t.Errorf("gen tree not right,source_in:%s,source_gen:%s", v, source)
		}
	}
}

type travalTreeTestCase struct {
	tree      []string
	preOrder  []string
	inOrder   []string
	postOrder []string
}

func TestTravalTree(t *testing.T) {
	ts := []travalTreeTestCase{
		{
			/*           1
			 *          / \
			 *        2    3
			 *       / \   / \
			 *            4   5
			 *		    /   \
			 * 		  6       7
			 */
			tree:      []string{"1", "2", "3", "null", "null", "4", "5", "6", "7"},
			preOrder:  []string{"1", "2", "3", "4", "6", "7", "5"},
			inOrder:   []string{"2", "1", "6", "4", "7", "3", "5"},
			postOrder: []string{"2", "6", "7", "4", "5", "3", "1"},
		},
		/*				  1
		 *				/  \
		 *			   2    3
		 *			  / \  / \
		 *			 4  5 6   7
		 */
		{
			tree:      []string{"1", "2", "3", "4", "5", "6", "7"},
			preOrder:  []string{"1", "2", "4", "5", "3", "6", "7"},
			inOrder:   []string{"4", "2", "5", "1", "6", "3", "7"},
			postOrder: []string{"4", "5", "2", "6", "7", "3", "1"},
		},
	}
	for _, tc := range ts {
		log.Println("tree:           ", tc.tree)
		tree := Strs2TreeNode(tc.tree)
		if tree == nil && len(tc.tree) == 0 {
			continue
		}
		pre := Tree2PreOrder(tree)
		log.Println("tree,pre_order: ", pre)
		if !reflect.DeepEqual(pre, tc.preOrder) {
			t.Errorf("preOrder not right,source_in:%s,source_gen:%s", tc.preOrder, pre)
		}
		in := Tree2InOrder(tree)
		log.Println("tree,in_order:  ", in)
		if !reflect.DeepEqual(in, tc.inOrder) {
			t.Errorf("inOrder not right,source_in:%s,source_gen:%s", tc.inOrder, in)
		}
		post := Tree2PostOrder(tree)
		log.Println("tree,post_order:", post)
		if !reflect.DeepEqual(post, tc.postOrder) {
			t.Errorf("postOrder not right,source_in:%s,source_gen:%s", tc.postOrder, post)
		}
	}
}
