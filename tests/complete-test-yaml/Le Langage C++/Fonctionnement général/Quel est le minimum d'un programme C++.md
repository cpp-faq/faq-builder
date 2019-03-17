---
display-name: Quel est le minimum d'un programme C++ ?
end-links:
- url: q://Quels sont les prototypes autorisés pour la fonction main
---
Un programme en C++ doit être à minima composé de la fonction main :

```cpp
int main()
{    
    return 0;
}
```

Ici, l'instruction `return` est même optionnelle. Par convention, retourner `0` signifie que le programme s'est exécuté correctement.
