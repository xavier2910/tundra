# Library Concepts and Methodology

\(README Part 2\)

So I realized when creating this, I didn't have a super
clear idea of how, exactly, everything was supposed to fit
together. So I'm writing this as much as an aid to myself
as for anyone who ends up using my library \(if anyone does :\).

Maybe this document is redundant with the godocs. Sue me.

## Structures

So generally, the structures reperesent logical units in
the game, such as the player, game locations, and
manipulable game objects. Documentation for each structure
can be found in the source code \(godocs\).

## Commands

This gets a little hairier. A command is just a function that runs when a specific thing is typed by the user. The approach here is to let everything have commands, and the commands \(generally speaking\) should belong to the actor. this is how the `turnBased` command processor works. If you
write your own command processor, keep in mind that this
is what I intended. See the godocs for more info.
