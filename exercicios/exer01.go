package main

import(
	"fmt"
	"strconv"
	//"strings"
	"regexp"
)

func GetSum(a, b int) int {
	soma := 0
	
	  if a == b {
		soma = a
	   
	  } else {
		 soma = a + b
	  }
  
	return soma
  }

  func GetSum1(a, b int) int {
	soma := 0
	
	  if a == b {
		soma = a
	  }
	
	if a < b {
	   for i:= a; i <= b; i++ {
		 soma = soma + i
	}
	  }
	
	  if b < a {
	   for i:= b; i <= a; i++ {
		 soma = soma + i
	}
	  }
  
	return soma
  }

func Feast(beast string, dish string) bool {
	a := string(dish[0])
	b := string(beast[0])
	
	a1:= string(dish[len(dish) -1])
	b1:= string(beast[len(beast) -1])
	
	if a == b && a1 == b1 {
	  return true
	} else {
	  return false
	}
  }

func FakeBin2(x string) (result string) {
	return regexp.MustCompile("([5-9])").ReplaceAllString(regexp.MustCompile("([0-4])").ReplaceAllString(x, "0"), "1")
}

func sum(x []int) []int {
	y := make([]int, len(x))
	for i, valor := range x {
		y[i] = valor *2
		
	}
	return y
}

func CheckForFactor(base int, factor int) bool {
    result := base%factor
  if result == 0{
    return true
  }
  return false
}

func Points(games []string) (res int) {
	for _,v:=range games {
	if v[0]>v[2] {res +=3}
	if v[0]==v[2] {res +=1}
  }
	return
  }

func Invert(arr []int) []int {
	y := make([]int, len(arr))
  for i, valor := range arr {
 
	y[i] = valor * -1

  } 
  return y

}

func additiveInverse(numbers []int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = -num
    }
    return result
}

func QuarterOf(month int) int {
	
	switch month {
	  case 1,2,3:
	  return 1
	  case 4,5,6:
	  return 2
	  case 7,8,9:
	  return 3
	  case 10,11,12:
	  return 4
	default:
		return 0
	}
  }

  func SquareSum(numbers []int) int {
	soma := 0
	  for i := 0; i < len(numbers); i++ {
		soma = soma + (numbers[i] * numbers[i])
	  }
	return soma
  }

  func SquareSum1(numbers []int) int {
    res := 0
    for _ ,v := range numbers {
      res += v*v
    }
    return res
}

func countSheep(num int) string {
	result := ""
	for i := 1; i <=num; i++ {
	  result = result +  strconv.Itoa(i) + "sheep..."
	} 
	return result
  }

  func NumberToString(n int) string {
	convert := strconv.Itoa(n)
	return convert
  }

  func NumberToString1(n int) string {
	return fmt.Sprintf("%d", n)
  }

  func RepeatStr(repetitions int, value string) string {
// 	import "strings"

// func RepeatStr(repititions int, value string) string {
//   return strings.Repeat(value, repititions)
// }
	result := " "
	for i := 0; i < repetitions; i++ {
	  result = result + value
	}
	return result
  }


  func Digitize(n int) []int {
	str := strconv.Itoa(n)
	arr := make([]int, len(str))
	for k, v := range str {
	  arr[len(str)-1-k] = int(v - '0')
	}
  
	return arr
  }
  //o primeiro valor vai para o ultimo indice, o segundo valor vai para o penultimo indice...
  // primeira interação i vale 0, na segunda i vale 1, na terceira i vale 3...
  //num array de tamanho 5 
  //tamnho -1 -0 vai ser o ultimo indice, 4 que recebe 
  //tamnho -1 -1 vai ser o penultimo indice
  //tamanho -1 -2 vai ser o antepenultimo indice

  func SmallestIntegerFinder(numbers []int) int {

	menor := numbers[0]
	
	for _, v := range numbers {
	  if v < menor{
		menor = v
	  }
	}
  
	return menor 
  }
  func FakeBin(x string) string {
	b:=[]byte(x)
	for i,v:=range b{
	  if v<'5' {b[i]='0'}else{b[i]='1'}
	}
	return string(b)
  }

  func FakeBin1(x string) (result string) {
    for _, char := range x {
        if char < '5' {
            result += "0"
        } else {
            result += "1"
        }
    }
    return
}
func NbYear(p0 int, percent float64, aug int, p int) int {
  
	years := 0
	
	for p0 < p {
	  increase := int(float64(p0) * (percent / 100))
		  p0 += increase + aug
		  years++
	}
   return years
  }

  func Grow(arr []int) int{
  
	result := 1 
	
	for _, v := range arr { 
	  result = result * v
	}
	return result
  }
  
  func lastPetalPhrase(petals int) string {
    phrases := []string{"I love you", "a little", "a lot", "passionately", "madly", "not at all"}

    // Calculate the index of the last phrase
    lastPhraseIndex := (petals - 1) % len(phrases)

    return phrases[lastPhraseIndex]
}

func main() {

	r3:= GetSum1(2,6)
	fmt.Println("r3:", r3)

	r2:= FakeBin("5632332")
	fmt.Println("r2" + r2)

	r1 := Digitize(2345)
	fmt.Println("array invertido", r1)

	r := lastPetalPhrase(14)
	fmt.Println(r)

	res:= RepeatStr(3, "txt")
	fmt.Println(res)

	tipoSlice := []int{11, 12, -13, 14, 15}

	resp := Invert(tipoSlice)
	fmt.Println("Funcao inverter", resp)

	res1 := SquareSum(tipoSlice)
	fmt.Println("funcao soma dos quadrado", res1)

	result := sum(tipoSlice)
	fmt.Println(result)

	ovelha := countSheep(4)
	fmt.Println(ovelha)
}
