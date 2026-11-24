#include <cmath>
#include <cstdlib>
#include <iostream>

int main(int argc, char** argv) {
    if (argc != 5) { std::cerr << "usage: score age income cards carloans\n"; return 2; }
    const double w[4] = {0.02291608894566896, 3.743010242991806e-05, -0.234389214456067, -0.30267949575853875};
    const double b = -2.507835092167375;
    double z = b;
    for (int i = 0; i < 4; ++i) z += w[i] * std::atof(argv[i + 1]);
    std::cout << 1.0 / (1.0 + std::exp(-z)) << "\n";
    return 0;
}
