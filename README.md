# GOBOX my golang sandbox

![](https://github.com/haagor/gobox/blob/main/img/line.png)

## Intro - motivation

Suite à un entretien j'ai reçu un feedback. Ce feedback n'était pas négatif, il était correctif. Il m'a permis de regarder en face une réalité que je ne voulais pas voir : mes compétences techniques relatives au dev backend ne sont pas au niveau.
En effet je me suis un peu dispersé. J'adore les problématiques liées à l'humain, aux process, à l'organisation et aussi l'architecture... J'y ai investi du temps, notamment chez Leboncoin. Mais j'ai perdu le focus sur le code. Cela ne serait pas un problème si je souhaitais naviguer vers d'autre métier, mais ce n'ai pas le cas. Je veux faire du dev back.  

A la suite de cet entretien, je dois réagir. Et je veux garder une trace de cette expérience parce qu'elle me donne un exemple de la valeur du feedback. Comment suite à un échec, un feedback peut me donner de l'énergie pour corriger et m'améliorer.  

Ma volonté: je veux affiner mes connaissances techniques, les approfondir. Ce n'est pas normal de passer 2 ans à travailler en tant que dev back Go sur du micro service et des API chez Leboncoin et arriver plus d'un ans après en entretien et ne pas savoir expliquer ce qu'est un "context" dans Go.  

Mes actions: installer une habitude de lecture, d'expérimentation de code et d'outil. Faire des katas. Les ressources ne manquent pas! Aujourd'hui j'ai du temps, c'est le moment. Le bon moment aurait été avant de passer à côté d'entretiens. Un petit auto shame et au travail!  
Si j'arrive à faire régulièrement du sport, je dois réussir à investir régulièrement du temps sur ma technique, indépendamment de mes expériences professionnelles.

## Liens

- [Dave.Cheney](https://dave.cheney.net/)
- [Katas](https://codingdojo.org/kata/)

## Katas

Je vais faire un premier Kata sans objectif particulier. J'ai seulement lu l'introduction de [Practical-Go](https://dave.cheney.net/practical-go/presentations/qcon-china.html
) et il y a beaucoup de guide line dans ce texte. Je vais essayer à chacun de mes Katas d'ajouter des guide lines. Ainsi Mon premier sera libre, puis il y aura de plus en plus de point. Je devrais pouvoir observer une evolution interressante de mon code.

### Bank OCR

Pour ce premier Kata je n'ai pas grand chose à ajouter. Je pense que le point central est la manière de gérer les digits. J'ai traduit un digits en 3 lignes de chacune 3 valeurs qui se lit de haut en bas et de gauche à droite. Ainsi un 1 signifie qu'il y a un `|` si on est en colonne 1 ou 3, ou un `_` en colonne 2. Cette string est coûteuse et convertir ça en binaire puis en entier serait plus efficace.
 
Dans ce premier kata je n'ai pas pris en compte de guide line. En revanche c'est l'occasion de creuser certaines définitions et concepts que je ne maîtrise pas.
Le `checksum` est une string de chiffres et de lettres qui indique le nombre de bits qui doit être présent dans une transmission de données. Si le nombre de bits indiqué par le checksum et le nombre de bits réel sont différents, il y a un problème. Le `checksum` est utilisé avec les protocoles `TCP` et `UDP`. Le `checksum` value est assigné en exécutant une `cryptographic hash function`. C'est un algorithme qui traduit un message en bit d'une taille fixe appelé `hash-value` ou message digest. Ce mapping est déterministe et quasi impossible à inverser. Il existe plusieurs exemples de fonctions comme les `SHA` ou `DSA` ou encore `MD5`.

Un des contextes où j'ai rencontré ces algorithmes est la génération de clés ssh. J'ai voulu testé leur déterminisme en rentrant la même passphrase, sauf que je n'obtiens pas le même résultat. L'occasion de découvrir l'existence de la couche de `RNG (Random Number Generator)` dans l'algorithme de génération de clé. Plus le `RNG` est exécuté dans un environnement avec de l'`entropy` plus il est efficace. L'occasion d'essayer de comprendre le principe d'`entropy` avec une [video](https://www.youtube.com/watch?v=YM-uykVfq_E).
 
Une `hash table` ou `hash map` fonctionne sur le même principe: on map un clé en la donnant à un `hash fonction` avec un index derrière lequel se trouve la valeur que l'on souhaite trouver. Il est donc bien question d'une map mais avec une complexité de recherche et d'insertion en moyenne de Θ(1).
La 'hashmap' est utilisée quand on crée des `index` en base de données. Il y a un coût en espace mémoire de Θ(n). C'est ce point qui oblige à utiliser les `index` de manière réfléchie.



## TODO

- array / slice
- API
  - JWT
  - Token / certificat
  - http / https
  - context
- DataBase
  - exp PSQL
  - lib go PSQL
  - journalisation BD


---

Simon P