#include <iostream>
#include <vector>
#include <string>
using namespace std;

typedef unsigned long long ull;

const ull factor = 1000000009;

ull getHash(const vector<ull>& hashes, const vector<ull>& pows, int left, int right) {
    return hashes[right + 1] - pows[right - left + 1] * hashes[left];
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);

    string s;
    cin >> s;

    int m;
    cin >> m;

    vector<ull> hashes(s.length() + 1);
    vector<ull> pows(s.length() + 1);

    hashes[0] = 0;
    pows[0] = 1;

    for (int i = 1; i <= s.length(); i++) {
        pows[i] = pows[i - 1] * factor;
    }

    for (int i = 0; i < s.length(); i++) {
        hashes[i + 1] = hashes[i] * factor + s[i];
    }

    for (int i = 0; i < m; i++) {
        int a, b, c, d;
        cin >> a >> b >> c >> d;

        if (getHash(hashes, pows, a - 1, b - 1) == getHash(hashes, pows, c - 1, d - 1)) {
            cout << "Yes\n";
        } else {
            cout << "No\n";
        }
    }

    return 0;
}
