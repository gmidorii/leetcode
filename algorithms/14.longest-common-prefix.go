package main

/**
 * <p>Write a function to find the longest common prefix string amongst an array of strings.</p>

<p>If there is no common prefix, return an empty string <code>&quot;&quot;</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> strs = [&quot;flower&quot;,&quot;flow&quot;,&quot;flight&quot;]
<strong>Output:</strong> &quot;fl&quot;
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> strs = [&quot;dog&quot;,&quot;racecar&quot;,&quot;car&quot;]
<strong>Output:</strong> &quot;&quot;
<strong>Explanation:</strong> There is no common prefix among the input strings.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>0 &lt;= strs.length &lt;= 200</code></li>
	<li><code>0 &lt;= strs[i].length &lt;= 200</code></li>
	<li><code>strs[i]</code> consists of only lower-case English letters.</li>
</ul>

**/
/**
 * ["flower","flow","flight"]
**/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	match := func(org, tgt string) string {
		rt := []rune(tgt)
		var res []rune
		for i, o := range org {
			if i >= len(tgt) {
				break
			}
			if o != rt[i] {
				break
			}
			res = append(res, o)
		}
		return string(res)
	}
	var common string
	for i, v := range strs {
		if i == 0 {
			common = v
			continue
		}
		common = match(common, v)
	}
	return common
}
