---
display-name: A quoi correspond l'attribut [[noreturn]] ?
end-links:
- descr: "cppreference.com – C++ attribute: noreturn"
  url: https://en.cppreference.com/w/cpp/language/attributes
  options:
    prefix: EN
- descr: open-std.org | n2761 'General Attributes for C++'
  url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2761.pdf
  options:
    prefix: EN
options:
  since: C++11
---
L'attribut ```[[noreturn]]```, introduit avec **C++11**, permet de signaler qu'une fonction ne retourne pas. Il peut s'agir d'une fonction qui lève une exception dans tous les cas ou d'un appel à une fonction qui termine le programme (par exemple ```std::terminate```) ou qui change le contexte d'exécution (```std::longjmp```).

Si la fonction peut effectivement se retourner, il s'agit d'un **undefined behavior**.

```cpp
std::jmp_buf jmp_buf;
[[noreturn]] foo(int a) {
    /* ... */
    std::longjmp(jmp_buf, a);
}

[[noreturn]] bar() { throw false; }
```

La liste des fonctions standards marquées ```[[noreturn]]``` sont listées sur [cette page](https://en.cppreference.com/w/cpp/language/attributes/noreturn).
