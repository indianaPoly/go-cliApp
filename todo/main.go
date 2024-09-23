package main

/**
* bufio : 버퍼링된 입출력을 제공하는 패키지 (데이터를 버퍼에 모아 처리)
* encoding/json : 구조체를 JSON(인코딩)으로 변환 하거나 혹은 JSON을 구조체(디코딩)로 변환
* fmt : 다양한 형식의 데이터를 출력하거나 입력
* os : 운영체제와 상호작용 하기 위한 기능, 파일 읽기/쓰기, 환경 변수 접근, 명령줄 인자 처리
* strings : 문자열 조작과 관련된 유틸리티 함수 제공
 */
import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// json: "" 은 태그라고 불림
// 인코딩 디코딩시 지정되는 메타데이터터
// 구조체 선언 (앞 글자 대문자)
type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

// 변수 선언 및 타입 선언언
var todos []Todo

// 할일 목록 보기 설정
func listTodos() {
	if len(todos) == 0 {
		fmt.Println("할 일이 없습니다,")
		return
	}

	// range 반복문에서는 첫 번째 요소는 인덱스
	// 두 번째 요소는 데이터를 의미함.
	for _, todo := range todos {
		status := "[ ]"

		if todo.Completed {
			status = "[X]"
		}
		fmt.Printf("%d. %s %s\n", todo.ID, status, todo.Task)
	}
}

// 할 일 목록 추가하기
func addTodo() {
	// Go에서 표준 입력(데이터)를 읽기 한 리더 객체 생성
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("새로운 할 일 : ")
	// \n 문자열이 나올 때 까지 문자열을 계속 읽음
	task, _ := reader.ReadString('\n')
	// 문자열 앞 뒤에 있는 공백 문자를 제거거
	task = strings.TrimSpace(task)

	id := len(todos) + 1
	todos = append(todos, Todo{ID: id, Task: task, Completed: false})

	fmt.Println("할 일이 추가되었습니다.")
}

// 완료 한 일에 대해서 처리리
func completeTodo() {
	var id int

	fmt.Print("처리 완료할 ID: ")
	_, err := fmt.Scanln(&id)

	if err != nil {
		fmt.Println("입력 오류 : ", err)
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = true
			fmt.Println("할 일에 대해서 표시 완료")
			return
		}
	}

	fmt.Println("해당 ID의 할 일을 찾을 수 없습니다.")
}

// 삭제
func deleteTodo() {
	var id int
	fmt.Print("삭제할 ID: ")
	_, err := fmt.Scanln(&id)

	if err != nil {
		fmt.Println("ID 에러 : ", err)
		return
	}

	found := false
	for i, todo := range todos {
		if todo.ID == id {
			// todos i 번째 전까지만 선택을 하여 하나의 리스트로 만든다.
			// 그리고 todos i + 1 부터 리스트의 요소를 분리하여 1번의 리스트에 추가한다. 그리고 이를 todos로 새로이 반환한다.
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("할 일이 삭제되었습니다.")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("해당 ID의 할 일을 찾을 수 없습니다.")
	}
}

// 저장하는 함수
func saveTodos() {
	// 인코딩
	file, _ := json.MarshalIndent(todos, "", " ")
	// 0644는 파일 권한 부여 = 0400 + 0200 + 040 + 04
	// 소유자는 읽기, 쓰기 가능
	// 그룹과 기타 사용자는 읽기 권한만 존재
	_ = os.WriteFile("todos.json", file, 0644)
}

func main() {
	// 파일을 불러와 정보를 전역변수에 저장
	file, err := os.ReadFile("todos.json")

	// nil이 아닌 다른 에러가 반환이 된다면 리턴 처리리
	if err != nil {
		fmt.Println("파일 읽기 오류:", err)
		return
	}

	// Unmarshal (Json을 구조체로 디코딩)
	// file에 담겨 있는 정보를 todos 주소로 하여 저장
	if err := json.Unmarshal(file, &todos); err != nil {
		fmt.Println("JSON 파싱 중 에러 발생:", err)
		return
	}

	// 무한 반문 실행
	for {
		fmt.Println("\n할 일 관리 애플리케이션")
		fmt.Println("1. 할 일 목록 보기")
		fmt.Println("2. 새 할 일 추가")
		fmt.Println("3. 할 일 완료 표시")
		fmt.Println("4. 할 일 삭제")
		fmt.Println("5. 종료")
		fmt.Print("선택: ")

		var choiceNumber int
		// 입력 번호 저장
		_, err := fmt.Scanln(&choiceNumber)

		if err != nil {
			fmt.Println("입력에 오류가 있습니다.")
		}

		switch choiceNumber {
		case 1:
			listTodos()
		case 2:
			addTodo()
		case 3:
			completeTodo()
		case 4:
			deleteTodo()
		case 5:
			saveTodos()
			fmt.Println("애플리케이션이 종료되었습니다.")
			return
		default:
			fmt.Println("잘못된 선택입니다. 다시 선택해주세요.")
		}
	}
}
