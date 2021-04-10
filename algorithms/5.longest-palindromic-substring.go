package main

/**
 * <p>Given a string <code>s</code>, return&nbsp;<em>the longest palindromic substring</em> in <code>s</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;babad&quot;
<strong>Output:</strong> &quot;bab&quot;
<strong>Note:</strong> &quot;aba&quot; is also a valid answer.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;cbbd&quot;
<strong>Output:</strong> &quot;bb&quot;
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;a&quot;
<strong>Output:</strong> &quot;a&quot;
</pre>

<p><strong>Example 4:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;ac&quot;
<strong>Output:</strong> &quot;a&quot;
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 1000</code></li>
	<li><code>s</code> consist of only digits and English letters (lower-case and/or upper-case),</li>
</ul>
**/
/**
 * "babad"
**/
func longestPalindrome(s string) string {
	rev := make([]rune, len(s))
	for i, c := range s {
		rev[len(s)-i-1] = c
	}
	rs := []rune(s)
	results := []string{}
	for i := range rs {
		for j := range rev {
			if rs[i] == rev[j] {
				ans := []rune{rs[i]}
				x := 1
				for {
					if len(rev)-(j+x)-1 == i {
						ans = append(ans, rs[i+x])
						results = append(results, string(ans))
						break
					}
					if j+x >= len(rev)-1 {
						if len(ans) == 1 {
							results = append(results, string(ans))
						}
						break
					}
					if i+x >= len(rs)-1 {
						break
					}
					if rs[i+x] == rev[j+x] {
						ans = append(ans, rs[i+x])
						x = x + 1
						continue
					}
					break
				}
			}
		}
	}
	var res string
	for _, result := range results {
		if len(result) > len(res) {
			res = result
		}
	}
	return res
}
