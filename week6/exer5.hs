import Data.Ratio

-- Built-In Functions
myIterate :: (a -> a) -> a -> [a]
myIterate f x = [x] ++ myIterate f (f x)

myTake :: Int -> [a] -> [a]
myTake 0 _ = []
myTake _ [] = []
myTake n (x:xs) = x : myTake (n-1) xs

myDrop :: Int -> [a] -> [a]
myDrop 0 xs = xs
myDrop _ [] = []
myDrop n (x:xs) = myDrop (n-1) xs

mySplitAt :: Int -> [a] -> ([a], [a])
mySplitAt a b = (myTake a b, myDrop a b)

-- Rational Numbers
rationalSum :: Integral a => a -> [Ratio a]
rationalSum n = [a % (n-a) | a <- [1..n-1]]

-- Lowest Terms Only]
rationalSumLowest :: Integral a => a -> [Ratio a]
rationalSumLowest n = [a % (n-a) | a <- [1..n-1], gcd a (n-a) == 1]

-- All Rational Numbers
rationals :: Integral a => [Ratio a]
rationals = concat [rationalSumLowest n | n <- [1..]]

-- Input/Output
splitAtSeparator :: Eq a => a -> [a] -> [[a]]
splitAtSeparator sep [] = []
splitAtSeparator sep content = first : splitAtSeparator sep rest
    where
    first = takeWhile (/= sep) content
    firstlen = length first
    rest = drop (firstlen+1) content

readInt :: String -> Int
readInt = read

convertToSum :: [String] -> Int
convertToSum [] = 0
convertToSum xs = readInt (head xs) + convertToSum (myDrop 1 xs)

sumFile :: IO ()
sumFile = do
    a <- readFile("input.txt")
    let sum = convertToSum (splitAtSeparator '\n' a)
    print sum