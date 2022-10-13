// https://practice.geeksforgeeks.org/problems/game-with-string4100/1

class Solution{
public:
    int minValue(string s, int k){
        unordered_map<char, int> mp;
        priority_queue<int> pq;
        int sum = 0;
        
        for(int i = 0; i < s.length(); i++) { mp[s[i]]++; }
        
        for(auto x: mp) { pq.push(x.second); }
        
        while(k--) {
            int maxi = pq.top();
            pq.pop();
            pq.push(--maxi);
        }
        
        while(pq.size() != 0) {
            sum += pq.top()*pq.top();
            pq.pop();
        }
        
        return sum;
    }
};