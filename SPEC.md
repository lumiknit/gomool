# Lang Spec (Editing)

## Scratch Pad

```
# <- 줄 주석
# `A #=> B` 는 gomool 코드 A가 golang 코드 B로 컴파일 된다는 뜻.

# ;이 2개 이상 연달아 있으면 newline이랑 동일.
# ;이 하나인 것은 별도의 separator임.

# -- 패키지

package PKG #=> package PKG
import "module" #=> import "module"
import alias "module" #=> import (alias "module")

# -- 식별자 명명
# 만약 식별자가 그대로 go에서 쓰일 수 있다면 그대로 사용.
# 그렇지 않은 식별자(e.g. break 등의 키워드나 operator, unicode) 역시 사용할
# 수는 있는데, 이 경우에 해당하는 경우에는 무조건 gommol id encoding rule을 따름.
# Gomool ID Encoding Rule:
# - 처음에는 `gx_`나 `Gx_`로 시작함. (Priv/Publ에 따라 적절히 사용)
# - 원래 Go에서 허용되는 문자들은 그대로 사용
# - `_`는 `__`로 사용.
# - ASCII 범위의 문자들은 대응되는 `_XX` 형태의 코드가 있어서 이로 변환.
#   XX에는 hexadecimal characters를 제외한 알파벳이 들어감.
# - 그 외의 유니코드는 `_HHHH`처럼 hex로 나타내게 됨.
# 기본적으로 public/private는 golang의 규칙과 같이 첫글자가 대문자인 경우에만
# public인데, 맨 앞에 `$`를 붙여서 강제로 public으로 만들 수 있음.
# 다만, 이렇게 만들면 강제로 Gomool ID Encoding Rule을 따르게 됨.
# 만약 이름이 `gx_`나 `Gx_`로 시작해도 강제로 Gomool ID Encoding Rule을 따름.
# 만약 이름이 keyword라면 마찬가지로 강제로 위 encoding을 따름.
# 만약에 연산자 등 `_` 외의 특수문자를 포함하는 이름을 만들고 싶으면 backtick으로
# 감싸서 사용.

abcdefg = "priv"
Abcdefg = "publ"
$abcdefg = "publ" #=> var Gx_abcdefg = "publ"
`+-` = "op" #=> var gx__pl_mn = "op"; which is private
$`+-` = "op" #=> var Gx__pl_mn = "op"; which is public

# -- 연산자
# 만약 Gomool에서 연산자로 사용가능한 문자로만 이루어진 단어?가 있으면
# 연산자로 사용 가능.
# 만약 public, private가 각각 정의가 되어 있으면 private를 우선시 함.
# 무조건 public을 써야하면 연산자 앞에 $를 붙이면 됨.
# 단항연산은 연산자 맨 뒤에 `_`를 붙인 것으로 취급함. e.g. `+_`를 정의하면 `+a` 같이
# 사용할 때 접근됨.

# -- 상수
# 기본적으로는 Golang의 것을 그대로 사용하면 됨.
# 대신 backtick string (or raw string)은 backtick이 identifier로 사용되므로
# 대신에 ''처럼 multiple quote를 사용하면 됨.

# -- 변수 정의 / 대입
# 다른 언어처럼 id = expr 형태로 사용하면 충분.
# 변수가 처음 사용되는 시점에서 `var a`를 따로 정의하게 됨.
hello = 42 #=> var hello = 42
hello = 50 #=> hello = 50
# 흔히 쓰이는 패턴처럼 여러 변수를 한번에 대입하는 것도 가능.
val, ok = somethingMayGoWrong(boom)

