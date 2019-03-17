# Fonctionnement général

**En cours de rédaction.**

### Sommaire

 - [Quel est le minimum d'un programme C++ ?](#Quel%20est%20le%20minimum%20d%27un%20programme%20C++).
 - [Quel est le minimum d'un programme C++ ?](#A%20quoi%20ressemble%20le%20programme%20%27Hello%20World%27%20en%20C++).
 - [Quels sont les différents paradigmes supportés par le C++ ?](#Quels%20sont%20les%20diff%C3%A9rents%20paradigmes%20support%C3%A9s%20par%20le%20C++).
 - [Quels sont les concepts fondamentaux du C++ ?](#Quels%20sont%20les%20concepts%20fondamentaux%20du%20C++).

## Quel est le minimum d'un programme C++ ?

Un programme en C++ doit être à minima composé de la fonction main :

```cpp
int main()
{    
    return 0;
}
```

Ici, l'instruction `return` est même optionnelle. Par convention, retourner `0` signifie que le programme s'est exécuté correctement.

#### Liens et compléments
 - [](q://Quels sont les prototypes autorisés pour la fonction main).

## Quel est le minimum d'un programme C++ ?

Plusieurs versions du programme **Hello World!** sont généralement proposées pour le C++. En voici une :

```
#include <iostream>

int main()
{
     std::cout << "Hello World !\n";
}
```

Généralement les variations proposées tournent autour de l'utilisation de la directive `using namespace` ([Pourquoi devrait-on éviter l'utilisation de la directive using namespace std ?](https://github.com/cpp-faq/cpp-faq/tree/develop/faq/fr-FR/.faq/404.md)) et de la présence ou non du retour chariot ('\n'), de `std::endl`, et des arguments de la fonction `main`.

## Quels sont les différents paradigmes supportés par le C++ ?

**En cours d'écriture**

## Quels sont les concepts fondamentaux du C++ ?

**En cours d'écriture**
