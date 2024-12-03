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
    
    mul_enabled = true
    sum = 0
    
    # Find all mul, do(), and don't() instructions
    mul_matches = collect(eachmatch(r"mul\((\d+),(\d+)\)", file_content))
    do_matches = collect(eachmatch(r"do\(\)", file_content))
    dont_matches = collect(eachmatch(r"don't\(\)", file_content))
    
    # in order
    instructions = sort(
        vcat(
            [(m.offset, "mul", parse(Int, m.captures[1]), parse(Int, m.captures[2])) for m in mul_matches],
            [(m.offset, "do", 0, 0) for m in do_matches],
            [(m.offset, "dont", 0, 0) for m in dont_matches]
        ), 
        by=x->x[1]
    )
    
    for (_, type, x, y) in instructions
        if type == "do"
            mul_enabled = true
        elseif type == "dont"
            mul_enabled = false
        elseif type == "mul" && mul_enabled
            sum += x * y
        end
    end
    
    println(sum)
end

main()