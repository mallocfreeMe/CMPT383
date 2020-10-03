import Data.Time.Calendar
import Data.Time.Calendar.OrdinalDate

-- Merging
merge :: (Ord a) => [a] -> [a] -> [a]
merge a [] = a
merge [] b = b
merge (a:as) (b:bs)
    | a < b  = [a] ++ merge as (b:bs)
    | a >= b  = [b] ++ merge (a:as) bs


-- Merge Sort
mergeSort :: (Ord a) => [a] -> [a]
mergeSort [] = []
mergeSort [a] = [a] 
mergeSort a = merge (mergeSort firstHalf) (mergeSort secondHalf) 
    where firstHalf = take (div (length a) 2) a 
          secondHalf = drop (div (length a) 2) a

-- Haskell Library and Dates
daysInYear :: Integer -> [Day]
daysInYear y = [jan1..dec31]
    where jan1 = fromGregorian y 1 1
          dec31 = fromGregorian y 12 31

-- learned the $ from https://stackoverflow.com/questions/2035742/is-there-a-haskell-library-for-dates/46548811#46548811
isFriday :: Day -> Bool
isFriday d =  (snd $ mondayStartWeek $ d) == 5

-- Divisors from exercise 2
divisors :: Int -> [Int]
divisors n = [i | i <- [2..(n `div` 2)], n `mod` i == 0]

getDay :: (Integer, Int, Int) -> Int
getDay (y,m,d) = d

isPrimeDay :: Day -> Bool
isPrimeDay d = (length.divisors.getDay.toGregorian) d == 0

primeFridays :: Integer -> [Day]
primeFridays y = filter isPrimeDay (filter isFriday (daysInYear y))
