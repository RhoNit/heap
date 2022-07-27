/* https://practice.geeksforgeeks.org/problems/adding-array-element4756/1 */

class Solution {
  public:
    int minOperations(int arr[], int n, int k) {
        priority_queue<int, vector<int>, greater<int>> pq;
        
        for(int i = 0; i < n; i++) { pq.push(arr[i]); }
        
        int count = 0;
        while(pq.top() < k) {
            int x = pq.top();
            pq.pop();
            int y = pq.top();
            pq.pop();
            
            pq.push(x+y);
            ++count;
        } 
        
        return count;
    }
};