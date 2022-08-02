class Solution {
public:
    int lastStoneWeight(vector<int>& stones) {
        priority_queue<int> pq;
        for(int i = 0; i < stones.size(); i++) { pq.push(stones[i]); }
        while(pq.size() > 1) {
            int max1 = pq.top();
            pq.pop();
            int max2 = pq.top();
            pq.pop();
            if(max1 != max2) {
                pq.push(max1-max2);
            }
        }
        if(pq.size()) { return pq.top(); }
        return 0;
    }
};