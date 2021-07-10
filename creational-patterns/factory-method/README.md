Factory Method
---
## Intent
**Factory Method**는 슈퍼 클래스에서 개체를 생성하기 위한 인터페이스를 제공하지만
하위 클래스가 생성 될 개체의 유형을 변경할 수 있도록 하는 `Creational Design Pattern`이다.
![img.png](https://refactoring.guru/images/patterns/content/factory-method/factory-method-en-2x.png)

## Structure
![](https://media.vlpt.us/images/chrishan/post/c5d690ef-e770-4f21-a3f2-1150f104c7c8/Screen%20Shot%202021-05-09%20at%2013.21.19.png)

## How to implement
1. 모든 제품이 동일한 인터페이스를 따르도록합니다. 이 인터페이스는 모든 제품에서 의미있는 메서드를 선언해야합니다.<br /><br />
2. 생성자 클래스 내부에 빈 Factory Method를 추가합니다. 메서드의 반환 유형은 공통 제품 인터페이스와 일치해야합니다.<br /><br />
3. 작성자의 코드에서 제품 생성자에 대한 모든 참조를 찾습니다. 하나씩 Factory Method에 대한 호출로 대체하고 제품 생성 코드를 Factory Method로 추출합니다.<br /><br />Returned product의 유형을 제어하기 위해 Factory Method에 임시 매개 변수를 추가해야 할 수 있습니다.<br /><br />이 시점에서 Factory Method의 코드는 매우 추하게 보일 수 있습니다. 인스턴스화 할 제품 클래스를 선택하는 큰 스위치 연산자가 있을 수 있습니다. 하지만 걱정하지 마세요. 곧 수정하겠습니다.<br /><br />
4. 이제 Factory Method에 나열된 각 제품 유형에 대한 작성자 하위 클래스 세트를 작성하십시오. 하위 클래스에서 Factory Method를 재정의하고 기본 메서드에서 적절한 구성 코드 비트를 추출합니다.<br /><br />
5. 제품 유형이 너무 많고 모든 유형에 대해 하위 클래스를 만드는 것이 타당하지 않은 경우 하위 클래스의 기본 클래스에서 제어 매개 변수를 재사용 할 수 있습니다.<br /><br />예를 들어, 다음과 같은 클래스 계층이 있다고 가정 해 봅시다. 두 개의 하위 클래스가있는 기본`Mail` 클래스 :`AirMail` 및`GroundMail`; `Transport` 클래스는`Plane`,`Truck` 및`Train`입니다. `AirMail` 클래스는`Plane` 객체 만 사용하지만`GroundMail`은`Truck` 및`Train` 객체 모두에서 작동 할 수 있습니다. 두 경우를 모두 처리하기 위해 새 하위 클래스 (예 : 'TrainMail')를 만들 수 있지만 다른 옵션이 있습니다. 클라이언트 코드는`GroundMail` 클래스의 Factory Method에 인수를 전달하여 수신하려는 제품을 제어 할 수 있습니다.<br /><br />
6. 모든 추출 후 기본 Factory Method가 비어있는 경우 추상화로 만들 수 있습니다. 남은 것이 있으면 이를 메서드의 기본 동작으로 만들 수 있습니다.

## Pros and Cons
<span style="color:green;">O</span>. Creator와 concrete products간의 coupling을 피합니다.<br /><br />
<span style="color:green;">O</span>. 단일 책임 원칙. 제품 생성 코드를 프로그램의 한 위치로 이동하여 코드를 더 쉽게 지원할 수 있습니다.<br /><br />
<span style="color:green;">O</span>. Open/Closed Principle. 기존 클라이언트 코드를 손상시키지 않고 새로운 유형의 제품을 프로그램에 도입 할 수 있습니다.<br /><br />
<span style="color:red;">X</span>. 패턴을 구현하기 위해 많은 새로운 서브 클래스를 도입해야하므로 코드가 더 복잡해질 수 있습니다. 가장 좋은 시나리오는 creator classes의 기존 계층 구조에 패턴을 도입하는 경우입니다.<br /><br />