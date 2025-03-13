use std::collections::HashSet;
use std::fs::read_to_string;

type Position = (usize, usize);

fn read_input(filename: &str) -> (Vec<String>, String) {
    let content = read_to_string(filename).expect("Failed to read file");
    let mut map_lines = Vec::new();
    let mut moves = String::new();
    
    for line in content.lines() {
        if line.contains('#') || line.contains('@') {
            map_lines.push(line.to_string());
        } else if line.contains('^') || line.contains('v') || line.contains('<') || line.contains('>') {
            moves.push_str(line.trim());
        }
    }
    
    (map_lines, moves)
}

fn find_robot_position_and_boxes(warehouse: &[String]) -> (Position, HashSet<Position>) {
    let mut robot_pos = (0, 0);
    let mut boxes = HashSet::new();
    
    for (r, row) in warehouse.iter().enumerate() {
        for (c, char) in row.chars().enumerate() {
            if char == '@' {
                robot_pos = (r, c);
            } else if char == 'O' {
                boxes.insert((r, c));
            }
        }
    }
    
    (robot_pos, boxes)
}

fn simulate_robot_moves(mut warehouse: Vec<String>, moves: &str) -> HashSet<Position> {
    let directions = [('^', (-1, 0)), ('v', (1, 0)), ('<', (0, -1)), ('>', (0, 1))];
    let direction_map: std::collections::HashMap<char, (isize, isize)> = directions.iter().cloned().collect();
    
    let (mut robot_pos, mut boxes) = find_robot_position_and_boxes(&warehouse);
    let mut warehouse_chars: Vec<Vec<char>> = warehouse.iter().map(|row| row.chars().collect()).collect();
    
    for move_char in moves.chars() {
        if let Some(&(dr, dc)) = direction_map.get(&move_char) {
            let new_robot_pos = ((robot_pos.0 as isize + dr) as usize, (robot_pos.1 as isize + dc) as usize);
            
            if warehouse_chars[new_robot_pos.0][new_robot_pos.1] == '#' {
                continue;
            }
            
            let mut current_pos = new_robot_pos;
            let mut chain_positions = Vec::new();
            
            while boxes.contains(&current_pos) {
                let next_pos = ((current_pos.0 as isize + dr) as usize, (current_pos.1 as isize + dc) as usize);
                
                if warehouse_chars[next_pos.0][next_pos.1] == '#' {
                    break;
                }
                
                chain_positions.push(current_pos);
                current_pos = next_pos;
            }
            
            if !chain_positions.is_empty() && warehouse_chars[current_pos.0][current_pos.1] == '.' {
                for &pos in chain_positions.iter().rev() {
                    let next_pos = ((pos.0 as isize + dr) as usize, (pos.1 as isize + dc) as usize);
                    boxes.remove(&pos);
                    boxes.insert(next_pos);
                    warehouse_chars[next_pos.0][next_pos.1] = 'O';
                    warehouse_chars[pos.0][pos.1] = '.';
                }
            }
            
            if !boxes.contains(&new_robot_pos) {
                warehouse_chars[robot_pos.0][robot_pos.1] = '.';
                warehouse_chars[new_robot_pos.0][new_robot_pos.1] = '@';
                robot_pos = new_robot_pos;
            }
        }
    }
    
    boxes
}

fn calculate_gps_sum(boxes: &HashSet<Position>) -> usize {
    boxes.iter().map(|&(r, c)| 100 * r + c).sum()
}

fn main() {
    let (map_lines, moves) = read_input("input.txt");
    let final_boxes = simulate_robot_moves(map_lines, &moves);
    let gps_sum = calculate_gps_sum(&final_boxes);
    println!("Sum of all boxes' GPS coordinates: {}", gps_sum);
}
