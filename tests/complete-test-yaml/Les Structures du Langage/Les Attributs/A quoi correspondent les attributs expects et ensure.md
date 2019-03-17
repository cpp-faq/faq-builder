---
display-name: A quoi correspondent les attributs [[expects]] et [[ensures]] ?
end-links:
- descr: Comment définir une pré/postcondition en C++ ?
  url: faq://Comment définir une pré/postcondition en C++ ?
- descr: "cppreference.com – C++ attribute: expects, ensures, assert"
  url: https://en.cppreference.com/w/cpp/language/attributes/contract
  options:
    prefix: EN
- descr: open-std.org | p0840r2 'Support for contract based programming in C++'
  url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2018/p0542r5.html
  options:
    prefix: EN
---
```[[expects]]``` et ```[[ensure]]```, ajouté avec **C++20**, sont les **contracts condition**. Ils permettent d'exprimer respectivement la précondition et la postcondition.

Ils s'appliquent à une fonction de la manière suivante :

```cpp
void foo(int x) [[expects: x > 0]] // La précondition est violée si le paramètre x est inférieur ou égal à 0.

void foo(const std::string& s) [[ensures: !s.empty]]; // La postcondition est violée si s est vide après l'appel de la fonction.
```

Les deux attributs acceptent une expression booléenne. Les arguments de la fonction et toutes les variables accessibles peuvent être utilisées. Dans le cas de ```[[ensures]]```, il est possible d'introduire l'identifiant qui sera utilisé comme retour de fonction de manière à l'utiliser dans la post condition :

```
std::size_t read()
  [[ensures r: r != 0]]
{
    /* ... */
    std::size_t r = read(something);
    /* ... */
    return r;
}
```

Tout comme ```[[assert]]``` ces attributs peuvent avoir un **build-level** [Qu'est-ce que le build level ?](https://github.com/cpp-faq/cpp-faq/tree/develop/faq/fr-FR/.faq/404.md). Les indication fournies par ces attributs donnent des informations à l'optimiseurs et aux outils statiques, ainsi qu'au développeurs qui lisent le code.
