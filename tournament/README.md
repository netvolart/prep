## Taks 2

Question
You are hosting a tennis tournament with many players participating. All your players are pre-ranked based on their historical performance. In an elimination tournament, the loser of each match-up is immediately eliminated from the tournament. Each winner will play another in the next round, until the final match-up, whose winner becomes the tournament champion.

Part 1 
Can you write a program to predict the players in each round? Your program should naively assume that a higher ranked player always wins against a lower ranked player. 

Input       
 [1, 2, 3, 4, 5, 6, 7, 8]  //ranks of each player and the match ups. The smaller the number, the higher the rank.

Output:
 1, 2, 3, 4, 5, 6, 7, 8 //round 1
 1, 3, 5, 7  //round 2
 1, 5  //round 3:
 1   //round 4

round1     round2      round3       round 4

1 - - \
            1  - - \
2 - - /
                       1  - - \
3 - - \
            3 - - /
4 - - /
                                    1 
5 - - \
            5 - - \
6 - - /
                        5 - - /
7 - - \
            7 - - /
8 - - /


Part 2 
Can you create a draw generator that makes sure the better players do not meet in the earlier rounds? To be more specific, a higher-ranked player should never be eliminated in an earlier round than someone ranked lower according to your prediction in part 1.

The input is a total number of players, output a possible draw

Input : 8 
1, 2, 3, 4, 5, 6, 7, 8
Output:
1 5 3 7 2 6 4 8
Explanation:
1 5 3 7 2 6 4 8
1 3 2 4
1 2
1
