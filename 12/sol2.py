from collections import deque

def read_input(file_path):
    with open(file_path, 'r') as f:
        return [list(line.strip()) for line in f]

def get_neighbors(x, y, rows, cols):
    directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]
    return [(x + dx, y + dy) for dx, dy in directions if 0 <= x + dx < rows and 0 <= y + dy < cols]

def bfs(grid, x, y, visited):
    rows, cols = len(grid), len(grid[0])
    queue = deque([(x, y)])
    region_cells = []
    visited.add((x, y))
    plant_type = grid[x][y]
    
    while queue:
        cx, cy = queue.popleft()
        region_cells.append((cx, cy))
        
        for nx, ny in get_neighbors(cx, cy, rows, cols):
            if (nx, ny) not in visited and grid[nx][ny] == plant_type:
                visited.add((nx, ny))
                queue.append((nx, ny))
    
    return region_cells

def count_sides(region_cells, grid):
    rows, cols = len(grid), len(grid[0])
    edges = set()
    
    for x, y in region_cells:
        for dx, dy in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
            nx, ny = x + dx, y + dy
            if (nx, ny) not in region_cells:
                edges.add(((x, y), (nx, ny)))
    
    return len(edges) // 2

def calculate_fence_cost(grid):
    rows, cols = len(grid), len(grid[0])
    visited = set()
    total_cost = 0
    
    for x in range(rows):
        for y in range(cols):
            if (x, y) not in visited:
                region_cells = bfs(grid, x, y, visited)
                area = len(region_cells)
                sides = count_sides(region_cells, grid)
                total_cost += area * sides
    
    return total_cost

def main():
    grid = read_input("input.txt")
    result = calculate_fence_cost(grid)
    print(result)

if __name__ == "__main__":
    main()
