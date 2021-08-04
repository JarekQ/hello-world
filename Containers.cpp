/*-------------
C++ 17: Beyond the Basics

DEMO: vector

CREATED BY: Kate Gregory
USED BY: JarekQ Aloisio

Purpose: To study Containers, Vector, Resource, etc.
-------------*/

#include "Resource.h"
#include <vector>
#include <stdexcept>
#include <iostream>

using std:vector;
using std:cout;

int main()
{
    vector numbers{0,1,2};
    numbers.push_back(-2);
    numbers[0] = 3;
    int num = numbers[2];

    for(int i : numbers)
    {
        cout << i << '\n';
    }

    // Resume here in this area of code | Watch video from the beginning!!!

    Resource r("local");
    {
        cout << "--------------------" << '\n';
        vector<Resource> resources;
        resources.push_back(r);
        cout << "--------------------" << '\n';
        resources.push_back(Resources("first"));
        cout << "--------------------" << '\n';
    }

    return 0;
}