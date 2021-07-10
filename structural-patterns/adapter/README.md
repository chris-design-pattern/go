Adapter
---
## Intent

**Adapter**는 호환되지 않는 interface가 있는 object가 공동 작얼을 할 수 있도록 하는 구조적 design pattern입니다.

![](https://images.velog.io/images/chrishan/post/e2e7bcdf-f978-4e6b-a640-128b3ff33f6e/adapter-en-2x.png)

## Structure

#### Object adapter

이 구현에서는 개체 구성 원칙을 사용합니다. Adapter는 한 개체의 interface를 구현하고 다른 개체를 wrapping합니다. 널리 사용되는 모든 프로그래밍 언어로 구현할 수 있습니다.
![](https://images.velog.io/images/chrishan/post/a5567231-f0e0-4d49-aba4-084748672390/structure-object-adapter-indexed-2x.png)

1. **Client**는 프로그램의 기존 비지니스 로직을 포함하는 클래스입니다.
2. **Client interface**는 다른 클래스가 client 코드와 협업 할 수 있도록 따라야 하는 프로토콜을 설명합니다.
3. **Service**는 유용한 클래스(보통 타사 또는 레거시)입니다. Interface가 호환되지 않기 때문에 client는 이 클래스를 직접 사용할 수 없습니다.
4. **Adapter**는 Client와 Service 모두에서 작동 할 수 있는 클랙스입니다. 서비스 객체를 wrapping하면서 client interface를 구현합니다. Adapter는 adapter interface를 통해 client로부터 호출을 수신하고 이를 이해 할 수 있는 형식으로 wrapping된 service 개체에 대한 호출로 변환합니다.
5. Client 코드는 client interface를 통해 adapter와 함께 작동하는 한 구체적인 Adapter 클래스에 연결되지 않습니다. 덕분에 기존 client 코드를 손상시키지 않고 새로운 유형의 adapter를 프로그램에 도입 할 수 있습니다. 이는 service 클래스의 interface가 변경되거나 교체 될 때 유용 할 수 있습니다. Client 코드를 변경하지 않고 새 adapter 클래스를 만들 수 있습니다.
#### Class adapter

이 구현에서는 상속을 사용합니다. Adapter는 동시에 두 개체의 interface를 상속합니다. 이 접근 방식은 C++와 같이 다중 상속을 지원하는 프로그래밍 언어로만 구현할 수 있습니다.
![](https://images.velog.io/images/chrishan/post/6b478fcf-5947-406a-b070-fb2b7f977f0a/structure-class-adapter-indexed-2x.png)

1. Class adapter는 client와 service 모두에서 동작을 상속하므로 객체를 wrapping 할 필요가 없습니다. 재정의 된 method 내에서 적응이 발생합니다. 결과 adapter는 기존 client class 대산 사용할 수 있습니다.

## How to Implement

1. Interface가 호환되지 않는 class가 두개 이상 있는지 확인하십시오.
- 변경 할 수 없는 유용한 서비스 class(종종 타사, 레거시 또는 많은 기존 종속성 포함)
- Service class 사용으로 이점을 얻을 수 있는 하나 또는 여러 client class.
2. Client interface를 선언하고 client service와 통신하는 방법을 설명합니다.<br /><br />
3. Adapter class를 만들고 client interface를 따르도록 합니다. 지금은 모든 method를 비워 둡니다.<br /><br />
4. Service class에 대한 참조를 저장하려면 adapter class에 필드를 추가합니다. 일반적인 방법은 생성자를 통해 이 필드를 초기화하는 것이지만 method 호출할 때 adapter에 전달하는 것이 더 편리한 경우도 있습니다.<br /><br />
5. 하나씩 adapter class에서 client interface의 모든 method를 구현합니다. Adapter는 interface 또는 데이터 형식 변환만 처리하여 대부분의 실제 작업을 서비스 개체에 위임해야 합니다.<br /><br />
6. Client는 client interface를 통해 adapter를 사용해야 합니다. 이렇게 하면 client 코드에 영향을 주지 않고 adapter를 변경하거나 확장 할 수 있습니다.
## Pros and Cons

<span style="color:#346751;">O</span> 단일 책임 원칙. 프로그램의 기본 비즈니스 로직에서 interface 또는 데이터 변환 코드를 분리 할 수 있습니다.<br /><br />
<span style="color:#346751;">O</span> 개방/폐쇄 원칙. Client interface를 통해 adapter와 함께 작동하는 한 기존 client 코드를 손상시키지 않고 새로운 유형의 adapter를 프로그램에 도입 할 수 있습니다.<br /><br />
<span style="color:#C84B31;">X</span> 새로운 interface 및 class 세트를 도입해야 하므로 코드의 전반적인 복잡성이 증가합니다. 때로는 나머지 코드와 일치하도록 service class를 변경하는 것이 더 간단합니다.