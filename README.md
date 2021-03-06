# halken
>Video games are meant to be just one thing. Fun. Fun for everyone.

*Satoru Iwata*

<p align="center">
  <img src="https://i.imgur.com/opqDZes.gif">
</p>

Halken is a Game Boy emulator written in Go being developed during my time at the [Recurse Center](https://recurse.com).
The name is one used by HAL Laboratories for a time. HAL was the first company where Satoru Iwata was a video game programmer.

I intend to have lots of comments as well as a document regarding the process. Once I am happy with 32KB games generally working I'll be cleaning the code and writing documentation for others who want to tackle the same project.

## Known working games

**Usage**: `halken /path/to/rom`

1. Tetris
2. Dr. Mario
3. Flipull
4. World Bowling
5. Minesweeper
6. Motorcross Maniacs
7. Hyper Lode Runner
8. Battlecity

## Known bugs
* LCD STAT interrupt is only partially correct - need to make sure all cases are covered
  * Breaks certain games like Game of Harmony, which gets weird graphics due to it not firing when it should
* Sprite color palettes besides default are unimplemented
* 8x16 sprites aren't drawn properly, get cut off

## TODO

1. ~Implement CPU opcodes~
    * ~Non-CB opcodes implemented~
     * ~Write dispatch loop~
     * Use blargg's test output to fix instructions
       * ~01 - special~
       * 02 - interrupts
       * ~03 - op sp,hl~
       * ~04 - op r,imm~
       * ~05 - op rp~
       * ~06 - ld r,r~
       * ~07 - jr,jp,call,ret,rst~
       * ~08 - misc instrs~
       * ~09 - op r,r~
       * ~10 - bit ops~
       * ~11 - op a,(hl)~
2. ~Implement memory~
3. ~Test GB bootstrap ROM~
4. ~Draw tiles~
5. ~Draw background~
6. ~Graphics loop~
7. Interrupts
   * ~VBlank~
   * ~LCD~
   * ~Timer~
   * ~Serial~
   * ~Joypad~
8. ~Timer~
9. ~Refactor (direct memory access vs. abstractions), lower LOC, comment~
10. Document process for learners