# -- 함수 정의 / 대입
# 기본적으로 value level에서 func에 대응되는 키워드는 =>임.
f = (x, y) => y #=> func f(x, y) { return y }
# 다만, 함수를 변수에 저장하는게 아니라 함수 자체를 사용한다면,
f(x, y) = y
# 같은 형태도 허용함.
# 다만, 함수 내에서는 f(x, y) = y를 f = (x, y) => y의 syntatic sugar로 보지만
# top level에서는 f(x, y) = y는 func f(x, y) {...}, f = (x, y) => y는
# var f = func(x, y) {...}으로 컴파일함.
# 메소드를 정의하고 싶으면
obj.f(x, y) = y #=> func (obj any) f(x, y) { y }

# -- 타입 지정
# 만약 변수나 함수를 만들 때 타입을 지정하고 싶으면 정의하기 전에 미리 id :: Type
# 형태로 타입을 지정해주어야 함.
# 예를 들어서
f: (int, int) => int
f(x, y) = y
# 라고 하면 컴파일 시
func f(x int, y int) int { return y }
# 와 같이 타입 정보를 합쳐서 넣게 됨.
# 참고로 type은 golang의 것을 거의 그대로 사용하되, 함수의 경우
f: (a1, a2, ...) => ret
f: type.(a1, a2, ...) => ret
# 형태로 사용

# -- Function Call, Type Conversion, Type Assertion, Indexing
# Golang과 동일
# 단, function call, type conversion은 인자가 1개면 괄호를 안 해도 됨.
f x

# -- 타입 정의
# :=를 사용하면 type을 정의함.
myInt := int32
# type level에서 struct는 @{}만으로 감싸기만 해도 됨.
# 타입 지정을 할 때 ::를 사용하고, 세미콜론으로 구분. 만약 :: 지정 없으면 타입이 옴.
# 타입이 오는 경우에 이름은 _1, _2, ...처럼 작명됨.
box := @{ int; a, b, c :: int }
# type level에서 interface는 @.{}로 감싸기만 하면 됨.
# 다만 함수 타입을 지정할 때 메소드가 아닌 일반 함수 형식으로 해야됨.
addible := @.{ add :: (int) => int; zero :: int }

# -- 블록
# gomool도 파이썬처럼 indentation된 block을 block으로 판단함.
# 예를 들어서 multiline function은
f(x, y) =
  a = 20
  b = 30
  return a + b + x + y
# 처럼 만들 수 있음. 즉, 그 줄에서 필요한 내용이 안 나오고 다음 줄로 넘기게 되면
# 인덴트한 것에 따라 적절히 처리하게 됨.

# -- 조건문
# golang style if
if val
  print 20
else if err, v = boom(); err
else
  boom

# switch
switch val
  a -> 20
  b -> 30
  _ -> default-case

# cond-switch
switch _
  f() > 0 -> 20
  b -> 30

# golang style for
for init; cond; inc
  int

# -- goto, return, break, continue, defer, chan, go, etc.
# golang처럼 하면 됨.
``label``
  goto label

# select
ch <- 20
select
  a := <-ch ->
    print 30

go f(x)
go =>
  print(20)
#=> go func(){ print(20) }()
```

## Example code

```
package main

import "fmt"
       sub "github.com/abc/sub"

main() =
  fmt.Println("Hello, World!")

fibo: (int) => int
fibo(n) = switch _
  n <= 1 -> n
  _      -> fibo(n - 1) + fibo(n - 1)

fiboFast: (int) => int
fiboFast(n) =
  a, b = 0, 1
  for i = 0; i < n; i += 1
    a, b = b, a + b
  a

typedFunc = (x: int, y: float) =>:(float)
  float(x) + y

makeGoroutines: (int) => _
makeGoroutines(n) =
  wg = sync.WaitGroup()
  for i = 0; i < n; i += 1
    wg.Add()
    go =>
      sync.Sleep(10)
      wg.Done()
  wg.Wait()

Cons := @{ any; any }
Nil := @{}

Len_able := @.{ Len: () => int }

Cons.Len: () => int
x.Len() = 1 + Len_able(x).Len()

Nil.Len: () => int
x.Len() = 0
```
