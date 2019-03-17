---
display-name: Qu'est-ce qu'un attribut ?
end-links:
- descr: "cppreference.com – C++ attribute: carries_dependency"
  url: https://en.cppreference.com/w/cpp/language/attributes/carries_dependency
  options:
    prefix: EN
- descr: "cppreference.com - std::kill_dependency"
  url: https://en.cppreference.com/w/cpp/atomic/kill_dependency
  options:
    prefix: EN
- descr: open-std.org | n2761 'General Attributes for C++'
  url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2761.pdf
  options:
    prefix: EN
---
L'attribut ```[[carries_dependency]]``` introduit avec **C++11**, permet d'indiquer qu'une chaîne de dépendance se propage à l'intérieur ou à l'extérieur d'une fonction (avec un ```std::memory_order``` **release** ou **consume**) et qu'il est donc possible d'éviter des barrières inutiles.

Si cet attribut est placé devant le paramètre d'une fonction (par exemple ```foo(int* [[carries_dependency]] v)```), il indique que l'initialisation du paramètre a une dépendance dans la conversion *lvalue-vers-rvalue*.

Si il est appliqué à une fonction (comme ```[[carries_dependency]] int* foo()```) il indique que la valeur de retour à une dépendance par rapport à l'appelant.
