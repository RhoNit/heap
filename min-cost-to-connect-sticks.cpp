/*   https://www.codingninjas.com/codestudio/problems/minimum-cost-to-connect-sticks_1402396?leftPanelTab=1   */

#include <bits/stdc++.h>
using namespace std;

long long int minimumCostToConnectSticks(vector<int> &arr) {
    priority_queue<int, vector<int>, greater<int>> pq;
    for(int i = 0; i < arr.size(); i++) { pq.push(arr[i]); }
    int sum = 0;
    while(pq.size() > 1) {
        int min1 = pq.top();
        pq.pop();
        int min2 = pq.top();
        pq.pop();
        sum += min1 + min2;
        pq.push(min1+min2);
    }
    return sum;
}