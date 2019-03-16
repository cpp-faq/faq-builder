# L'inférence de type

**En cours d'écriture**

#### Sommaire

**En cours d'écriture**

## Que signifie auto dans les arguments d’une fonction ?

Depuis **C++14**, il est possible de créer des fonctions anonymes (*lambdas*) polymorphiques/génériques (*generic/polymorphic lambda functions*) à l’aide du mot clef `auto`.

Le compilateur peut inférer le type des arguments passé, et même en déduire le type de retour pour les fonctions lambdas :

```cpp
#include <complex>

using namespace std::literals::complex_literals;

auto add = [](auto a, auto b) { return a + b; };

std::cout << add(1.2, 5) << std::endl; // (1)
std::cout << add(0i, 1i) << std::endl; // (2)
```

Dans le premier appel *(1)*, les arguments déduits sont `double` et `int`, le type de retour est donc `double`. Dans le second cas (2), le compilateur déduit `std::complex<double>` pour les deux arguments et donc le type de retour.

Les *generic lambdas* sont donc l’équivalent d’une fonction template pour les lambdas (dans la fonction `add`, le compilateur lèverait une erreur si les types inférés ne supportaient pas l’addition). Le mot clef `auto` est ici utilisé pour préciser l’inférence de type pour les arguments de la fonction. Puisque le type de retour n’est pas précisé, celui-ci est également déduit à partir de l’expression du `return`.

## Peut-on utiliser auto avec la syntaxe d’initialisation uniforme ?

Depuis **C++14**, il est en effet possible d’utiliser `auto` et l’*uniform initialization syntax* conjointement :

```cpp
auto i{ 0 }; // i est un int.
auto i = 0; // i est un int.
auto il = {0}; // i est du type ‘std::initialization_list<int>’.
```

En **C++11** en revanche, le comportement n’est pas celui attendu :

```cpp
auto i{ 0 }; // i est du type ‘std::initialization_list<int>’.
auto i = 0; // i est un int.
```

## Que signifie auto int i; et pourquoi le type int est-il précisé ?

Avant la norme **C++11**, `auto` était un *storage class specifier* (spécificateur de classe de stockage) au même titre que `extern`, `static`, `register` et `thread_local` (**C++11**). Il permettait alors de spécifier qu’une variable était dans la classe de stockage automatique, c’est-à-dire que la durée de vie de celle-ci est gérée automatiquement par le compilateur.

Cependant, il s’agit par là du comportement par défaut et c’est pourquoi la norme **C++11** a changé la signification de ce mot clef qui n’était (pratiquement) jamais utilisé :

```cpp
auto int = 25; // i est une variable automatique. Invalide depuis C++11.
// Même signification (en C++03):
int i = 25;
```

Le mot clef `auto` était très rarement utilisé avant son changement de signification (tout comme il est très peu utilisé en **C**), cependant vous pouvez encore croiser du vieux code C++ utilisant `auto` avec son ancien sens. Si ce code n’est pas compatible avec **C++11**, un rechercher/remplacer sur l’éditeur ou un petit script permettent le mettre à jour rapidement.
