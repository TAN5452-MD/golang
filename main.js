/**
 * @param {string[]} strs
 * @return {string[][]}
 */
    var groupAnagrams = function(strs) {
        let result = {};
        for (let i = 0; i < strs.length; i++) {
            let str = strs[i];
            let sortedStr = str.split('').sort().join('');
            if (!result[sortedStr]) {
                result[sortedStr] = [];
            }
            result[sortedStr].push(str);
        }
        console.log(result);
        return Object.values(result);
    };

console.log(groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]
));