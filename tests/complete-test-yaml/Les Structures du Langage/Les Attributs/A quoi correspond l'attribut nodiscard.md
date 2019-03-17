---
display-name: A quoi correspond l'attribut [[nodiscard]] ?
end-links:
- descr: "cppreference.com – C++ attribute: fallthrough"
  url: https://en.cppreference.com/w/cpp/language/attributes/nodiscard
  options:
    prefix: EN
options:
  since: C++17
---
L’attribut ```[[nodiscard]]```, ajouté avec **C++17**, permet de refuser le droit du programmeur d’ignorer le retour d’une fonction.

Assez simplement, une fonction peut être marquée avec ```[[nodiscard]]``` ce qui indique au compilateur que le retour ne devrait pas être ignoré :

```cpp
[[nodiscard]] error_code f();

int main() {
   f(); // GCC-7.2 : warning: ignoring return value of 'error_code f()',
        // declared with attribute nodiscard [-Wunused-result]
}
```

L’attribut ```[[nodiscard]]``` peut également être appliqué à un type. Dans ce cas, toutes les fonctions retournant un objet de ce type sont implicitement marquées ```[[nodiscard]]``` :

```struct [[nodiscard]] error_code {};

error_code f();
error_code g();

int main() {
   f(); // warning [[nodiscard]].
   g(); // warning [[nodiscard]].
}
```
