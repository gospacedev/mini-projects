#include<iostream>

using namespace std;

/*
The Collatz conjecture in mathematics asks whether repeating two simple 
arithmetic operations will eventually transform every positive integer 
into one. If n is even divide by 2 and when odd times n by 3 and add 1.
*/
int main()
{
    int n;
    cout<<"Enter a number:\n";
    cin >> n;
    int counter=0;

    for (;;)
    {
        if (n % 2 == 0)//even
        {
            n/=2;
        }
        else if (n == 1)//solved
        {
            cout << "The problem has been solved!" << endl;
            break;
        }
        else//odd
        {
            n=(n*3)+1;
        }
        counter++;
        cout<<"n: "<<n<<endl;
        cout<<"Step: "<<counter<<endl;
    }
}
