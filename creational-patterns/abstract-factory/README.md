Abstract Factory
---
## Intent

**Abstract Factory**는 구체적인 클래스를 지정하지 않고 관련성을 갖는 객체들의 집합을 생성하거나 서로 독립적인 객체들의 집합을 생성할 수 있는 인터페이스를 제공하는 패턴이다.

![](https://images.velog.io/images/chrishan/post/64b62edd-e018-409d-a397-d017cf95a317/abstract-factory-en-2x.png)

## Structure
![](https://images.velog.io/images/chrishan/post/85b01a09-3c9e-42b2-8f52-5ce4611d84ab/structure-indexed-2x.png)
1. **Abstract Products**는 제품군을 구성하는 별개의 관련 제품 세트에 대한 인터페이스를 선언합니다.<br /><br />
2. **Concrete Products**는 변형별로 그룹화된 추상 제품의 다양한 구현입니다. 각 Abstract Product(의자/소파)은 주어진 모든 변형(Victorian/Modern)으로 구현되어야 합니다.<br /><br />
3. **Abstract Factory** 인터페이스는 각각의 추상 제품을 생성하기 위한 메소드 세트를 선언합니다.<br /><br />
4. **Concrete Factories**는 추상 팩토리의 생성 방법을 구현합니다. 각 concrete factory는 제품의 특정 변형에 해당하며 해당 제품 변형만 생성합니다.<br /><br />
5. Concrete Factory는 concrete product를 인스턴스화하지만 생성 방법의 signature는 해당 _abstract_ product를 반환해야 합니다. 이렇게 하면 팩토리를 사용하는 클라이언트 코드가 팩토리에서 가져온 제품의 특정 변형에 연결되지 않습니다. **Client**는 추상 인터페이스를 통해 개체와 통신하는 한 구체적인 공장/제품 변형으로 작업할 수 있습니다.

## How to Implement

1. 고유한 제품 유형과 이러한 제품의 변형 매트릭스를 매핑 한다.<br /><br />
2. 모든 제품 유형에 대한 Abstract Interface를 선언한다. 그런 다음 모든 구체적인 제품 Class가 이러한 Interface를 구현하도록 한다.<br /><br />
3. 모든 추상 제품에 대한 일련의 생성 방법을 사용하여 Abstract Factory Interface를 선언한다.<br /><br />
4. 각 제품 변형에 대해 하나씩 구체적인 Factory Class 집합을 구현한다.<br /><br />
5. App 어딘가에 Factory initialization code를 만든다. Application 구성이나 현재 환경에 따라 Concrete factory class 중 하나를 인스턴스화 해야 한다. 이 Factory object를 제품을 구성하는 모든 Class에 전달한다.<br /><br />
6. Code를 스캔하여 제품 생성자에 대한 모든 직접 호출을 찾는다. Factory object에 대한 적절한 Creation method에 대한 호출을 대체한다.<br /><br />

## Pros and Cons

<span style="color:green;">O</span>. Factory에서 얻어지는 제품이 서로 호환되는지 확인 할 수 있다.<br /><br />
<span style="color:green;">O</span>. 구체적인 제품과 Client code간의 긴밀한 결합을 피한다.<br /><br />
<span style="color:green;">O</span>. Single Responsibility Principle. 제품 생성 코드를 한 곳으로 추출하여 코드를 더 쉽게 지원할 수 있다.<br /><br />
<span style="color:green;">O</span>. Open/Closed Principle. 기존 Client code를 손상시키지 않고 새로운 제품 변형을 도입 할 수 있다.<br /><br />
<span style="color:red;">X</span>. 패턴과 함께 새로운 Interface와 Class가 도입 되었기 때문에 코드는 생각보다 복잡해 질 수 있다.<br /><br />

