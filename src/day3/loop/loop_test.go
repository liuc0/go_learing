package loop

import "testing"
/**
 * for循环的遍历方式
 */
func TestWhileLoop(t *testing.T)  {
	n := 0;
	for n<5  {
		t.Log(n);
		n++;
	}
}
