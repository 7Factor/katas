fn get_prime_factors(mut number: i32) -> Vec<i32> {
    let mut result: Vec<i32> = Vec::new();

    let mut divisor = 2;
    
    while number > 1 {
        while number % divisor == 0 {
            result.push(divisor);
            number = number / divisor;
        }
        divisor += 1;
    }

    return result;
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn prime_factor_works() {
        let result = get_prime_factors(10);
        assert_eq!(result.len(), 2);
        assert_eq!(result[0], 2);
        assert_eq!(result[1], 5);
    }

    #[test]
    fn prime_factor_works_for_large_factors() {
        let result = get_prime_factors(3065);
        assert_eq!(result.len(), 2);
        assert_eq!(result[0], 5);
        assert_eq!(result[1], 613);
    }
}