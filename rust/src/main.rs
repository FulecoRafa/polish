use std::{env};

fn main() {
    let mut stack: Vec<f32> = Vec::new();
    for operand in env::args().skip(1) {
        match operand.parse::<f32>() {
            Ok(num) => stack.push(num),
            Err(_) => {
                let right = stack.pop().unwrap();
                let left = stack.pop().unwrap();
                let resul = match &operand[..] {
                    "+" => left+right,
                    "-" => left-right,
                    "*" => left*right,
                    "/" => left/right,
                    "^" => left.powf(right),
                    _ => panic!("{} is not a valid operation", operand)
                };
                stack.push(resul);
            }
        };
    };
    println!("{}", stack.pop().unwrap())
}
