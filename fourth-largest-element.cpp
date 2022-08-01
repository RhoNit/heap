#include <bits/stdc++.h>
using namespace std;

int getFourthLargest(int arr[], int n) {
    priority_queue<int> pq;
    for(int i = 0; i < n; i++) {
        pq.push(arr[i]);
    }
    int t = 4;
    while(t-1 > 0) {
        pq.pop();
        t--;
    }
    if(pq.size() != 0) { return pq.top(); }
    return -2147483648;
}