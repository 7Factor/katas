fn fizz_buzz(min: i32, max: i32) -> Vec<String> {
    let mut v : Vec<String> = Vec::new();

    for i in min..max+1 {
        v.push(output(i))
    }
    return v;
}

fn output(value: i32) -> String {
    let mut output = "".to_string();

    if value % 3 == 0 && value % 5 == 0 {
        output = "FizzBuzz".to_string();
    } else if value % 3 == 0 {
        output = "Fizz".to_string();
    } else if value % 5 == 0 {
        output = "Buzz".to_string();
    } else  {
        output = value.to_string();
    }

    return output;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn fizz_buzz_works() {
        let result = fizz_buzz(1, 100);
        assert_eq!(result.len(), 100);
    }

    #[test]
    fn fizz_buzz_works_2() {
        let result = fizz_buzz(1, 100);
        assert_eq!(result[2 - 1], "2");
        assert_eq!(result[3 - 1], "Fizz");
        assert_eq!(result[4 - 1], "4");
        assert_eq!(result[5 - 1], "Buzz");
        assert_eq!(result[15 - 1], "FizzBuzz");
    }

}
