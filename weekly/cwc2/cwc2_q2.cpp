#include <iostream>
#include <bits/stdc++.h>
using namespace std;
bool isChar(char ch) {
    return ch >= 'a' && ch <= 'z';
}
bool isDigit(char ch) {
    return ch >= '0' && ch <= '9';
}
long calc(vector<int> a, vector<int> b) {
    long res = 0;
    for(int i=0; i<a.size(); i++) {
        res += a[i]*b[0];
    }
    return res;
}
int main() {
    string str = "w:/a\\bc::/12\\xyz";
    int size = str.length();
    int mode = 0;
    /*
       0 - starting position
       1 - find char {digit, :, /}
       2 - find digit, : /
       3 - find char after /
       4 - / find char \
       5 - / find digit \
       6 - \ char
    */
    vector<int> count, count1, count2;
    long ans = 0;
    for(int i=0;i<size;i++) {
        char ch = str[i];
        if(mode == 0) {
            if(isChar(ch)) {
                mode = 1;
                count.push_back(1);
            }
        } else if(mode == 1) {
            if(isChar(ch)) {
                count.back()++;
            } else if(isDigit(ch) || ch == ':') {
                mode = 2;
            } else if(ch == '/') {
                mode = 3;
            } else if(ch == '\\') {
                mode = 0;
                count.clear();
            }
        } else if(mode == 2) {
            if(isChar(ch)) {
                mode = 1;
                count.push_back(1);
            } else if(ch == '/') {
                mode = 3;
            } else if(ch == '\\') {
                mode = 0;
                count.clear();
            }
        } else if(mode == 3) {
            if(isChar(ch)) {
                mode = 4;
                count1.push_back(1);
            } else if(ch == ':') {
                mode = 0;
                count.clear();
                count1.clear();
            } else if(ch == '/') {
                mode = 0;
                count.clear();
            } else if(ch == '\\') {
                if(str[i-1] == '/') {
                    mode = 0;
                    count.clear();
                } else {
                    mode = 6;
                    count2.push_back(0);
                }
            }
        } else if(mode == 4) {
            if(isChar(ch)) {
                count1.back()++;
            } else if(isDigit(ch)) {
                mode = 5;
            } else if(ch == ':') {
                mode = 2;
                count = count1;
                count1.clear();
            } else if(ch == '/') {
                mode = 3;
                count = count1;
                count1.clear();
            } else if(ch == '\\') {
                mode = 6;
                count2.push_back(0);
            }
        } else if(mode == 5) {
            if(isChar(ch)) {
                count1.push_back(1);
                mode = 4;
            } else if(ch == ':') {
                mode = 2;
                count = count1;
                count1.clear();
            } else if(ch == '/') {
                mode = 3;
                count = count1;
                count1.clear();
            } else if(ch == '\\') {
                mode = 6;
                count2.push_back(0);
            }
        } else if(mode == 6) {
            if(isChar(ch)) {
                count2.back()++;
            } else if(isDigit(ch) || ch == ':') {
                if(count2.back() == 0) {
                    mode = 0;
                    count.clear();
                    count1.clear();
                    count2.clear();
                } else {
                    mode = 2;
                    ans += calc(count, count2);
                    count = count2;
                    count1.clear();
                    count2.clear();
                }
            } else if(ch == '/') {
                if(count2.back() == 0) {
                    mode = 0;
                    count.clear();
                    count1.clear();
                    count2.clear();
                } else {
                    mode = 3;
                    ans += calc(count, count2);
                    count = count2;
                    count1.clear();
                    count2.clear();
                }
            } else if(ch == '\\') {
                mode = 0;
                ans += calc(count, count2);
                count.clear();
                count1.clear();
                count2.clear();
            }
        }
    }
    if (mode == 6) {
        ans += calc(count, count2);
    }
	return 0;
}