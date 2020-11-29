1. 两数之和

   #### 

   给定一个整数数组 `nums` 和一个目标值 `target`，请你在该数组中找出和为目标值的那 **两个** 整数，并返回他们的数组下标。

   你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

   示例

   ```
   给定 nums = [2, 7, 11, 15], target = 9
   
   因为 nums[0] + nums[1] = 2 + 7 = 9
   所以返回 [0, 1]
   ```

   

   **哈希表**解决更好：

   思路及算法

   注意到方法一的时间复杂度较高的原因是寻找 target - x 的时间复杂度过高。因此，我们需要一种更优秀的方法，能够快速寻找数组中是否存在目标元素。如果存在，我们需要找出它的索引。

   使用哈希表，可以将寻找 target - x 的时间复杂度降低到从 O(N)O(N)O(N) 降低到 O(1)O(1)O(1)。

   这样我们创建一个哈希表，对于每一个 x，我们首先查询哈希表中是否存在 target - x，然后将 x 插入到哈希表中，即可保证不会让 x 和自己匹配。

   ```java
   class Solution {
       public int[] twoSum(int[] nums, int target) {
           Map<Integer, Integer> hashtable = new HashMap<Integer, Integer>();
           for (int i = 0; i < nums.length; ++i) {
               if (hashtable.containsKey(target - nums[i])) {
                   return new int[]{hashtable.get(target - nums[i]), i};
               }
               hashtable.put(nums[i], i);
           }
           return new int[0];
       }
   }
   ```

   

   

   