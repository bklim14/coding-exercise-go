package exercise

/*Given the string, check if it is a palindrome.

Example

For inputString = "aabaa", the output should be
checkPalindrome(inputString) = true;
For inputString = "abac", the output should be
checkPalindrome(inputString) = false;
For inputString = "a", the output should be
checkPalindrome(inputString) = true.
Input/Output

[execution time limit] 4 seconds (go)

[input] string inputString

A non-empty string consisting of lowercase characters.

Guaranteed constraints:
1 ≤ inputString.length ≤ 105.

[output] boolean

true if inputString is a palindrome, false otherwise.
*/

func checkPalindrome(inputString string) bool {
	rStr := []rune(inputString)

	for i, j := 0, len(rStr)-1; i < j; i, j = i+1, j-1 {
		rStr[i], rStr[j] = rStr[j], rStr[i]
	}

	if inputString == string(rStr) {
		return true
	}

	return false
}
