#include <bits/stdc++.h>
using namespace std;

void sortKsorted(vector<int> &v, int n, int k) {
    priority_queue<int, vector<int>, greater<int>> pq;
    int x = 0;
    
    for(int i = 0; i < v.size(); i++) {
        if(pq.size() <= k) {
            pq.push(v[i]);
        }
        else {
            v[x] = pq.top();
            x++;
            pq.pop();
            pq.push(v[i]);
        }
    }
    while(x < n and pq.size() != 0) {
        v[x] = pq.top();
        pq.pop();
        x++;
    }
    
    for(int i = 0; i < v.size(); i++) { cout << v[i] << " ";}
    cout << endl;
}

int main() {
	int t;
	cin >> t;
	while(t--) {
	    int n, k;
	    cin >> n >> k;
	    
	    vector<int> vec(n);
	    for(int i = 0; i < n; i++) { cin >> vec[i]; }
	    
	    sortKsorted(vec, n, k);
	}
	return 0;
}