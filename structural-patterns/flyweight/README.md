Flyweight
---
## Indent
**Flyweight**는 각 객체의 모든 데이터를 유지하는 대신 여러 객체간에 공통 상태 부분을 공유하여 사용 가능한 RAM 양에 더 많은 객체를 맞출 수 있는 Structural design pattern입니다.
![](https://images.velog.io/images/chrishan/post/e9e6eb2c-763d-497d-8ae8-f146254ef560/flyweight-2x.png)

## Structure
![](https://images.velog.io/images/chrishan/post/76d79992-8dba-4044-8d10-5576038a0945/structure-indexed-2x.png)
1. Flyweight pattern은 단지 최적화일 뿐입니다. 적용하기 전에 프로그램이 동시에 메모리에 유사한 개체를 대량으로 보유하는 것과 관련된 RAM 소비 문제가 있는지 확인하십시오. 이 문제가 다른 의미 있는 방법으로 해결되지 않도록 하십시오.<br /><br />
2. **Flyweight** 클래스에는 여러 객체 간에 공유할 수 있는 original 객체의 상태 부분이 포함되어 있습니다. 동일한 flyweight 객체가 다양한 상황에서 사용될 수 있습니다. Flyweight 내부에 저장된 상태를 _intrinsic_이라고 합니다. Flayweight의 메소드에 전달된 상태를 _extrinsic_이라고 합니다.<br /><br />
3. **Context** 클래스는 모든 원본 객체에서 고유한 외부 상태를 포함합니다. 컨텍스트가 flyweight 객체 중 하나와 쌍을 이루면 original 객체의 전체 상태를 나타냅니다.<br /><br />
4. 일반적으로 original 객체의 동작은 flyweight 클래스에 남아 있습니다. 이 경우 flyweight의 메소드를 호출하는 사람은 외부 상태의 적절한 비트도 메소드의 매개변수에 전달해야 합니다. 다른 한편으로, 동작은 연결된 flyweight를 단순히 데이터 객체로 사용하는 컨텍스트 클래스로 이동할 수 있습니다.<br /><br />
5. **Client**는 flyweight의 외부 상태를 계산하거나 저장합니다. 클라이언트의 관점에서 flyweight는 일부 컨텍스트 데이터를 메서드의 매개변수에 전달하여 런타임에 구성할 수 있는 템플릿 개체입니다.<br /><br />
6. **Flyweight Factory**는 기존 flyweight pool을 관리합니다. Factory에서 client는 flyweight를 직접 만들지 않습니다. 대신에 원하는 flyweight의 고유 상태 비트를 전달하여 팩토리를 호출합니다. 팩토리는 이전에 생성된 flyweight를 살펴보고 검색 기준과 일치하는 기존 flyweight를 반환하거나 아무것도 발견되지 않으면 새 flyweight를 생성합니다.
## How to Implement
1. Flyweight가 될 클래스의 필드를 두 부분으로 나눕니다.
- the intrinsic state: 많은 개체에 걸쳐 복제된 변경되지 않는 데이터를 포함하는 필드
- the extrinsic state: 각 객체에 고유한 컨텍스트 데이터를 포함하는 필드
2. 클래스의 고유 상태를 나타내는 필드를 그대로 두되, 변경할 수 없는지 확인하십시오. 생성자 내부에서만 초기 값을 가져와야 합니다.<br /><br />
3. 외부 상태의 필드를 사용하는 방법을 살펴보십시오. 메소드에 사용된 각 필드에 대해 새 매개변수를 도입하고 필드 대신 사용하십시오.<br /><br />
4. 선택적으로 flyweight pool을 관리하기 위한 팩토리 클래스를 생성합니다. 새 flyweight를 만들기 전에 기존 flyweight를 확인해야 합니다. Factory가 설치되면 client는 factory를 통해서만 flyweight를 요청해야 합니다. 그들은 factory에 고유 상태를 전달하여 원하는 flyweight를 설명해야 합니다.<br /><br />
5. 클라이언트는 flyweight 객체의 메서드를 호출할 수 있도록 외부 상태(컨텍스트) 값을 저장하거나 계산해야 합니다. 편의상, 플라이웨이트 참조 필드와 함께 외부 상태는 별도의 컨텍스트 클래스로 이동될 수 있다.
## Pros and Cons
<span style="color:green;">O</span>. 프로그램에 유사한 개체가 많이 있다고 가정하면 많은 RAM을 절약할 수 있습니다.<br /><br />
<span style="color:red;">X</span>. 누군가가 플라이웨이트 메서드를 호출할 때마다 컨텍스트 데이터의 일부를 다시 계산해야 하는 경우 CPU 주기보다 RAM을 거래할 수 있습니다.<br /><br />
<span style="color:red;">X</span>. 코드가 훨씬 더 복잡해집니다. 새로운 팀원들은 항상 엔티티의 상태가 왜 그런 식으로 분리되었는지 궁금해 할 것입니다.