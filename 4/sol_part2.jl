function main()
    # Check if a file was provided as an argument
    if length(ARGS) != 1
        println("Please provide exactly one file as a parameter.")
        return
    end

    # Read the input file
    filename = ARGS[1]
    file_content = ""
    try
        file_content = read(filename, String)
    catch e
        println("Error reading file: $e")
        return
    end

    # Prepare the grid
    lines = split(file_content, '\n')
    grid = [collect(line) for line in lines]

    counter = 0

    centre_char = 'A'
    prev_char = 'M'
    next_char = 'S'
    
    @assert (grid_size = length(grid)) == length(grid[1])

    directions = [
        (1, 1),  # down right
        (1, -1), # down left
        (-1, 1), # up right
        (-1, -1) # up left
    ]

    for row in 2:grid_size-1
        for col in 2:grid_size-1
            if grid[row][col] == centre_char
                dir_count = 0
                for direction in directions
                    if grid[row - direction[1]][col - direction[2]] == prev_char && 
                       grid[row + direction[1]][col + direction[2]] == next_char
                        dir_count += 1
                    end
                    if dir_count == 2
                        counter += 1
                        break
                    end
                end
            end
        end
    end

    println(counter)
end

# Run the program
main()
