#include <iostream>
#include <vector>
#include <string>
#include <utility>

class Graph {
    std::vector<std::vector<int>> adj;
    std::vector<bool> visited;
    std::vector<int> entryTime;
    std::vector<int> exitTime;
    int timer;

public:
    Graph(int n) : adj(n), visited(n), entryTime(n), exitTime(n), timer(0) {}

    void add(int u, int v) {
        adj[u].push_back(v);
    }

    void dfs(int v) {
        visited[v] = true;
        timer++;
        entryTime[v] = timer;

        for (int u : adj[v]) {
            if (!visited[u]) {
                dfs(u);
            }
        }

        timer++;
        exitTime[v] = timer;
    }

    bool isPredecessor(int u, int v) {
        return entryTime[u] <= entryTime[v] && exitTime[u] >= exitTime[v];
    }
};

int main() {
    std::ios_base::sync_with_stdio(false);
    std::cin.tie(nullptr);

    int n, m;
    std::cin >> n >> m;

    Graph g(n);

    if (n > 1) {
        for (int i = 0; i < n - 1; i++) {
            int parent;
            std::cin >> parent;
            g.add(parent - 1, i + 1);
        }
    }

    std::vector<std::pair<int, int>> pairs(m);
    for (int i = 0; i < m; i++) {
        int u, w;
        std::cin >> u >> w;
        pairs[i] = std::pair(u - 1, w - 1);
    }

    g.dfs(0);

    for (const auto& p : pairs) {
        std::cout << (g.isPredecessor(p.first, p.second) ? 1 : 0) << '\n';
    }

    return 0;
}
