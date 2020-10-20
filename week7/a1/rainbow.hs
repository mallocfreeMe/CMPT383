import RainbowAssign
import qualified Data.Map as Map
import Data.Maybe

-- Parameters
pwLength, nLetters, width, height :: Int
filename :: FilePath
pwLength = 8            -- length of each password
nLetters = 5            -- number of letters to use in passwords: 5 -> a-e
width = 40              -- length of each chain in the table
height = 1000           -- number of "rows" in the table
filename = "table.txt"  -- filename to store the table

-- pwReduce functioncon
-- param: takes one argment: a hash value 
-- post: should rerturn a possible password
pwReduce :: Hash -> Passwd
pwReduce hash = map toLetter (reverse $ take pwLength (leastSignificant $ fromEnum hash))     
  where
    leastSignificant :: Int -> [Int]
    leastSignificant hashInt = [mod hashInt nLetters] ++ leastSignificant (div hashInt nLetters)

-- rowRecursive function: a helper method for rainbowTable
-- param: takes two argument: the width of the table, an inital Passwd value
-- post: should return the final Hash value at the end of the chain 
rowRecursive :: Int -> Passwd -> Hash
rowRecursive 0 passowrd = pwHash(passowrd)
rowRecursive width password = rowRecursive (width - 1) (pwReduce(pwHash(password)))

-- ranbow Table function
-- param: takes two arguments: the width of the table, and the list of initial passwords.
-- post: should return a Map.Map that maps the final Hash values onto the Passwd values at the start of the chain.
rainbowTable :: Int -> [Passwd] -> Map.Map Hash Passwd
rainbowTable width [] = Map.empty
rainbowTable 0 (p:ps) = Map.fromList (map formTuple (p:ps))
  where 
       formTuple :: Passwd -> (Hash, Passwd)
       formTuple p = (pwHash p, p)
rainbowTable width (p:ps) = Map.fromList (map formRowRecursiveTuple (p:ps))
  where 
       formRowRecursiveTuple :: Passwd -> (Hash, Passwd)
       formRowRecursiveTuple p = (rowRecursive width p, p)

-- newPassword Function: helper method for createRow
-- PARAM: takes two arugument: width, password
-- POST: should return a new password
newPassword :: Int -> Passwd -> Passwd
newPassword 0 password = password
newPassword width password = newPassword (width-1) (pwReduce(pwHash(password)))

-- createRow Function: helper method for createTable
-- PARAM: takes two arugument: width, password
-- POST: should return a row of the rainbow table in a list format
createRow :: Int -> Passwd -> [(Hash, Passwd)]
createRow 0 password = [(pwHash(password), password)]
createRow width password = [(pwHash(newPassword (width) (password)), (newPassword (width) (password)))] ++ (createRow (width-1) (password))

-- createTable Function: helper method for findPassword
-- PARAM: takes two arugument: width, a list of passwords
-- POST: should return the whole rainbow table in a list format
createTable :: Int -> [Passwd] -> [(Hash, Passwd)]
createTable width [] = []
createTable width (p:ps) = (createRow width p) ++ createTable width (ps)

-- findPassword Function
-- PARAM: takes three arugument: the rainbow table, the width of the table, and the hash value the function trying to reverse.
-- POST: The function should return a Maybe Passwd
findPassword :: Map.Map Hash Passwd -> Int -> Hash -> Maybe Passwd
findPassword table width hash = Map.lookup hash fullTable
  where 
    fullTable = Map.fromList(createTable width (Map.elems table))

-- generateTable function: generate a new table and save it to disk
generateTable :: IO ()
generateTable = do
  table <- buildTable rainbowTable nLetters pwLength width height
  writeTable table filename

-- test1 function: To read the table back
test1 = do
  table <- readTable filename
  return (Map.lookup 0 table)

-- test2 function: tries to crack n randomly generated passwords, and reports the results
test2 :: Int -> IO ([Passwd], Int)
test2 n = do
  table <- readTable filename
  pws <- randomPasswords nLetters pwLength n
  let hs = map pwHash pws
  let result = Data.Maybe.mapMaybe (findPassword table width) hs
  return (result, length result)

main :: IO ()
main = do
  generateTable
  res <- test2 10000
  print res 
