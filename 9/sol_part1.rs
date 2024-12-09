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

    let mut disk = String::new();

    for (i, ch) in content.chars().enumerate() {
        println!("Index: {}, Character: {}", i, ch);

        if let Some(digit) = ch.to_digit(10) {
            if i % 2 == 0 {
                // Append the result of multiplication
                disk += &(i/2).to_string().repeat(digit as usize);
            } else {
                // Append a dot
                disk += &".".repeat(digit as usize);
            }
        } else {
            println!("Character '{}' is not a digit", ch);
        }
    }

    let mut i = 0;
    let mut j = disk.chars().count() - 1;

    let mut chars: Vec<char> = disk.chars().collect();

    while i < j {
        if chars[i] != '.' {
            i = i + 1;
        } else {
          if chars[j] == '.' {
            j = j - 1;
          } else {
            chars[i] = chars[j];
            chars[j] = '.';
            i = i + 1;
            j = j - 1;
          }
        }
    }

    let mut sum = 0;

    for (i, ch) in chars.iter().enumerate() {
        if let Some(digit) = ch.to_digit(10) {
            sum = sum + i * (digit as usize);
        } else {
            println!("Character '{}' is not a digit", ch);
        }
    }

    println!("{}", sum);

    Ok(())
}
