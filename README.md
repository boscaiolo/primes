# Primes Multiplication Table

A golang application to calculate the multiplication table of the first N primes

##Run
    primes.exe
    
##Run Tests    
    go test
    
###Limitations
    Using N*N as limit for generating primes. I have tried to use the formula:
          
          Pn / log Pn ~ N
          
     Where Pn is the Nth prime but could not manage to make it work due to float/int conversion loss
     
     Wanted to use files for the test data but also did not have more time
     
     The app hangs with input 20000 with a deadlock wich might be the fmt going out of mem need to investigate  
    
    