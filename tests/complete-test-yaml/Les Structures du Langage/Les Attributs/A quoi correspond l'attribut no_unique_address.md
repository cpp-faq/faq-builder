---
display-name: A quoi correspond l'attribut [[no_unique_address]] ?
end-links:
- descr: "cppreference.com – C++ attribute: no_unique_address"
  url: https://en.cppreference.com/w/cpp/language/attributes/no_unique_address
  options:
    prefix: EN
- descr: open-std.org | p0840r2 'Language support for empty objects'
  url: http://open-std.org/JTC1/SC22/WG21/docs/papers/2018/p0840r2.html
  options:
    prefix: EN    
---
L’attribut ```[[no_unique_address]]```, ajouté avec **C++20**, est destiné à indiquer qu’une variable membre n’a pas besoin d’avoir une adresse séparée des autres membres de la classes. Cela revient à autoriser un équivalent de l’[Empty Base Optimisation](https://github.com/cpp-faq/cpp-faq/tree/develop/faq/fr-FR/.faq/404.md) pour un membre d’une classe.

```
struct Empty {}; // empty class

struct X {
    int i;
    Empty e;  
};

struct Y {
    int i;
    [[no_unique_address]] Empty e;
};

int main()
{
    // the size of any object of empty class type is at least 1
    static_assert(sizeof(Empty) >= 1);

    // at least one more byte is needed to give e a unique address
    static_assert(sizeof(X) >= sizeof(int) + 1);

    // empty member optimized out
    static_assert(sizeof(Y) == sizeof(int));
}
```

Dans cet exemple tiré de *cppreference*, ```Empty``` est une **empty class**, c’est à dire qu’elle ne contient aucune donnée membre. Puisque le C++ impose que même une classe vide fasse une taille d’au moins 1 byte, la première assertion est donc vraie.

La classe ```X``` contient une donnée membre de type ```Empty```, la classe ```X``` a donc une taille au moins égale à la taille d’un ```int``` plus un byte, soit ```sizeof(int) + 1``` comme le confirme la seconde assertion.

La classe ```Y``` contient elle aussi une donnée membre de type ```Empty```, mais accompagnée de l’attribut ```[[no_unique_address]]```. Dans ce cas, le compilateur est autorisé à optimiser ce membre de manière à ne lui faire occuper aucune place. Par conséquent, la taille de ```Y``` est égale à la taille d’un ```int```, le membre ```e``` n’occupant pas d’espace en mémoire.
