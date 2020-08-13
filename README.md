# design_pattern
Design pattern with golang. There are 3 categories in design pattern.

- Creational Patterns
- Structural Patterns
- Behavioral Patterns

## Creational Patterns
- Abstract Factory

`Abstract Factory` is creating a base class/interface for the other classes, like create a template for the other classes so the other classes have the same behavior. The factory class is responsible for the other classes.

- Builder

`Builder` is separating the construction and the representation so the same class type can represent different things.

- Object Pool

`Object Pool` is used to manage the object caching so the we can avoid to create a new object.

- Fatory Method

`Factory Method` is creating a base class for other classes for the other classes so the other classes have the same behavior.

- Prototype

`Prototype` is creating a clone from the actual class.

- Singleton

`Singleton` is one class/object only need declared once.

## Structural Pattern
- Adapter

`Adapter` is converting an interface to a class/object that needed by the client/requester.

- Bridge

`Bridge` is decoupling the abstraction from its implementation. The abstraction of object behavior is stated before the object created.

- Composite

`Composite` is dividing a process to be a hierarchical collection. The collection is a recursive collection.

- Decorator

`Decorator` is giving an additional things to a class/object. The purpose is to enchance the parent class/object.

- Fly Weight

`Fly Weight` is to share the same class/object to the other class.

- Private Class

`Private Class` is encasulate a class/object attribute so it only can be accessed from a function. The `private class` data design pattern prevents undesirable manipulation after class declaration.

## Behavoiral Pattern
- Chain of Responsibility

`Chain of Responsibility` is to divide the class/object to its purpose, create an object-oriented linked list with recursive traversal connection, like `layered architecture based application`

- Command

`Command` is to create a chain command of proccess. There is a sender and a receiver, the sender will send a command to the receiver and the receiver have to execute it.

- Intepreter

`Intepreter` is to represent an information in different representation, like a JSON object can be represent in a `string` type or `json object` type or `its own` type.

- Iterator

`Iterator` is to execute command sequencially.

- Mediator

`Mediator` is creating a middleware to connect two parties that may have `many to many` relationship.

- Memento

`Memento` is creating a snapshot or checkpoint of a process so it can be changed back (undo or rollback).

- Null Object

`Null Object` is to encapsulate empty class/object, for example null string represent an empty string, not a string with empty value (null != "")

- Observer

`Observer` is creating a one to many relationship, which the "one" is a class/object that responsible for monitoring the "many".

- State

`State` is a flag in an object/class to indicate its behavior. If the state is changing, a behavior of a class can be changed too. Like in `Fly Weight design`, the behavior of first request for a specific class and the next request is different.

- Strategy

`Strategy` is creating different algorithms and let the client chose what it want to use, like create one abstract class and several classes that implement the abstract class. The classes can be used based on the request.

- Template Method

`Template Method` is creating a template that consist of several processes/features for an operation, the class that use that template can chose the process that they need. Like template for a pc is casing, logical board, vga, cooler, psu. A custom build pc may only has logical board, vga, cooler, psu without casing.
