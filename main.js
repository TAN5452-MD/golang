var findMedianSortedArrays = function(nums1, nums2) {
    const s = [...nums1,...nums2].sort()
    console.log(s);
    return  (s.length & 1) === 0 ? 
     (s[s.length/2]+s[(s.length/2)-1])/2   : s[Math.floor(s.length/2)]

};

const r = findMedianSortedArrays([1,3],[2])

console.log(r)