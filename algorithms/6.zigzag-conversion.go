package main

/**
 * <p>The string <code>&quot;PAYPALISHIRING&quot;</code> is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)</p>

<pre>
P   A   H   N
A P L S I I G
Y   I   R
</pre>

<p>And then read line by line: <code>&quot;PAHNAPLSIIGYIR&quot;</code></p>

<p>Write the code that will take a string and make this conversion given a number of rows:</p>

<pre>
string convert(string s, int numRows);
</pre>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;PAYPALISHIRING&quot;, numRows = 3
<strong>Output:</strong> &quot;PAHNAPLSIIGYIR&quot;
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;PAYPALISHIRING&quot;, numRows = 4
<strong>Output:</strong> &quot;PINALSIGYAHRPI&quot;
<strong>Explanation:</strong>
P     I    N
A   L S  I G
Y A   H R
P     I
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;A&quot;, numRows = 1
<strong>Output:</strong> &quot;A&quot;
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 1000</code></li>
	<li><code>s</code> consists of English letters (lower-case and upper-case), <code>&#39;,&#39;</code> and <code>&#39;.&#39;</code>.</li>
	<li><code>1 &lt;= numRows &lt;= 1000</code></li>
</ul>

**/
/**
 * "PAYPALISHIRING"
3
**/
func convert(s string, numRows int) string {
	queue := make([][]rune, numRows)
	for i := 0; i < numRows; i++ {
		queue[i] = make([]rune, len(s))
	}
	rowIdx := 0
	colIdx := 0
	for _, c := range s {
		queue[rowIdx][colIdx] = c
		// next idx
		if numRows == 1 {
			colIdx = colIdx + 1
			continue
		}
		if colIdx%(numRows-1) != 0 {
			rowIdx = rowIdx - 1
			colIdx = colIdx + 1
			continue
		}
		if rowIdx == (numRows - 1) {
			rowIdx = rowIdx - 1
			colIdx = colIdx + 1
			continue
		}
		rowIdx = rowIdx + 1
	}

	var res []rune
	for _, row := range queue {
		for _, col := range row {
			// empty
			if col == 0 {
				continue
			}
			res = append(res, col)
		}
	}

	return string(res)
}
