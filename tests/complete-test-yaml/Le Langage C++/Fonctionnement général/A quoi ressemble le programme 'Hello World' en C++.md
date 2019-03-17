---
display-name: Quel est le minimum d'un programme C++ ?
end-links: []
---
Plusieurs versions du programme **Hello World!** sont généralement proposées pour le C++. En voici une :

```
#include <iostream>

int main()
{
     std::cout << "Hello World !\n";
}
```

Généralement les variations proposées tournent autour de l'utilisation de la directive `using namespace` ([Pourquoi devrait-on éviter l'utilisation de la directive using namespace std ?](https://github.com/cpp-faq/cpp-faq/tree/develop/faq/fr-FR/.faq/404.md)) et de la présence ou non du retour chariot ('\n'), de `std::endl`, et des arguments de la fonction `main`.
