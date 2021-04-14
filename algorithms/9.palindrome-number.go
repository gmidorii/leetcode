package main

/**
 * <p>Given an integer <code>x</code>, return <code>true</code> if <code>x</code> is palindrome integer.</p>

<p>An integer is a <strong>palindrome</strong> when it reads the same backward as forward. For example, <code>121</code> is palindrome while <code>123</code> is not.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> x = 121
<strong>Output:</strong> true
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> x = -121
<strong>Output:</strong> false
<strong>Explanation:</strong> From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> x = 10
<strong>Output:</strong> false
<strong>Explanation:</strong> Reads 01 from right to left. Therefore it is not a palindrome.
</pre>

<p><strong>Example 4:</strong></p>

<pre>
<strong>Input:</strong> x = -101
<strong>Output:</strong> false
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>-2<sup>31</sup>&nbsp;&lt;= x &lt;= 2<sup>31</sup>&nbsp;- 1</code></li>
</ul>

<p>&nbsp;</p>
<strong>Follow up:</strong> Could you solve it without converting the integer to a string?
**/
/**
 * 121
**/
func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 {
		return false
	}
	l := []int{}
	for {
		balance := x % 10
		l = append(l, balance)
		x = (x - balance) / 10
		if x == 0 {
			break
		}
	}
	r := make([]int, len(l))
	j := 0
	for i := len(l) - 1; i >= 0; i-- {
		r[i] = l[j]
		j++
	}

	for i, v := range l {
		if r[i] != v {
			return false
		}
	}
	return true
}
