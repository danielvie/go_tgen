#include <windows.h>
#include <iostream>

typedef int (*MathFunc)(int, int);

int main() {
    HMODULE hDll = LoadLibraryW(L"zig-out/bin/mathlib.dll");
    if (!hDll) {
        std::cerr << "Failed to load DLL!" << std::endl;
        return 1;
    }

    MathFunc sum = (MathFunc)GetProcAddress(hDll, "sum");
    MathFunc sub = (MathFunc)GetProcAddress(hDll, "sub");

    if (!sum || !sub) {
        std::cerr << "Failed to get function address!" << std::endl;
        FreeLibrary(hDll);
        return 1;
    }

    std::cout << "Sum: " << sum(5, 3) << std::endl;
    std::cout << "Sub: " << sub(5, 3) << std::endl;

    FreeLibrary(hDll);
    return 0;
}
