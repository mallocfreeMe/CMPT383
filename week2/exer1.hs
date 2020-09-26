-- First Haskell Code
det a b c =  b^2 - 4*a*c
quadsol1 a b c = (-b - sqrt(det a b c))/2*a
quadsol2 a b c = (-b + sqrt(det a b c))/2*a

-- Writing Your First Code
third_a xs = xs!!2
third_b (_:_:c:_) = c

-- Factorial
fact 0 = 1
fact 1 = 1
fact a = (fact(a - 1)) * a

-- Hailstone Function
hailstone n
  | (even n == True) = n `div` 2
  | otherwise = 3*n+1

-- Hailstone Length
hailLen 0 = 0
hailLen 1 = 0
hailLen n = hailLen (hailstone n) + 1