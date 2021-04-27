#!/bin/python3
import os


#
# Complete the 'countMatches' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. STRING_ARRAY grid1
#  2. STRING_ARRAY grid2
#
def search_grid(grid1, grid2, i, j):
    """
    Searches the grid horizontally and vertically to identify a matching region.
    :param grid1: first grid
    :param grid2: second grid
    :param i: row index to search from
    :param j: column index to search from
    :return: True if a matching region is found else False
    """
    if i < 0 or j < 0 or i >= len(grid1) or j >= len(grid1[0]):  # boundary check
        return True
    match = grid1[i][j] == grid2[i][j]
    if grid1[i][j] == 0 or grid2[i][j] == 0:
        return match
    # once a cell becomes a part of a matching region, set it to 0. This makes sure that the cell
    # is not counted for another matching region.
    grid1[i][j] = 0
    grid2[i][j] = 0
    match = search_grid(grid1, grid2, i - 1, j) and match
    match = search_grid(grid1, grid2, i, j - 1) and match
    match = search_grid(grid1, grid2, i + 1, j) and match
    match = search_grid(grid1, grid2, i, j + 1) and match
    return match


def countMatches(g1, g2):
    """
    Loop thru all the cells in the grid, if a potential matching region is found, calls search_grid.
    :param g1: first grid
    :param g2: second grid
    :return: number of matching regions
    """
    if g1 is None or g2 is None or len(g1) == 0 or len(g1[0]) == 0:  # sanity check
        return 0
    count = 0
    for i in range(len(g1)):
        for j in range(len(g1[0])):
            if g1[i][j] == g2[i][j] == 1 and search_grid(g1, g2, i, j):
                count = count + 1
    return count


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')
    grid1_count = int(input().strip())
    grid1 = []
    for _ in range(grid1_count):
        grid1_item = input()
        grid1.append(list(map(int, list(grid1_item))))
    grid2_count = int(input().strip())
    grid2 = []
    for _ in range(grid2_count):
        grid2_item = input()
        grid2.append(list(map(int, list(grid2_item))))
    result = countMatches(grid1, grid2)
    fptr.write(str(result) + '\n')
    fptr.close()
