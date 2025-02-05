const fs = require('fs');

function convertToGrid(input) {
    return input.split('\n').map(line => line.split(''));
}

function findGuard(grid) {    
    const specialChars = ['^', '>', '<', 'v'];
    
    for (let row = 0; row < grid.length; row++) {
        for (let col = 0; col < grid[row].length; col++) {
            if (specialChars.includes(grid[row][col])) {
                return { row, col, char: grid[row][col] };
            }
        }
    }
    
    // Return null if not found
    return null;
}

function execute(grid, guard) {
    let outside = false;
    let prevStep = 0; // 1 -> row++, 2 -> row--, 3 -> col++, 4 -> col--
    const visited = new Set();

    while (!outside) {
        const key = `${guard.row},${guard.col},${guard.char}`;
        if (visited.has(key)) {
            return true; // Loop detected
        }
        visited.add(key);

        if (guard.row >= grid.length || guard.row < 0 || guard.col >= grid[guard.row].length || guard.col < 0) {
            outside = true;
        } else if (grid[guard.row][guard.col] == "#") {
            switch (prevStep) {
                case 1:
                    guard.char = ">";
                    guard.row++;
                    break;
                case 2:
                    guard.char = "<";
                    guard.row--;
                    break;
                case 3:
                    guard.char = "v";
                    guard.col--;
                    break;
                case 4:
                    guard.char = "^";
                    guard.col++;
                    break;
                default:
                    console.log("Error: invalid state.") 
            }
        } else {
            if (grid[guard.row][guard.col] == ".") {
                grid[guard.row][guard.col] = "X";
            }

            switch (guard.char) {
                case "^":
                    guard.row--;
                    prevStep = 1;
                    break;
                case ">":
                    guard.col++;
                    prevStep = 3;
                    break;
                case "<":
                    guard.col--;
                    prevStep = 4;
                    break;
                case "v":
                    guard.row++;
                    prevStep = 2;
                    break;
                default:
                    console.log("Error: invalid state.") 
            }
        }
    }

    return false; // No loop detected
}

function part2(grid, guard) {
    const originalGrid = JSON.parse(JSON.stringify(grid));
    let count = 0;

    for (let row = 0; row < grid.length; row++) {
        for (let col = 0; col < grid[row].length; col++) {
            if (grid[row][col] === '.' && !(row === guard.row && col === guard.col)) {
                grid[row][col] = '#';
                const loopDetected = execute(grid, { ...guard });
                if (loopDetected) {
                    count++;
                }
                grid = JSON.parse(JSON.stringify(originalGrid)); // Reset grid
            }
        }
    }

    return count;
}

function main() {
    if (process.argv.length !== 3) {
        console.log('Please provide exactly one file as a parameter.');
        return;
    }

    const filePath = process.argv[2];
    let fileContent = "";

    try {
        fileContent = fs.readFileSync(filePath, 'utf-8');
    } catch (error) {
        console.error(`Error reading file: ${error.message}`);
    }

    const grid = convertToGrid(fileContent.trim());
    const guard = findGuard(grid);

    const result = part2(grid, guard);
    console.log(result);
}

main();