Proxy
---
## Intent
## Intent
**Proxy**는 다른 개체에 대한 대체 또는 자리 표시자를 제공할 수 있는 구조적 디자인 패턴입니다. 프록시는 원래 개체에 대한 액세스를 제어하므로 요청이 원래 개체에 전달되기 전이나 후에 수행할 수 있습니다.

![](https://images.velog.io/images/chrishan/post/752f5439-089e-467e-b399-11c9a5ba5a45/proxy-2x.png)

## Structure
![](https://images.velog.io/images/chrishan/post/ea5c84ec-d8ec-4e17-8b1f-873b319f981a/structure-indexed-2x.png)
1. **Service Interface**는 서비스의 인터페이스를 선언합니다. Proxy는 서비스 개체로 위장할 수 있도록 이 인터페이스를 따라야 합니다.<br />
2. **Service**는 몇 가지 유용한 비즈니스 로직을 제공하는 클래스입니다.<br />
3. **Proxy** 클래스에는 서비스 개체를 가리키는 참조 필드가 있습니다. 프록시가 처리(e.g., lazy initialization, logging, access control, caching, etc.)를 완료한 후 요청을 서비스 개체에 전달합니다.<br />일반적으로 프록시는 서비스 개체의 전체 수명 주기를 관리합니다.<br />
4. **Client**는 동일한 인터페이스를 통해 서비스와 프록시 모두에서 작업해야 합니다. 이렇게 하면 서비스 개체가 필요한 모든 코드에 프록시를 전달할 수 있습니다.

## How to Implement
1. 기존 서비스 인터페이스가 없는 경우 프록시와 서비스 객체를 교환할 수 있도록 생성합니다. 서비스 클래스에서 인터페이스를 추출하는 것이 항상 가능한 것은 아닙니다. 해당 인터페이스를 사용하려면 서비스의 모든 클라이언트를 변경해야 하기 때문입니다. 플랜 B는 프록시를 서비스 클래스의 하위 클래스로 만드는 것이며, 이렇게 하면 서비스의 인터페이스를 상속합니다.<br /><br />
2. 프록시 클래스를 만듭니다. 서비스에 대한 참조를 저장하기 위한 필드가 있어야 합니다. 일반적으로 프록시는 서비스의 전체 수명 주기를 생성하고 관리합니다. 드문 경우지만 클라이언트가 생성자를 통해 서비스를 프록시에 전달합니다.<br /><br />
3. 목적에 따라 프록시 메소드를 구현하십시오. 대부분의 경우 일부 작업을 수행한 후 프록시는 작업을 서비스 개체에 위임해야 합니다.<br /><br />
4. 클라이언트가 프록시를 받을지 실제 서비스를 받을지 결정하는 생성 방법을 도입하는 것을 고려하십시오. 이것은 프록시 클래스의 간단한 정적 메서드 또는 본격적인 팩토리 메서드가 될 수 있습니다.<br /><br />
5. 서비스 개체에 대해 지연 초기화 구현을 고려하십시오.
## Pros and Cons
<span style="color:green;">O</span>. 클라이언트가 알지 못하는 상태에서 서비스 개체를 제어할 수 있습니다.<br />
<span style="color:green;">O</span>. 클라이언트가 신경 쓰지 않을 때 서비스 개체의 수명 주기를 관리할 수 있습니다.<br />
<span style="color:green;">O</span>. 프록시는 서비스 개체가 준비되지 않았거나 사용할 수 없는 경우에도 작동합니다.<br />
<span style="color:green;">O</span>. Open/Closed Principle. 서비스나 클라이언트를 변경하지 않고 새 프록시를 도입할 수 있습니다.<br />
<span style="color:red;">X</span>. 많은 새 클래스를 도입해야 하므로 코드가 더 복잡해질 수 있습니다.<br />
<span style="color:red;">X</span>. 서비스의 응답이 지연될 수 있습니다.