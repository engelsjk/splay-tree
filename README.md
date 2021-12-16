# splay-tree

A minimal Go port of the JavaScript [w8r/splay-tree](https://github.com/w8r/splay-tree) data structure.

## Implementation Status

At the moment, this library is only intended to support the development of a pure Go implementation of the Martinez-Rueda-Feito polygon clipping algorithm, based on the JavaScript library [mfogel/polygon-clipping](https://github.com/mfogel/polygon-clipping). Therefore, it does not have full parity with [w8r/splay-tree](https://github.com/w8r/splay-tree) and for now only implements those methods used in [mfogel/polygon-clipping](https://github.com/mfogel/polygon-clipping).

### Tree Methods

* [X] Insert
* [X] Add
* [X] Remove
* [X] Pop
* [ ] FindStatic
* [X] Find
* [ ] Contains
* [ ] ForEach
* [ ] Range
* [ ] Keys
* [ ] Values
* [X] Min
* [X] Max
* [X] MinNode
* [X] MaxNode
* [ ] At
* [X] Next
* [X] Prev
* [ ] Clear
* [ ] ToList
* [ ] Load
* [ ] IsEmpty
* [X] Size
* [ ] Root
* [ ] ToString
* [ ] Update
* [ ] Split