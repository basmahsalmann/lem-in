# Lem in

## Description

This project takes files that consists of a map with connecting rooms. It utilizes an algorithm that finds the most optimized way to send ants from the start to end room, with the least amount of steps.

## Authors
- Zainab Saeed (zsaeed)
- Basma Sallman (bsalman)
- Malak Aljamri (maljamri)
- Nourose Mohammed (nmohammed)

## Algorithm Overview

1. BFS Pathfinding:
    Identifies all valid paths from the start to the end room using BFS .
2. Path Filtering:
    Recursively removes paths with shared rooms, retaining one path per room until all combinations are processed.
3. Optimal Path Selection:
    Compares path lengths and selects the set with the maximum non-clashing paths.
4. Ant Assignment and Movement:
    Assigns and moves ants along the optimal paths efficiently.

## How It Works

Starting Point: Ants begin in the ##start room.
Goal: Move ants to the ##end room with minimal moves.
Movement Constraints: Ants move one room per turn with minimal turns as possible.

## Running the Project

1. Ensure you have a Linux terminal and Go installed.
2. Clone the repository:
3. Navigate to the project directory:

```bash
cd lem_in
Run the project:
go run . Examples/<file.txt>
Example:
go run . Examples/example00.txt
go run . Examples/example01.txt
go run . Examples/example03.txt



