#include <iostream>
#include <math.h>
using namespace std;

bool isPrime(int n){
    if (n <= 1)
    {
        return 0;
    }
 
    for(int i=2 ; i <= sqrt(n); i++)
    {
        if(n%i == 0)
        {
            return 0;
        }
    }
    return 1;
}

int main(){
    int n;
    cout<<"Enter a number:"<<endl;
    cin >> n;

    if (isPrime(n)== true)
    {
        cout<<"The number is a prime number"<<endl;
    }else if (isPrime(n)==false)
    {
        cout<<"The number is not a prime number"<<endl;
    }
}
