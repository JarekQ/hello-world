/*-------------
Reading Legacy C++

DEMO: C-Style Strings

CREATED BY: Kate Gregory
USED BY: JarekQ Aloisio

Purpose: To study the C-Style Arrays 
-------------*/

#include <cstring>
#include <string>

int main()
{
    char const* s1 = "Hello";
    int i = strlen(s1);

    char* s2 = new char[strlen(s1) + strlen(" Kate")];
    strcpy(s2,s1);
    strcat(s2," Kate");
    // Borrowed from someone else's memory for the Null character (end)
    char c = s2[10];

    // Improper String Initialization
    char hello[] = {' ','w','o','r','l','d'};
    size_t j = strlen(hello);

    // Proper String Initialization
    char better[] = "better";
    int k = strlen(better);

    // N
    // Mix and Match New and Old
    std::string s(s1);
    s += " modern C++";
    char* s3 = new char[s.length() + 1];
    // s3 = s;
    // s3 = s.c_str();

    strcpy(s3, s.c_str());

    delete[] s2;
    return 0;
}