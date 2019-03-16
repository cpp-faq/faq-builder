# Les modificateur const et volatile

### Sommaire

**En cours d'écriture**

## Quelle est la différence entre const int et int const ?

 Le modificateur ```const``` qualifie l’élément directement à sa gauche et s’il n’y a rien à gauche, il s’applique à l’élément immédiatement à droite. Les formes ```const int``` et ```int const``` sont donc strictement équivalents et désignent toutes les deux un entier dont la valeur ne peut être modifiée.

Le choix de positionnement du mot clef ```const``` est une question de préférence pour les types simples et les référence (```const int&``` et ```int const&``` sont équivalents également), l'important étant de rester cohérent et de garder une notation homogène. Pour les pointeurs le positionnement a de l'importance :

```cpp
int* p; // Déclaration d'un pointeur sur un entier.  

const int* cp1; // Déclaration d'un pointeur sur un entier constant.
int const* cp2; // Déclaration d'un pointeur sur un entier constant.

int* const pc{ &i }; // Déclaration d'un pointeur constant sur un entier.

const int* const cpc1{ cp1 }; // Déclaration d'un pointeur constant sur un entier constant.
int const* const cpc2{ cp1 }; // Déclaration d'un pointeur constant sur un entier constant.
}
```

La position du modificateur ```const``` a donc une signification en ce qui concerne les pointeurs. La chose peut d’ailleurs rapidement se compliquer :
```cpp
int* const* const** const* const* ptr; // Cas extrème.
```

## Quelle est la différence entre const et constexpr ?

Le mot clef ```const``` permet de représenter la constance (l’immuabilité) d’un objet. Il est utilisé comme modificateur de type ou pour spécifier des fonctions membres constantes.

Le mot clef ```constexpr``` permet de déclarer des expressions constantes (constant expression), c’est-à-dire des expressions pouvant être évaluées à la compilation plutôt qu’à l’exécution. Une expression constante est implicitement const.

## Que signifie le mot clef mutable sur une variable membre ?

Le mot-clef ```mutable```, lorsqu'il est appliqué à une variable membre, indique d'un membre d'une classe peut-être modifier même si l'objet est constant :

```cpp
struct A
{
    int i;
    mutable int j;

    A() noexcept = default;

    void foo(int k) const
    {
        i = k; // GCC 8.1 : error: assignment of member 'A::i' in read-only object
        j = k; // Ok.
    }

};

int main()
{
    const A a = {1, 2};
    a.i = 0; // GCC 8.1 : error: assignment of member 'A::i' in read-only object
    a.j = 1; // Ok.      
}
```

Dans cet exemple, la modification du membre ```j``` est possible dans un contexte constant grâce à ```mutable```.

```mutable``` est pratique sur les membres d'une classes qui sont des détails d'implémentations et dont la modification ne modifie pas l'état observable de la classe. Même si l'objet a été modifié d'un point de vue de la mémoire, son état "logique" reste constant. C'est notamment très pratique lorsqu'on implémente un cache : la modification du cache relève de l'optimisation de l'implémentation et n'influence pas l'état logique de l'objet.

#### Liens et compléments
 - **[EN]** [cppreference.com | cv-qualifiers](http://en.cppreference.com/w/cpp/language/cv)

## Que signifie le mot clef mutable sur une lambda ?

Par défaut, les captures par valeur d'une lambda ne sont pas modifiables, ```mutable``` peut-être spécifié (depuis **C++11**) pour permettre la modification de ces variables :

```cpp
int i = 0;

auto l = [i]() { i = 22; }; // GCC 8.1 : error: assignment of read-only variable 'i'
auto ll = [i]() mutable { i = 22; }; // Ok.
```

#### Liens et compléments
 - **[EN]** [cppreference.com | Lambda expressions](https://en.cppreference.com/w/cpp/language/lambda)

## A quoi sert le mot-clef volatile ?

Le mot-clef ```volatile``` fait partie des **cv-qualifiers** au même titre que ```const```. Il s'applique aux variable, aux variables membres ou aux paramètres d'une fonction.

Les membres d'un objet ```volatile``` seront également considérés comme ```volatile```, qu'ils soient mutable ou non.Considérez l'exemple suivant :

Une variable ```volatile``` indique au compilateur qu'il ne possède pas le contrôle exclusif sur celle-ci et que par conséquent il ne devrait considérer que toutes opérations sur la variable ```volatile``` comme ayant un effet de bord.

```cpp
int main()
{
    volatile int a{ 0 };
    a = 0;
}
```

Dans le cas précédent, la plupart des compilateurs auraient tout bonnement optimisé le code en supprimant la variable ```a``` qui n'est pas utilisée. D'autant que l'affectation à 0 est inutile.

Avec le qualificateur ```volatile```, le compilateur va considérer qu'il est possible que quelque-chose survienne entre la déclaration et l'affectation de la variable et donc qu'il ne doit pas l'optimiser.

Le qualificateur volatile n'apporte aucune garantie en matière d’accès concurrent et ne devrait pas être utilisé en ce sens.

L'opérateur de conversion ```const_cast``` permet également d'enlever la qualification volatile d'une variable, le passage d'un objet non-volatile à un objet volatile est quant à elle implicite, comme pour la constance.

#### Liens et compléments
 - **[EN]** [cppreference.com | cv-qualifiers](http://en.cppreference.com/w/cpp/language/cv)
 - [A quoi sert l'opérateur de conversion const_cast ?](https://github.com/cpp-faq/cpp-faq/tree/develop/faq/fr-FR/.faq/404.md)

 ## Que signifie le mot clef restrict ?

 Le mot clef ```restrict``` n'est **pas** un mot-clef en C++ et n'a aucune signification particulière, c'est un mot-clef spécifique au *langage C*.

 En C, ```restrict``` est un qualificateur de type complémentaire à ```const``` et ```volatile``` qui est réservé aux pointeurs sur objets. Il permet d'indiquer au compilateur qu'il n'y a pas de risque de **pointer aliasing** avec un pointeur, c'est-à-dire qu'il peut sans crainte partir du principe que cette adresse mémoire ne chevauche pas avec une autre, et par conséquent lui permettre d'effectuer des optimisations supplémentaire. Si deux plages mémoire se chevauchent, il s'agit d'un *comportement indéterminé*.

 Ce mot-clef n'existe pas en C++ et c'est d'ailleurs l'une des rares incompatibilités entre le C et le C++. Cela étant, certains compilateur proposent des extensions telles que ```__restrict__``` sur GCC et chez le compilateur d'IBM ou ```__restrict``` sur MSVC.

 #### Liens et compléments
  - **[EN]** [cppreference.com | restrict type qualifier](https://en.cppreference.com/w/c/language/restrict)
  - **[EN]** [wikipedia.org | Restrict](https://fr.wikipedia.org/wiki/Restrict)
  - [Un programme C est-il valide en C++ ?](https://github.com/cpp-faq/cpp-faq/tree/develop/faq/fr-FR/.faq/404.md)
