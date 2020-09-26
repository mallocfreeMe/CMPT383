-- Primes and Divisors
divisors :: Int -> [Int]
divisors n = [i | i <- [2..(n `div` 2)], n `mod` i == 0]
primes :: Int -> [Int]
primes n = [i | i <- [2..n], divisors i == []]

-- Pythagorean Triples
pythagorean :: Int -> [(Int, Int, Int)]
pythagorean n = [(a, b, c) | a <- [1..n-1], b <- [1..n-1], c <- [1..n], a ^ 2 + b ^ 2 == c ^ 2 && a < b && b < c]

-- Joining Strings
join :: String -> [String] -> String 
join _ [] = []
join _ [s] = s
join separator (s:xs) = s ++ separator ++ join separator xs

-- Factorial with a fold
fact' :: Int -> Int
fact' n = foldl (*) 1 [1..n]

-- Tail Recursive Hailstone
hailLen :: Int -> Int
hailLen n = hailTail 0 n
  where
    hailTail a 1 = a
    hailTail a n = hailTail (a+1) (hailstone n)

-- Hailstone Function from exercise 1
hailstone :: Int -> Int
hailstone n
  | (even n == True) = n `div` 2
  | otherwise = 3 * n + 1