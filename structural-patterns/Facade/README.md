Facade
---
## Intent
**Facade** is a structural design pattern that provides a simplified interface to a library, a framework, or any other complex set of classes.
![](https://refactoring.guru/images/patterns/content/facade/facade-2x.png)

## Structure
![](https://images.velog.io/images/chrishan/post/377c6734-983e-47df-843b-8fd0e6b43409/structure-indexed-2x.png)

1. **Facade**는 하위 시스템 기능의 특정 부분에 대한 편리한 액세스를 제공합니다. 클라이언트의 요청을 어디로 보내야 하는지, 움직이는 모든 부품을 어떻게 작동해야 하는지 알고 있습니다.<br/><br/>
2. **Additional Facade** 클래스는 관련 없는 기능으로 단일 외관을 오염시켜 또 다른 복잡한 구조를 만들 수 있는 것을 방지하기 위해 생성할 수 있습니다. Additional Facade는 클라이언트와 다른 Facade 모두에서 사용할 수 있습니다.<br/><br/>
3. **Complex Subsystem**은 수십 개의 다양한 객체로 구성됩니다. 그들 모두가 의미 있는 일을 하도록 하려면 올바른 순서로 개체를 초기화하고 적절한 형식의 데이터를 제공하는 것과 같은 하위 시스템의 구현 세부 정보를 자세히 살펴봐야 합니다.<br/><br/>서브시스템 클래스는 Facade의 존재를 인식하지 못합니다. 그들은 시스템 내에서 작동하고 서로 직접 작동합니다.<br/><br/>
4. 클라이언트는 서브시스템 개체를 직접 호출하는 대신 Facade를 사용합니다.
## How to implement
1. Check whether it’s possible to provide a simpler interface than what an existing subsystem already provides. You’re on the right track if this interface makes the client code independent from many of the subsystem’s classes.<br /><br />
2. Declare and implement this interface in a new facade class. The facade should redirect the calls from the client code to appropriate objects of the subsystem. The facade should be responsible for initializing the subsystem and managing its further life cycle unless the client code already does this.<br /><br />
3. To get the full benefit from the pattern, make all the client code communicate with the subsystem only via the facade. Now the client code is protected from any changes in the subsystem code. For example, when a subsystem gets upgraded to a new version, you will only need to modify the code in the facade.<br /><br />
4. If the facade becomes too big, consider extracting part of its behavior to a new, refined facade class.
## Pros and Cons
<span style="color:#346751;">O</span> 서브시스템의 복잡성에서 코드를 분리할 수 있습니다.

<span style="color:#C84B31;">X</span> Facade는 앱의 모든 클래스에 결합된 **a god object**가 될 수 있습니다.