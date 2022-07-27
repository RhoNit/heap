/*   first approach (using Min Heap)   */

class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        priority_queue<int, vector<int>, greater<int>> pq;
        //min heap
        for(int i = 0; i < nums.size(); i++) {
            if(pq.size() < k) { pq.push(nums[i]); }
            else {
                if(pq.top() < nums[i]) {
                    pq.pop();
                    pq.push(nums[i]);
                }
            }
        }
        return pq.top();
    }
};

/* second approach (using Max Heap)  */

class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        priority_queue<int> pq;
        //used '-' sign for making the max heap act as a min heap
        for(int i = 0; i < k; i++) { pq.push(-nums[i]); }
        for(int i = k; i < nums.size(); i++) {
            if(-pq.top() < nums[i]) {
                pq.pop();
                pq.push(-nums[i]);
            }
        }
        return -pq.top();
    }
};