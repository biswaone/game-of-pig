# game-of-pig
go bootcamp exercise https://playbook.one2n.in/go-bootcamp/go-projects/game-of-pig

run the make file to create a binary ./pig </br>
```
cd game-of-pig
make
```
### story 1

```
./pig 20 30
 ```
 ### Output 
 ```
 Holding at 10 vs Holding at 20: wins: 4/10 (40.0%), losses: 6/10 (60.0%)
 ```
 ### story 2

```
./pig 20 1-100
 ```
 ### Output
 ```
 Holding at 21 vs Holding at 1: wins: 10/10 (100.0%), losses: 0/10 (0.0%)
Holding at 21 vs Holding at 2: wins: 10/10 (100.0%), losses: 0/10 (0.0%)
Holding at 21 vs Holding at 3: wins: 10/10 (100.0%), losses: 0/10 (0.0%)
...
...
Holding at 21 vs Holding at 10: wins: 6/10 (60.0%), losses: 4/10 (40.0%)
Holding at 21 vs Holding at 11: wins: 6/10 (60.0%), losses: 4/10 (40.0%)
Holding at 21 vs Holding at 12: wins: 7/10 (70.0%), losses: 3/10 (30.0%)
...
...
Holding at 21 vs Holding at 100: wins: 10/10 (100.0%), losses: 0/10 (0.0%)
 ```