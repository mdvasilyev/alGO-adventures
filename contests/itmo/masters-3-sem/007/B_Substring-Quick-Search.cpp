#include <iostream>
#include <vector>
#include <string>
using namespace std;

vector<int> findSubstrings(const string& p, const string& t) {
    vector<int> res;
    string combined = p + " " + t;
    vector<int> prefix(combined.length(), 0);

    for (int i = 1; i < combined.length(); i++) {
        int j = prefix[i - 1];
        while (j > 0 && combined[i] != combined[j]) {
            j = prefix[j - 1];
        }
        if (combined[i] == combined[j]) {
            j++;
        }
        prefix[i] = j;
    }

    for (int i = p.length() + 1; i < combined.length(); i++) {
        if (prefix[i] == p.length()) {
            res.push_back(i - 2 * p.length() + 1);
        }
    }

    return res;
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);

    string p, t;

    cin >> p;
    cin >> t;

    vector<int> res = findSubstrings(p, t);

    cout << res.size() << endl;
    for (int pos : res) {
        cout << pos << " ";
    }
    cout << endl;

    return 0;
}
