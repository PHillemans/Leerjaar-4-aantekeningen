#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void mystrcat(char * destination, char * source, int * pos_point);

int main() {
    char stringOne[] = "Zonder "; 
    char stringTwo[] = "toeters ";
    char stringThree[] = "en ";
    char stringFour[] = "bellen.";

    int pos = 0; // continue from the point written to this var
    int i = 0; // Initialize allocation
    i = strlen(stringOne) + strlen(stringTwo) + strlen(stringThree) + strlen(stringFour); // set length to expected string length
    char* bigString; // create pointer named big string
    bigString =  (char*) malloc (i + 1); // allocate memory to stringlength

    mystrcat(bigString, stringOne, &pos); 
    mystrcat(bigString, stringTwo, &pos); 
    mystrcat(bigString, stringThree, &pos); 
    mystrcat(bigString, stringFour, &pos); 

    printf("%s\n", bigString);

    return 0;
}    

void mystrcat(char * dest, char * src, int * pos_point) {
    dest += *pos_point; // set destination from where string needs to be concatted
    while (* dest++ = * src++) *pos_point += 1; // concat string
}
