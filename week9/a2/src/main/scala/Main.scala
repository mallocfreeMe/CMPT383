import java.io._
import java.time.LocalDate
import java.time.format.DateTimeFormatter
import java.util.Calendar

object Demo {
  // Exercise 2: the divisors, primes, join, and pythagorean functions.
  def divisors(a:Int) : Array[Int] = {
    var arr = Array.empty[Int];

    for(i<-2 to a/2) {
      if(a%i == 0) {
        arr = arr :+ i;
      }
    }
    return arr;
  }

  def primes(a:Int) : Array[Int] = {
    var arr = Array.empty[Int];

    for(i<-2 to a) {
      if(divisors(i).size == 0) {
        arr = arr :+ i;
      }
    }

    return arr;
  }

  def join(a:String, b:Array[String]) : String = {
    var s = "";

    if(b.size == 0) {
      s = "";
    } else if(b.size == 1) {
      s = b(0);
    } else {
      s = b(0);

      for(i<-1 to (b.size - 1)) {
        s = s + a + b(i);
      }
    }

    return s;
  }

  def pythagorean(n:Int) : Array[(Int, Int, Int)] = {
    var arr = Array.empty[(Int, Int, Int)];

    for(a<-0 to n-1) {
      for(b<-0 to n -1) {
        for(c<-0 to n) {
          if((a*a + b*b == c*c) && a < b && b < c) {
            arr = arr :+ (a, b, c);
          }
        }
      }
    }

    return arr; 
  }

  // Exercise 3: mergesort (and thus probably also merge)

  def merge(a:Array[Int], b:Array[Int]) : Array[Int]= {
    var arr = Array.empty[Int]

    if(a.size == 0 && b.size == 0) {
      return arr;
    } else if(a.size == 0 && b.size != 0) {
      arr = b;
    } else if(a.size != 0 && b.size == 0) {
      arr = a;
    } else {
      arr = a ++ b;
      for(i<-0 to (arr.size-1)) {
        for(j<-0 to (arr.size-1)) {
          if(arr(i) < arr(j)) {
            var temp = arr(i);
            arr(i) = arr(j);
            arr(j) = temp;
          }
        }
      }
    }

    return arr;
  }

  def merge2(a:String, b:String) : StringBuilder = {
    var result = new StringBuilder("");

    if(a.length == 0 && b.length == 0) {
      return result;
    } else if(a.length == 0 && b.length!= 0) {
      result = new StringBuilder(b);
    } else if(a.length!= 0 && b.length == 0) {
      result = new StringBuilder(a);
    } else {
      result = new StringBuilder(a ++ b);

      for(i<-0 to (result.length-1)) {
        for(j<-0 to (result.length-1)) {
          if(result(i).toInt < result(j).toInt) {
            var temp = result(i);
            result.setCharAt(i, result(j));
            result.setCharAt(j, temp);
          }
        }
      }
    }

    return result;
  }

  def mergesort(a:Array[Int]) : Array[Int] = {
    var arr = Array.empty[Int];

    if(a.size == 0) {
      return arr;
    } else if(a.size == 1) {
      arr = arr ++ a;
      return arr;
    } else {
      var arr1 = Array.empty[Int];
      var arr2 = Array.empty[Int];

      for(i<-0 to (a.size/2-1)) {
        arr1 = arr1 :+ a(i);
      }

      for(i<-(a.size/2) to (a.size - 1)) {
        arr2 = arr2 :+ a(i);
      }

      arr = merge(mergesort(arr1), mergesort(arr2));
    }

    return arr;
  }

  def mergesort2(a:String) : StringBuilder = {
    var arr = new StringBuilder("");

    if(a.length == 0) {
      return arr;
    } else if(a.length == 1) {
      arr = new StringBuilder(a);
      return arr;
    } else {
      var arr1 = new StringBuilder("");
      var arr2 = new StringBuilder("");

      for(i<-0 to (a.length/2-1)) {
        arr1 += a(i);
      }

      for(i<-(a.length/2) to (a.length - 1)) {
        arr2 += a(i);
      }

      arr = merge2(mergesort2(arr1.toString).toString, mergesort2(arr2.toString).toString);
    }

    return arr;
  }

  // Also Exercise 3: isPrimeDay and isFriday, using whatever date types are available in your language

  def isPrimeDay(y:Int,m:Int,d:Int) : Boolean = {
    if(divisors(d).size == 0) {
      return true;
    } else {
      return false;
    }
  }

  // learn how to use the java package from https://stackoverflow.com/questions/13647422/check-if-the-calendar-date-is-a-sunday
  def isFriday(y:Int,m:Int,d:Int) : Boolean = {
    var startDate = Calendar.getInstance();
    startDate.set(y,m-1,d);
    if (startDate.get(Calendar.DAY_OF_WEEK) == Calendar.FRIDAY) {
      return true;
    }

    return false;
  }

  ///////////////////////////////////////////////////////////////////////////
  ///////////////////////////////////////////////////////////////////////////
  /////////////////////       Examples            ///////////////////////////
  ///////////////////////////////////////////////////////////////////////////
  ///////////////////////////////////////////////////////////////////////////

  // sample program one 
  // I found a question online asking: How do you reverse an array in place in Java?
  // Scala is very simlar to Java, so I am going to use Scala to solve this question 
  // param: A Int arr
  // post: A Int arr in a reverse order
  def reverse(a:Array[Int]) : Array[Int] = {
    var arr = Array.empty[Int];

    if(a.size == 1) {
      arr = arr :+ a(0);
      return arr;
    }

    if(a.size != 0) {
      for(i<-(a.size-1) to 0 by -1) {
        arr = arr :+ a(i);
      }
    }

    return arr;
  }

  // A main function 
  def main(args: Array[String]) {
    // var arr = divisors(30)
    // for(i<-0 to (arr.size - 1)) {
    //   println(arr(i));
    // }

    // var arr = primes(100);
    // for(i<-0 to (arr.size - 1)) {
    //   println(arr(i));
    // }

    // var arr = Array("one","two","three");
    // var result = join(",", arr);
    // println(result);

    // var arr = pythagorean(30);
    // for(i<-0 to (arr.size - 1)) {
    //   println(arr(i));
    // }

    // var arr = merge(Array(4,5,7,8), Array(1,2,3,6,9));
    // for(i<-0 to (arr.size - 1)) {
    //   println(arr(i));
    // }

    // var s = merge2("aeguz", "ceptw");
    // println(s);

    // var arr = Array(6,2,4,8,9,5,3,1,7,10);
    // arr = mergesort(arr);

    // for(i<-0 to (arr.size - 1)) {
    //   println(arr(i));
    // }

    // var s = mergesort2("The quick brown fox jumps over the lazy dog.");
    // println(s);

    // var result = isPrimeDay(2018,6,23);
    // println(result);

    // var result = isFriday(2018,5,18);
    // println(result);

    // var arr = reverse(Array(1,2,3,4,5));
    // for(i<-0 to (arr.size - 1)) {
    //   println(arr(i));
    // }
  }
}