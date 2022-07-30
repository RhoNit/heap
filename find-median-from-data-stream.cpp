class MedianFinder {
private:
    priority_queue<double> pqmax;
    priority_queue<double, vector<double>, greater<double>> pqmin;
    
public:
    void addNum(int num) {
        if(pqmax.size() == 0) { pqmax.push(num); }
        else {
            if(pqmax.top() > num) { pqmax.push(num); }
            else { pqmin.push(num); }
        }

        //balance the median
        if(pqmax.size() > pqmin.size()+1) {
            pqmin.push(pqmax.top());
            pqmax.pop();
        }
        else if(pqmax.size() < pqmin.size()) {
            pqmax.push(pqmin.top());
            pqmin.pop();
        }
        return;
    }
    
    double findMedian() {
        if(pqmax.size() == pqmin.size()) {
            return (pqmax.top()+pqmin.top())/2;
        }
        return pqmax.top();
    }
};
