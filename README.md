# termix

Basic text editor with VIM like motion which brings nothing new.

## Motivation

Text editor is something I use on daily basis but how does it actually work? how is the text represented?
how does undo/redo, highlighting and many more things. There is a lot that is hiding behind even a simple
text editor implementation and I'd like to explore it and give it my take. This will be my attemp to create fully working
text editor from scratch


## references and resources
- https://en.wikipedia.org/wiki/Piece_table
- https://en.wikipedia.org/wiki/Red%E2%80%93black_tree
- https://www.cs.unm.edu/~crowley/papers/sds.pdf data structures for text sequances

- https://viewsourcecode.org/snaptoken/kilo tutorial on how to build text editor in C (uses vector)
- https://code.visualstudio.com/blogs/2018/03/23/text-buffer-reimplementation VScode textBuffer reimplementation
- https://bartoszmilewski.com/2013/11/25/functional-data-structures-in-c-trees Master mind RB tree implementation


## Requirements/features:
- [] load file
- [] save file
- [x] basic movement around text (VIM like)
- [] insert
- [] delete
- [] selection
- [x] mode switching (edit, normal)
- [] support for mouse

If I'll be brave enough ;)
- [] redo/undo
- [] word wrapping

