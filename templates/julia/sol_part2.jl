# sol_part1.jl
function main()
    # Check if a file was provided as an argument
    if length(ARGS) != 1
        println("Please provide exactly one file as a parameter.")
        return
    end
    
    # Get the filename from the arguments
    filename = ARGS[1]
    
    # Try to open the file and read its contents
    try
        file_content = read(filename, String)
        # println(file_content)
    catch e
        println("Error reading file: $e")
    end
end

# Call the main function to run the program
main()
