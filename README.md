# Tato

Tato is a simple implementation of a Volleyball statistics CLI, its primary goal is to provide a simple way to capture in-game statistics through the Tato code, which is specifically designed for it.

## Installation

[TODO]

## The Tato language

The Tato language allows to capture Volleyball actions encoded in the following way:
    
```
[player number][action-kind][class?][evaluation][starting-ending-zone?]
```

Some example of actions could be:
```
12SF+9
```

To understand this better we should go part by part on these elements:

### Player number

Perhaps the most simple of the elements, the jersey number of the player who performed the action, it could be one or two digits.e.g 0, 12, 23, 00, etc.


### Action Kind

The action kind represents the class of action performed by the player, these are well defined by the volleyball as a sport:

```
S - Service
R - Reception
P - Pass
A - Attack
D - Defense
B - Block
```

### Class

This one applies for the Service and Pass kind of action. This is the first of the optional elements as in some cases the user may not want to specify the class of the action.

#### For service:

```
J - Jumping Service
F - Floating Service
X - Jumping Floating Service
```

#### For pass:

```
A - First time pass (typically with the center player)
B - Second time pass
C - Third time pass
```

## Evaluation

One of the most important elements is the evaluation of the action as it allows to analyze performance of the players and the team.

```
= - Double negative
- - Negative
. - Neutral
+ - Positive
* - Double positive
```

## Starting and Ending zones

These are optional and zones can be from 1 to 9. One important thing that when we enter only one zone it means different things depending on the action type:

```
12S+9  > Served TO zone 9
12R=2   > Reception FROM the zone 2
12P=2   > Pass TO the zone 2
12A*3   > Attack FROM the zone 3
12B*3   > Block FROM the zone 3
12D*3   > Block FROM the zone 3
```

When both specified the first is starting and the second is ending zone.

## v1 Tasks

These are the tasks needed to consider Tato v1:

- [x] Action parsing
- [ ] CLI 
- [ ] Game start/end
- [ ] Simple Game statistics (TBD)









