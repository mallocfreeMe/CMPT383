import Data.Maybe

-- Pascal's Triangle
pascal :: Int -> [Int]
pascal 0 = [1]
pascal n = [1] ++ map (uncurry (+)) (zip prev (tail prev)) ++ [1] where prev = pascal (n - 1)

-- Pointfree Addition
addPair :: (Int, Int) -> Int 
addPair = uncurry (+)

-- Pointfree Filtering
withoutZeros :: Eq a => Num a => [a] -> [a]
withoutZeros = filter (/= 0)

-- Searching? Maybe?
elementFinding :: Eq a => a -> [a] -> Int -> Int
elementFinding a (x:xs) index
    | a == x = index
    | otherwise = elementFinding a xs (index + 1)

findElt :: Eq a => a -> [a] -> Maybe Int
findElt a b 
    | ((filter (== a) b)) == [] = Nothing
    | otherwise = Just (elementFinding a b 0)
