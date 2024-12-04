function main()
    # Check if a file was provided as an argument
    if length(ARGS) != 1
        println("Please provide exactly one file as a parameter.")
        return
    end

    # Get the filename from the arguments
    filename = ARGS[1]
    file_content = ""

    # Try to open the file and read its contents
    try
        file_content = read(filename, String)
    catch e
        println("Error reading file: $e")
        return
    end

    lines = split(file_content, '\n')
    grid = [collect(line) for line in lines]

    counter = 0

    for i in 1:size(grid, 1)
        for j in 1:length(grid[i])
            if grid[i][j] == 'X'
                # Top-left diagonal
                if i >= 4 && j >= 4 && 
                   grid[i-1][j-1] == 'M' && grid[i-2][j-2] == 'A' && grid[i-3][j-3] == 'S'
                    counter += 1
                end

                # Top
                if i >= 4 && 
                   grid[i-1][j] == 'M' && grid[i-2][j] == 'A' && grid[i-3][j] == 'S'
                    counter += 1
                end

                # Top-right diagonal
                if i >= 4 && j <= length(grid[i]) - 3 && 
                   grid[i-1][j+1] == 'M' && grid[i-2][j+2] == 'A' && grid[i-3][j+3] == 'S'
                    counter += 1
                end

                # Left
                if j >= 4 && 
                   grid[i][j-1] == 'M' && grid[i][j-2] == 'A' && grid[i][j-3] == 'S'
                    counter += 1
                end

                # Right
                if j <= length(grid[i]) - 3 && 
                   grid[i][j+1] == 'M' && grid[i][j+2] == 'A' && grid[i][j+3] == 'S'
                    counter += 1
                end

                # Bottom-left diagonal
                if i <= size(grid, 1) - 3 && j >= 4 && 
                   grid[i+1][j-1] == 'M' && grid[i+2][j-2] == 'A' && grid[i+3][j-3] == 'S'
                    counter += 1
                end

                # Bottom
                if i <= size(grid, 1) - 3 && 
                   grid[i+1][j] == 'M' && grid[i+2][j] == 'A' && grid[i+3][j] == 'S'
                    counter += 1
                end

                # Bottom-right diagonal
                if i <= size(grid, 1) - 3 && j <= length(grid[i]) - 3 && 
                   grid[i+1][j+1] == 'M' && grid[i+2][j+2] == 'A' && grid[i+3][j+3] == 'S'
                    counter += 1
                end
            end
        end
    end

    println(counter)
end

main()
