# Les 1 les
## toetsing

- Eerst wordt er 2 weken over low level dan 4 api's
- 1 punt per goed huis werk met 1 eind tentamen
- punten van huiswerk komt bij tentamen

## memory

### hoe werkt een programma
- je hebt 4 onderdelen in memory van het programma
    - stack
    - heap
    - data
    - code

### opdrachtje

    - Kijk welke varabel welk memory type bij welke variabele hoort?

    #### Code
    ```c
    #include <stdio.h>

    int A = 8; //DATA omdat dit een global waarde is

    int main() {
        int i = 10; //STACK Omdat het alleen aanwezig is tijdens de functie
        
        int sum = i+A; //STACK Zelfde reden als hiervoor

        printf("%d\n", sum);

        //program executed correctly
        return 0;
    }
    ```
    #### Uitleg

        - Stack: waar tijdelijke data staat (aanroepen van functie bijvoorbeeld)
        - Heap: voor dynamisch data (Bijvoorbeeld nieuwe klasse)
        - Data: voor static en global vars
        - Code: Hier komt je programma in die gerunt wordt


## Tijdcomplexiteit

- Grootste afhankelijkheid is altijd input
- "Big O" is vanuit het slechtste scenario berekend en staat voor de tijdscomplexiteit
    - Elkele operaties staat voor: O(1)
    - Evalueren van een conditional: O(1)
    - Itereren over een lijstje met lengte n: O(n)

### Opdracht
```c
int main() {
    int x = 10; //O(1)
    int y = 10; //O(1)

    int sum = x+y; //O(1)

    if (10 > 5) { //O(1)
        printf("True!\n"); //ook O(1) met alles erna ook
    } else {
        printf('False!\n");
    }

    int[] array = [];

    for (int i = 0; i < sizeof(array); i++) //O(n)
        //als binnenste forloop dan is die ook O(n) en dan is de bovenliggende forloop O(n^2)
        printf("%s", array[i])
    }

    return 0;
```

- Kortom forloop in forloop is O(n^2)
- En maken vars niet meer uit


## Huiswerk

- Ga naar learn-c.com
- werk t/m hoofdstuk "static"

