/* minimum number of operations to be done so that 
frequencies of every characters in a string become unique or distinct */

/* https://www.codingninjas.com/codestudio/problems/minimum-character-deletion_798648?leftPanelTab=0 */


#include <bits/stdc++.h>
using namespace std;

int minCharDeletion(string str) {
    unordered_map<char,int> mp;
    for(int i = 0; i < str.size(); i++) {
        mp[str[i]]++;
    }
    
    priority_queue<pair<int, char>> pq;
    for(auto it = mp.begin(); it != mp.end(); it++) {
        pq.push({it->second, it->first});
    }
    
    int count = 0;
    while(pq.size() != 0) {
        pair<int, char> p = pq.top();
        pq.pop();
        if(pq.size() == 0) { return count; }
        if(p.first == pq.top().first) {
            if(p.first > 1) {
                pq.push({p.first-1, p.second});
            }
            ++count;
        }
    }
    return count;
}