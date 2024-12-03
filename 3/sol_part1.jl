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
        # println(file_content)
    catch e
        println("Error reading file: $e")
    end

    pattern = r"mul\((\d{1,3}),(\d{1,3})\)"
    matches = [m.match for m in eachmatch(pattern, file_content)]

    sum = 0

    for str in matches
        m = match(pattern, str)

        if m !== nothing
            x = parse(Int, m.captures[1])
            y = parse(Int, m.captures[2])
            sum = sum + x * y
        end
    end

    println(sum)
end

# Call the main function to run the program
main()
