# Fonctionnement général

Cette section réponds aux questions concernant le fonctionnement général du langage C++.

### Sommaire


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
 - [Quels sont les prototypes autorisés pour la fonction main ?](https://github.com/cpp-faq/cpp-faq/tree/develop/faq/fr-FR/.faq/404.md)

## A quoi ressemble le programme "Hello World" en C++ ?

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
