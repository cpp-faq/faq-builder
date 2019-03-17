---
display-name: A quoi correspond l'attribut [[assert]] ?
end-links:
- descr: Comment faire une assertion en C++ ?
  url: faq://Comment faire une assertion en C++
- descr: "cppreference.com – C++ attribute: expects, ensures, assert"
  url: https://en.cppreference.com/w/cpp/language/attributes/contract
  options:
    prefix: EN
- descr: open-std.org | p0840r2 "Support for contract based programming in C++"
  url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2018/p0542r5.html
  options:
    prefix: EN
---
L'attribut ```[[assert]]``` fait partie des ajouts de **C++20** pour le support de la programmation par contrat. Cet attribut permet de déclarer une assertion, il peut être vu comme une version moderne de la macro **C** ```assert()```.

Cet attribut est assez particulier puisqu'il s'utilise à la manière d'une instruction et est suivis par un point-virgule :

```cpp
void foo(int i = 0) {
    /* ... */
    [[assert: x >= 0]];
    /* ... */
    [[assert: !something()]];
    /* ... */
}
```

```[[assert]]``` permet de vérifier la validité d'un prédicat à un point donné d'une fonction. Comme les autres attributs liés aux contrats, il est possible de spécifier un niveau de contrat (*contract-level*), ```default``` étant implicite.

Par rapport à l'ancien ```assert``` hérité du **C**, cet attribut apporte quelques avantages non négligeables. Déjà, il ne s'agit pas d'une macro et il permet d'éviter les risques associés (```[[assert: c == std::complex<float>{0, 0}]]``` compile contrairement à ```assert(c == std::complex<float>{0,0})```). Ensuite, les *contract-level* ainsi que la possibilité de régler le niveau de vérification voulu (associé à la possibilité de fournir un gestionnaire de violation de contrat personnalisé) :

```cpp
[[assert audit : x != nullptr && reachable(x)]]
[[assert axiom : sorted(database)]]
```

Enfin, ```[[assert]]``` permet à l'optimiseur et aux outils d'analyse statique d'avoir des informations précises sur l'assertion ce qui facilite l'analyse et peut ouvrir la voie à des optimisations supplémentaires.
