# Alias de types

**En cours d'écriture**

#### Sommaire

**En cours d'écriture**

## Comment faire un alias sur un type en C++ ?

Le mot clef `using` permet de créer un alias sur un type en C++ (depuis **C++11**) :

```cpp
using uint = unsigned int;
using str = std::string;

template<typename T> using MyVector = std::vector<T, MyAllocator<T>>;

using my_func_t = void(*)(int, int)
```

#### Liens et compléments
 - **[EN]** [cppreference.com – type alias](https://en.cppreference.com/w/cpp/language/type_alias)

## Est-ce qu'un alias est un nouveau type ?

Non. Un alias n'introduit pas un nouveau type. La norme l'indique d'ailleurs clairement :

> A typedef-name can also be introduced by an alias-declaration. The identifier following the using keyword becomes a typedef-name and the optional attribute-specifier-seq following the identifier appertains to that typedef-name. Such a typedef-name has the same semantics as if it were introduced by the typedef specifier. **In particular, it does not define a new type.**
>
> *n4478 - 9.1.3.2*  

Ainsi, un alias n'est qu'une autre façon de nommer un type existant. Le `typeid` d'un type et de ses aliases sont identiques :

```cpp
using ull = unsigned long long;
typedef unsigned long long ULL;

assert(typeid(ull) == typeid(unsigned long long));
assert(typeid(ull) == typeid(ULL));
```

## Comment faire un alias avec typedef ?

La syntaxe de création d'alias de type avec `typedef` est semblable à celle de la déclaration d'une variable :

```cpp
typedef unsigned int uint;
typedef std::string str;

typedef void(*my_func_t)(int, int);
```

En théorie, le mot clef `typdef` peut se trouver n'importe où dans la déclaration : `long long typedef unsigned int llui;`. On préférera cependant le spécifier au début de la déclaration pour des raisons de lisibilité.

Notez également que `typedef` ne permet pas directement la création d'aliases template ([Comment créer un alias sur un template avant C++11 ?](https://github.com/cpp-faq/cpp-faq/tree/develop/faq/fr-FR/.faq/404.md))

#### Liens et compléments
 - **[EN]** [cppreference.com – typedef](https://en.cppreference.com/w/cpp/language/typedef)

## Dois-je préférer typedef ou using pour déclarer un alias sur un type ?

Dans la mesure du possible, préférez l'utilisation de `using` pour déclarer un alias. `using` utilise une syntaxe plus naturelle et supporte nativement les aliases sur des types paramétrés.

Si votre compilateur ne supporte pas **C++11**, utilisez `typedef`.

Notez que `using` ne permet pas de remplacer `typedef` dans tous les cas. La norme fait la différence entre une *typedef-declaration* et une *alias-declaration* (avec `using`). Si les deux permettent de définir un alias sur un type (c'est à dire qu'ils introduisent un *typedef-name), il y a des cas où ils se comportent légèrement différemment.

Par exemple, l'exemple suivant montre l'utilisation d'un seul `typedef` tandis que son équivalent avec `using` nécessite deux directives :

```cpp
typedef struct node_t
{
    int value;
    node_t* prev;
    node_t* next;
} node, *p_node;
```

```cpp
struct node_t
{
    int value;
    node_t* prev;
    node_t* next;    
};

using node = node_t;
using pnode = node_t*;
```

## Comment créer un alias sur un template avant C++11 ?

`typedef` ne peut pas être utilisé directement pour créer un alias sur un type paramétré :

```cpp
template<typename T>
typedef std::vector<T> MyVector; // GCC 8.1 :  error: template declaration of 'typedef'
```

Il est cependant possible d'encapsuler `typedef` dans un type intermédiaire :

```cpp
template<typename T>
struct MyVector {
    typedef std::vector<T> type;
};

MyVector<int>::type vec;
```
