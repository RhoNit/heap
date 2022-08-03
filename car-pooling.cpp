class Solution {
public:
    bool carPooling(vector<vector<int>>& trips, int capacity) {
        priority_queue<pair<int,int>, vector<pair<int,int>>, greater<pair<int,int>>> pq;
        for(int i = 0; i < trips.size(); i++) {
            pq.push({trips[i][1], trips[i][0]});
            pq.push({trips[i][2], -trips[i][0]});
        }
        int p = 0;
        while(pq.size() != 0) {
            p += pq.top().second;
            pq.pop();
            if(p > capacity) { return false; }
        }
        return true;
    }
};