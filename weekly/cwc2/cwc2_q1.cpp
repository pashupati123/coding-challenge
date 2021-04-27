#include <bits/stdc++.h>
using namespace std;
string ltrim(const string &);
string rtrim(const string &);
vector<string> split(const string &);

int numPaths(vector<vector<int>> warehouse) {
int n=warehouse.size(), m=warehouse[0].size();
    for(int i=n-2;i>=0;i--)
        if(warehouse[i][m-1] == 1 && warehouse[i+1][m-1] == 1)
            warehouse[i][m-1] = 1;
        else
            warehouse[i][m-1] = 0;
    for(int j=m-2;j>=0;j--)
        if(warehouse[n-1][j] == 1 && warehouse[n-1][j+1] == 1)
            warehouse[n-1][j] = 1;
        else
            warehouse[n-1][j] = 0;
    for(int i=n-2;i>=0;i--) {
        for(int j=m-2;j>=0;j--) {
            if(warehouse[i][j]==1)
                warehouse[i][j]=(warehouse[i+1][j] + warehouse[i][j+1])%(1000000000 +7);
        }
    }
return warehouse[0][0];
}
int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));
    string warehouse_rows_temp;
    getline(cin, warehouse_rows_temp);
    int warehouse_rows = stoi(ltrim(rtrim(warehouse_rows_temp)));
    string warehouse_columns_temp;
    getline(cin, warehouse_columns_temp);
    int warehouse_columns = stoi(ltrim(rtrim(warehouse_columns_temp)));
    vector<vector<int>> warehouse(warehouse_rows);
    for (int i = 0; i < warehouse_rows; i++) {
        warehouse[i].resize(warehouse_columns);
        string warehouse_row_temp_temp;
        getline(cin, warehouse_row_temp_temp);
        vector<string> warehouse_row_temp = split(rtrim(warehouse_row_temp_temp));
        for (int j = 0; j < warehouse_columns; j++) {
            int warehouse_row_item = stoi(warehouse_row_temp[j]);
            warehouse[i][j] = warehouse_row_item;
        }
    }
    int result = numPaths(warehouse);
    fout << result << "\n";
    fout.close();
    return 0;
}
string ltrim(const string &str) {
    string s(str);
    s.erase(
        s.begin(),
        find_if(s.begin(), s.end(), not1(ptr_fun<int, int>(isspace)))
    );
    return s;
}
string rtrim(const string &str) {
    string s(str);
    s.erase(
        find_if(s.rbegin(), s.rend(), not1(ptr_fun<int, int>(isspace))).base(),
        s.end()
    );
    return s;
}
vector<string> split(const string &str) {
    vector<string> tokens;
    string::size_type start = 0;
    string::size_type end = 0;
    while ((end = str.find(" ", start)) != string::npos) {
        tokens.push_back(str.substr(start, end - start));
        start = end + 1;
    }
    tokens.push_back(str.substr(start));
    return tokens;
}