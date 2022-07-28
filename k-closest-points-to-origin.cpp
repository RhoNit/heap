class Solution {
public:
    vector<vector<int>> kClosest(vector<vector<int>>& points, int k) {
        priority_queue<pair<int, pair<int, int>>> pq;
        
        for(int i = 0; i < points.size(); i++) {
            pq.push({points[i][0]*points[i][0] + points[i][1]*points[i][1], {points[i][0], points[i][1]}});
            if(pq.size() > k) { pq.pop(); }
        }
        
        points.clear();
        
        while(pq.size() != 0) {
            auto p = pq.top().second;
            pq.pop();
            
            int x = p.first;
            int y = p.second;
            
            vector<int> v;
            v.push_back(x);
            v.push_back(y);
            
            points.push_back(v);
        }
        
        return points;
    }
};