# Les conversions

**En cours d'écriture**

#### Sommaire

**En cours d'écriture**

## Quels sont les différents types de conversion disponibles en C++ ?

Le C++ supporte quatre classes majeures de conversion, qui sont chacune explicitées dans cette section :
 - Les conversions implicites standards, qui correspondent à toute les conversions effectuées automatiquement par défaut par le compilateur, par exemple la conversion d'un `int` en `double`.
 - Les conversions explicites de type C.
 - Les conversions explicites C++ :
  - `static_cast`.
  - `const_cast`.
  - `dynamic_cast`.
  - `reinterpret_cast`.
 - Les conversions définies par l'utilisateur (*user-defined*), qui peuvent être implicites ou explicites.

Certaines fonctions de la bibliothèque standard, telles que `std::bit_cast` ou `std::any_cast`, peuvent également être rangées dans la catégorie des conversions explicites. De la même manière, les bibliothèque peuvent fournir leurs propres moyen de conversion pour les types non standards.

## Qu'est qu'une conversion implicite ?

Au sens large, une conversion implicite correspond à toutes les conversions qui se sont pas explicitement demandée par le programmeur. Cette définition inclut donc les conversions implicites effectuées par le langage (*standard conversion*) ainsi que les conversions implicites définies par l'utilisateur (*user-defined conversion*).

Le terme *conversion implicite* est souvent utilisé pour se référer uniquement aux conversions implicites standard. Ci-suit une liste non-exhaustive des conversions standards :
  - les conversions numérique : promotion de type (entière ou flottante), conversion entière, flottante, entier en flottant, pointeurs et pointeurs sur membre...
  - les conversions booléenne : depuis un entier, un flottant, un pointeur...
  - les conversions contextuelles : en booléen (dans un `if`, `switch`, une boucle...), avec `noexcept`, un prédicat de contrat...
  - les conversions de qualification (`const` et `volatile`).
  - les conversions de valeurs : *lvalue* -> *rvalue*, tableau en pointeur, fonction en pointeur...

#### Liens et compléments
 - **[EN]** [cppreference.com – C++ implicit conversion](https://en.cppreference.com/w/cpp/language/implicit_conversion)

## Qu'est-ce qu'une conversion explicite ?

Une conversion explicite correspond à une conversion qui est clairement demandée par le programmeur. Les opérateurs de conversion C++ standard (`static_cast`, `const_cast`, `dynamic_cast` et `reinterpret_cast`) ainsi que l'opérateur de conversion de type C font partie des conversions explicite C++. On peut également qualifier comme conversion explicites : les litéraux utilisateur de *C++14* et au sens large toute fonction permettant un transtypage (`std::any_cast`, `std::to_string` ou encore `my_lib::integer_to_mpz_t`).

## A quoi sert l'opérateur de conversion reinterpret_cast ?

## Qu'est-ce que la 'strict aliasing rule' ?

## Comment puis-je réinterpréter une séquence de byte vers un autre type ?
