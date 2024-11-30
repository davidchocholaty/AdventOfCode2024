use std::env;
use std::fs::File;
use std::io::{self, Read};

fn main() -> io::Result<()> {
    // Check if a file path is provided as a command-line argument
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 {
        eprintln!("Usage: cargo run <file-path>");
        std::process::exit(1);
    }

    // Get the file path from the first command-line argument
    let file_path = &args[1];

    // Open the file
    let mut file = File::open(file_path)?;

    // Read the content of the file into a string
    let mut content = String::new();
    file.read_to_string(&mut content)?;

    // Print the content of the file to the console
    // println!("{}", content);

    Ok(())
}
