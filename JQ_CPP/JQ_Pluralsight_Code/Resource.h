/*-------------
C++ 17: Beyond the Basics

DEMO: vector

CREATED BY: Kate Gregory
USED BY: JarekQ Aloisio

Purpose: To study Containers, Vector, Resource, etc.
-------------*/

#pragma once
#include <string>

class Resource
{
    private:
        std::string name;
    public:
        Resource(std::string n);
        Resource(const Resource& r);
        Resource& operator=(const Resource& r);
        ~Resource(void);
        std::string GetName() const {return name;}
};