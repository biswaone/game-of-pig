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

### story 3
```
./pig 1-100 1-100
```

### Output
```
Result: Wins, losses staying at k = 1: 399/990 (40.30%), 591/990 (59.70%)
Result: Wins, losses staying at k = 2: 389/990 (39.29%), 601/990 (60.71%)
Result: Wins, losses staying at k = 3: 445/990 (44.95%), 545/990 (55.05%)
...
...
Result: Wins, losses staying at k = 99: 242/990 (24.44%), 748/990 (75.56%)
Result: Wins, losses staying at k = 100: 285/990 (28.79%), 705/990 (71.21%)
```
