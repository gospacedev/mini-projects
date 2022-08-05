// Initialize random
Random rnd = new Random();

// Asking the user how many flips
Console.WriteLine("Enter how many flips: ");

// Reading line and converting it to 32-bit integer
int length = Convert.ToInt32(Console.ReadLine());

for (int i = 0; i < length; i++)// Generating a random number until the amount of flips has been met
{
    int coin  = rnd.Next(0, 2);// Initialize the coin to return a random integer

    if (coin == 0)// If the coin generates 0, the coin flipped heads
    {
        Console.WriteLine("The coin flipped heads.");
    }
    else if (coin == 1)// And if  the coin generates 1, the coin flipped tails
    {
        Console.WriteLine("The coin flipped tails.");
    }   
}
