# Bitmap Definition
In the 64-bit map, the least significant bit (LSB) will correspond to a1 and the most significant bit (MSB) will correspond to h8

| - | a  | b  | c  | d  | e  | f  | g  | h  |
| - | -  | -  |  - | -  | -  | -  | -  | -  |
| 8 | 56 | 57 | 58 | 59 | 60 | 61 | 62 | 63 |
| 7 | 48 | 49 | 50 | 51 | 52 | 53 | 54 | 55 |
| 6 | 40 | 41 | 42 | 43 | 44 | 45 | 46 | 47 |
| 5 | 32 | 33 | 34 | 35 | 36 | 37 | 38 | 39 |
| 4 | 24 | 25 | 26 | 27 | 28 | 29 | 30 | 31 |
| 3 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 |
| 2 | 8  | 9  | 10 | 11 | 12 | 13 | 14 | 15 |
| 1 | 0  | 1  | 2  | 3  | 4  | 5  | 6  | 7  |

## What data to store to undo a move:
capture: piece type -> 3 bits

captured side: 1 bit

castling: 4 bits

ep sq: ep available -> 1 bit

  ep sq -> 7 bits

| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 |
| - | - | - | - | - | - | - | - | - | - | -  | -  | -  | -  | -  |  - |
| capture side | > | cap type | < | wqsc | wksc | bqsc | bksc | ep available | >>> | >>  | >  | ep sq  | <  | <<  |  <<< |