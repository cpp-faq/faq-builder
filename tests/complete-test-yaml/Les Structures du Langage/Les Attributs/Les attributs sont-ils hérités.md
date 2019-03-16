---
display-name: Les attributs sont-ils hérités ?
end-links: []
---
Les attributs ne sont pas hérité. Ci-suit un exemple avec ```[[nodiscard]]``` ([A quoi correspond l'attribut [[nodiscard]] ?](https://github.com/cpp-faq/cpp-faq/tree/master/faq/fr-FR/04%20-%20Les%20structures%20du%20langage/Les%20attributs#a-quoi-correspond-lattribut-nodiscard-)) :

```cpp
struct [[nodiscard]] error_code {};
struct critical_error_code : error_code {};

struct A
{
    virtual [[nodiscard]] bool foo() = 0;
};

struct B : a
{
    bool foo() override;
};

critical_error_code foo();

int main() {
   h(); // pas de warning.
   A{}.foo(); // pas de warning.
}
```
