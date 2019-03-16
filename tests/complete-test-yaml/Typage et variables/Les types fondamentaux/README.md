# Les types primitifs

**En cours d'écriture**

#### Sommaire

**En cours d'écriture**

## Qu'est-ce que les types fondamentaux en C++ ?

## Quels sont les types fondamentaux en C++ ?

Les types fondamentaux en **C++** sont les suivants :
 - le type booléen : `bool`.
 - les types entiers : `short int`, `int`, `long int`, `long long int` (**C++11**) et équivalents non-signés ([Quels sont les types entiers fondamentaux ?](404)).
 - les types caractères : `char`, `signed char`, `unsigned char`, `wchar_t`, `char16_t` (**C++11**), `char32_t` (**C++11**) et `char8_t` (**C++20**).
 - les types flottant : `float`, `double` et `long double` (**C++11**).
 - le type `void`.

Le standard propose aussi d'autres types qui ne sont pas considérés comme *fondamentaux*, définis dans l'entête *<cstddef>* :
 - ̀`std::size_t` : type correspondant à une taille, (il s'agit aussi du type de retour des opérateurs `sizeof` et ̀`alignas`).
 - `std::nullptr_t` : il s'agit du type de `nullptr`, le littéral correspondant à la valeur d'un pointeur nul.
 - `std::ptrdiff_t` : un type entier signé obtenu en soustrayant deux pointeurs.
 - `std::byte` : un type entier non signé utilisé pour manipuler des bytes (sans leur attribuer de sémantique particulière).
 - `std::maxalign_t` : correspondant a un type dont l'alignement requis est au moins aussi grand que le maximum de tous les types scalaires (entiers et flottant).
 - On peut également citer les macros `NULL` et `offsetof`.

#### Liens et compléments
  - **[EN]** [cppreference.com - Fundamentals types](https://en.cppreference.com/w/cpp/language/types)
  - **[EN]** [cppreference.com - Type support](https://en.cppreference.com/w/cpp/types)

## Comment connaître la valeur maximale d'un type entier ou flottant (scalaires) ?

L'entête *<limits>* fourni la classe template `std::numeric_limit` permettant d'obtenir la valeur maximale et minimale d'un type scalaire :

```cpp
#include <limits>

// ...

std::cout << std::numeric_limits<int>::max() << '\n';       // sortie possible : 2147483647
std::cout << std::numeric_limits<double>::min() << '\n';    // sortie possible : 2.22507e-308
std::cout << std::numeric_limits<double>::lowest() << '\n'; // sortie possible : -1.79769e+308
```

`min` retourne la plus petite valeur positive possible pour les flottants, `lowest` permet d'obtenir la borne minimale négative. `numeric_limit` permet également de récupérer beaucoup d'autres informations telles que le style d'arrondi (̀`std::numeric_limit<char>::round_style()`) ou le maximum/minimum de l'exposant d'un type flottant (`std::numeric_limit<float>::min_exponent()`).

#### Liens et compléments
  - **[EN]** [cppreference.com - std::numeric_limits](https://en.cppreference.com/w/cpp/types/numeric_limits)

## Quelle est la taille des types entiers en C++ ?

Tout comme le **C**, le standard **C++** ne fixe pas la taille des types entiers. Cependant il existe un certain nombre de garanties proposées par la norme :
 - `1` <= `sizeof(short)` <= `sizeof(int)` <= `sizeof(long)` <= `sizeof(long long)`.
 - une taille minimale en bits est garanti pour les types entiers (et leurs équivalents non-signés) :
  - `short` : au moins 16 bits.
  - `int` : au moins 16 bits.
  - `long` : au moins 32 bits.
  - `long long` : au moins 64 bits.

#### Liens et compléments
 - **[EN]** [cppreference.com - Fundamentals types](https://en.cppreference.com/w/cpp/language/types)

## Comment sont codés les types entiers en C++ ?

Le codage en [complément à deux](https://fr.wikipedia.org/wiki/Compl%C3%A9ment\_%C3%A0\_deux) est imposé par la norme depuis **C++20**.

Avant **C++20** la norme ne spécifiait pas le codage des entiers signés et on pouvait potentiellement rencontrer du codage en *complément à un*, en *complément à deux* ou encore en *magnitude signée (bit de signe)*. Cependant, les machines modernes utilisent toutes le complément à deux.

#### Liens et compléments
 - **[EN]** [open-std.org – p0907r0 - "Signed Integers are Two’s Complement"](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2018/p0907r0.html)
 - **[EN]** [wikipedia.org - Signed number representations](https://en.wikipedia.org/wiki/Signed\_number\_representations#Signed\_magnitude\_representation)
