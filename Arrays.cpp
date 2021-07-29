/*-------------
Reading Legacy C++

DEMO: C-Style Arrarys

CREATED BY: Kate Gregory
USED BY: JarekQ Aloisio

Purpose: To study the C-Style Arrays 
-------------*/

int main()
{
    const int howmanynumbers = 4;
    int numbers[howmanynumbers];
    numbers[0] = 8;
    numbers[1] = 7;
    numbers[2] = 6;
    numbers[3] = 5;

    for(int i = 0; i < howmanynumbers; i++)
    {
        numbers[i] += 1;
    }

    *(numbers + 1) = 1;

    double morenumbers[] = {1.1,2.2,3.3,4.4,0};

    for(double* p = morenumbers; *p != 0; p++)
    {
        *p += 1.0;
    }

    int extent = numbers[0] - numbers[3]; // any on the fly calculation

    int* dynamicnumbers = new int[extent];

    dynamicnumbers[0] = 4;
    dynamicnumbers[1] = 3;
    dynamicnumbers[2] = 2;

    *(dynamicnumbers + 3) = 1;

    int localsize = sizeof(numbers) / sizeof(numbers[0]);

    int freestoresize = sizeof(dynamicnumbers) / sizeof(dynamicnumbers[0]);

    delete[] dynamicnumbers;

    return 0;
}