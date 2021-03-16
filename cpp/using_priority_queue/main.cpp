#include <queue>
#include <vector>
#include <functional>
#include <iostream>

int main()
{
    std::priority_queue<int,  std::vector<int>, std::greater<int>> q;
    for (auto n : {5, 3, 4, 2, 1}) {
        q.push(n);
    }

    while (q.size()) {
        std::cout << q.top() << " ";
        q.pop();
    }

    std::cout << std::endl;

    return 0;
}