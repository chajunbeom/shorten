# Shorten URL

너무 길거나 기억하기 어려운 URL들을 짧게 변경해 주며\
변경한 URL로 접근시 원본 사이트로 리다이렉트 해주는 웹어플리케이션 입니다.

## 선행작업 

### 1. 환경
> OSX,Linux

### 2. Go 설치
최신버전의 Go가 필요합니다.

>Go install\
https://golang.org/dl/

### 3. Go 환경 설정 및 소스 다운로드 
Go 소스코드를 실행하기 위해선 bin,pkg,src 폴더를 하위경로로 가지고 있는 Path를 환경변수 GOPATH로 설정해야 합니다.

```
$ mkdir project && cd project
$ mkdir bin pkg src
$ export GOPATH=$(pwd)
```

소스 코드는 src 폴더아래에 있어야 합니다.

```
$ cd $GOPATH/src/
$ git clone https://github.com/chajunbeom/shorten.git
```

## 빌드하기

$GOPATH/src 경로에서 Makefile을 실행하여 빌드합니다.
```
$ cd $GOPATH/src/shorten
$ make
```
빌드가 성공하면 bin 폴더가 생성 되며 다음 명령어로 어플리케이션을 실행합니다.

```
$ make run

어플리케이션 중지
$ make stop
```
정상적으로 실행 됬다면 http://localhost:23000  (Default)으로 접속합니다.

## 단위 테스트

Go 언어는 src/project/(pakage) 아래에 *_test.go 으로 된 단위 테스트 코드를 명령어로 실행 시킬 수 있습니다.

#### DB Test
```
$ cd $GOPATH/src/shorten/models
$ go test
 ... 생략
```
#### ShortenKey Test
```
$ cd $GOPATH/src/shorten/utils
$ go test
 ... 생략
```
## 문제 해결 전략

### 0. API
#### Redirect API
- Description :
생성한 shorten_key로 접근시 원본 url로 리다이렉팅 해주는 API
- Method : GET
- URI : /:shorten_key
- Parameter:
```
shorten_key를 uri path에 입력
```
- Response: Origin 주소로 리다이렉트

#### GetShortenURL API
- Description :
원본 URL을 shorten_key로 변환 해주는 API
- Method : POST
- URI : /convert
- ContentType: application/x-www-form-urlencoded
- Parameter:
```
key: origin_url
type: string(formData)
decription: 원본 url
```
- Response:
```
{
    "result": "OK", // 결과 코드
    "data": {
        "origin": "kakaopay.com", // 입력한 url
        "shorten_url": "DYg9Z1" // 변환한 shorten key
    }
}
```
### 1. ShortenKey 생성 알고리즘

조건 사항
1. 8자 이내의 shortenkey 생성

> 최대 길이 8, 문자열로만 구성된 랜덤키를 생성하기 위해서 Base64 encoding을 선택, 표준 Base64 Mapping table에서 +,/ 문자는 URL 표기법과 혼동 될 수 있으므로 -,_ 문자로 치환 합니다.

2. 동일한 url은 동일한 shortenkey 생성

> Base64로 인코딩할 유니크한 byte스트림이 필요 합니다.\
> CRC32 알고리즘으로 hash를 생성합니다. CRC32 를 선택한 이유는
> 1. 네트워크 패킷 에러 검출용으로 사용하는 만큼 충돌 확률이 적으며 속도도 빠릅니다.
>2. 또 8글자의 base64 string을 생성하려면 최대 6byte hash value가 필요한데 CRC32 코드는 4byte 이므로 용량 관점에서도 적합하다고 생각하였습니다.
>
> 따라서 본 웹어플리케이션에서 생성하는 shorten hash value는 다음과 같은 구조를 가지고 있습니다.\
> =========== 6byte =============\
> [ 2byte: 충돌값 ][ 4byte: CRC32 code ]\
> 첫 2byte는 0 으로 초기화 CRC32 code 충돌시 '충돌값' 1씩 증가\
> 동일한 CRC32 코드값, 최대 65535개의 hash 생성가능

3. 생성알고리즘 library사용 금지

> Base64 인코더 구현 (맵핑 테이블 수정 +/ 문자 => -_ 문자로 치환)\
> CRC32 Hash생성기 구현 (표준)

### 2. Data Schema
DataBase를 사용하진 않으며 In memory database로 동작 합니다.\
- primary_key : shortenURL,originURL
- value : struct{ shorten, origin }

### 3. Etc.
Shortenkey 생성 요청시 간단한 Origin URL 정규식 검사로 무분별한 Shortenkey 발급 억제
### 4. 개선사항
- Hash 충돌시 최악의 경우 65535번 동일한 로직을 수행
- MemoryDB 변경 필요
- 기타 등등...